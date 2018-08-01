package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Mock Redis
type mockClient struct{}

func (mclient *mockClient) Load(key string) (string, error) {
	data := ""
	err := errors.New("Not found")

	return data, err
}

func (mclient *mockClient) MultiLoad(keys []string) ([]interface{}, error) {
	return nil, nil
}

func (mclient *mockClient) Save(key string, data []byte) error {
	return nil
}

func (mclient *mockClient) GetHashKey(data, key string) string {
	h := sha1.New()
	h.Write([]byte(data))
	return fmt.Sprintf(key, hex.EncodeToString(h.Sum(nil)))
}

func (mclient *mockClient) FlushAll() error {
	return nil
}

// Tests
func TestGetHomepageRequest(t *testing.T) {
	expected := "The service is working!"
	env := Env{&mockClient{}}
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	http.HandlerFunc(env.handleHomeRequest).ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, expected, rec.Body.String())
}

func TestPassword(t *testing.T) {
	expected := "j-3Bt4dg9:(Lk#7i7OFCaH@vm6M6ZvxL"
	payload := Payload{
		Name:       "name",
		Passphrase: "passphrase",
		Service:    "service",
		Length:     32,
	}

	env := Env{&mockClient{}}
	rec := httptest.NewRecorder()
	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "/password", bytes.NewBuffer(body))
	http.HandlerFunc(env.handlePasswordRequest).ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
	b := rec.Body.String()
	assert.Equal(t, expected, b)
	assert.Equal(t, 32, len(b))
}
