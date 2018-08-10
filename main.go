package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/filipovi/redis"
	"github.com/filipovi/vault/generator"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

// Env contains the Redis client
type Env struct {
	cache redis.Cacher
	scope string
}

// Payload as used in the POST request
type Payload struct {
	Name       string `json:"name"`
	Passphrase string `json:"passphrase"`
	Service    string `json:"service"`
	Length     int    `json:"length"`
	Counter    int    `json:"counter"`
}

func (env *Env) handleHomeRequest(w http.ResponseWriter, req *http.Request) {
	send([]byte("The service is working!"), "text/plain", http.StatusOK, w)
}

func (env *Env) handlePasswordRequest(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	defer req.Body.Close()
	p := &Payload{}
	err := decoder.Decode(p)
	if err != nil {
		send([]byte(err.Error()), "text/plain", http.StatusBadRequest, w)

		return
	}

	key := env.cache.GetHashKey(getKey(env.scope, p), "password:%s")
	password, err := env.cache.Load(key)
	if err == nil && password != "" {
		send([]byte(password), "text/plain", http.StatusOK, w)
		return

	}

	password, err = generator.NewPassword(p.Name, p.Passphrase, p.Service, p.Length, p.Counter, env.scope)
	if err != nil {
		send([]byte(err.Error()), "text/plain", http.StatusNotAcceptable, w)
		return
	}

	if err = env.cache.Save(key, []byte(password)); err != nil {
		send([]byte(err.Error()), "text/plain", http.StatusNotAcceptable, w)
		return
	}

	send([]byte(password), "text/plain", http.StatusOK, w)
}

func getKey(scope string, p *Payload) string {
	var b strings.Builder
	b.Write([]byte(scope))
	b.Write([]byte(p.Name))
	b.Write([]byte(p.Passphrase))
	b.Write([]byte(p.Service))
	b.WriteString(fmt.Sprint(p.Length))
	b.WriteString(fmt.Sprint(p.Counter))

	return b.String()
}

func send(content []byte, contentType string, status int, w http.ResponseWriter) {
	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Content-Length", fmt.Sprintf("%v", len(content)))
	w.WriteHeader(status)
	w.Write(content)
}

func failOnError(err error, msg string) {
	if err == nil {
		return
	}
	log.Fatalf("%s: %s", msg, err)
	panic(fmt.Sprintf("%s: %s", msg, err))
}

func connect(file string) (*Env, error) {
	path, err := filepath.Abs(file)
	if err != nil {
		log.Fatal(err)
	}
	redis, err := redis.New(path)
	if nil != err {
		return nil, err
	}
	log.Println("Redis connected!")

	env := &Env{
		cache: *redis,
	}

	return env, nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func main() {
	env, err := connect("config.json")
	failOnError(err, "Failed to connect to Redis")
	env.scope = getEnv("SECRET", "")
	if env.scope == "" {
		log.Fatal("The scope is missing")
	}

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Get("/", env.handleHomeRequest)
	r.Post("/password", env.handlePasswordRequest)

	// Launch the Web Server
	addr := fmt.Sprintf("0.0.0.0:%s", getEnv("PORT", "3000"))
	srv := &http.Server{
		Handler:      r,
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Server run on http://" + addr)
	log.Fatal(srv.ListenAndServe())
}
