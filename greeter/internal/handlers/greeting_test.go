package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"

	"greeter/internal/gen"
	"greeter/internal/handlers"
)

func newRouter() http.Handler {
	r := chi.NewRouter()
	return gen.HandlerWithOptions(gen.NewStrictHandler(handlers.NewServer(), nil), gen.ChiServerOptions{
		BaseRouter:       r,
		ErrorHandlerFunc: handlers.WriteJSONError,
	})
}

func TestGetGreetingWithName(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/hello?name=Ada", nil)
	rec := httptest.NewRecorder()

	newRouter().ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var body gen.Greeting
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if body.Message != "Hello, Ada!" {
		t.Fatalf("expected %q, got %q", "Hello, Ada!", body.Message)
	}
}

func TestGetGreetingMissingName(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	rec := httptest.NewRecorder()

	newRouter().ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d: %s", rec.Code, rec.Body.String())
	}

	var body gen.Error
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if body.Message == "" {
		t.Fatal("expected a non-empty error message")
	}
}

func TestGetGreetingEmptyName(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/hello?name=", nil)
	rec := httptest.NewRecorder()

	newRouter().ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d: %s", rec.Code, rec.Body.String())
	}

	var body gen.Error
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if body.Message == "" {
		t.Fatal("expected a non-empty error message")
	}
}
