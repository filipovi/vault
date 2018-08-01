package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/filipovi/redis"
	"github.com/filipovi/vault/generator"
	"github.com/rs/cors"
	"github.com/urfave/negroni"
	goji "goji.io"
	"goji.io/pat"
)

// Env contains the Redis client
type Env struct {
	client redis.Cacher
	scope  string
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

	password, err := generator.NewPassword(p.Name, p.Passphrase, p.Service, p.Length, p.Counter, env.scope)
	if err != nil {
		send([]byte(err.Error()), "text/plain", http.StatusNotAcceptable, w)
		return
	}

	send([]byte(password), "text/plain", http.StatusOK, w)
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
		client: *redis,
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
	env.scope = getEnv(os.Getenv("SECRET"), "")
	if env.scope == "" {
		log.Fatal("The scope is missing")
	}

	n := negroni.Classic()

	// Routing
	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/"), env.handleHomeRequest)
	mux.HandleFunc(pat.Post("/password"), env.handlePasswordRequest)

	n.UseHandler(mux)

	// Middlewares
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://0.0.0.0"},
	})
	n.Use(c)

	// Launch the Web Server
	addr := fmt.Sprintf("0.0.0.0:%s", getEnv(os.Getenv("PORT"), "3000"))
	srv := &http.Server{
		Handler:      n,
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Server run on http://" + addr)
	log.Fatal(srv.ListenAndServe())
}
