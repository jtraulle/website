package http

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCorsMiddleware(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		wrapped := CorsMiddleware("https://www.example.net")(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("OK"))
		}))
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", nil)
		wrapped.ServeHTTP(w, r)
		if h := w.Header().Get("Access-Control-Allow-Origin"); h != "https://www.example.net" {
			t.Errorf("Unexpected header value %v", h)
		}
	})
}

func TestContentTypeMiddleware(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		wrapped := ContentTypeMiddleware("application/json")(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("OK"))
		}))
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", nil)
		wrapped.ServeHTTP(w, r)
		if h := w.Header().Get("Content-Type"); h != "application/json" {
			t.Errorf("Unexpected header value %v", h)
		}
	})
}

func TestOptoutMiddleware(t *testing.T) {
	wrapped := OptoutMiddleware("optout")(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hey there"))
	}))
	t.Run("with header", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", nil)
		r.AddCookie(&http.Cookie{
			Name: "optout",
		})
		wrapped.ServeHTTP(w, r)

		if w.Code != http.StatusNoContent {
			t.Errorf("Unexpected status code %d", w.Code)
		}

		if w.Body.String() != "" {
			t.Errorf("Unexpected response body %s", w.Body.String())
		}
	})
	t.Run("without header", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", nil)
		wrapped.ServeHTTP(w, r)

		if w.Code != http.StatusOK {
			t.Errorf("Unexpected status code %d", w.Code)
		}

		if w.Body.String() != "hey there" {
			t.Errorf("Unexpected response body %s", w.Body.String())
		}
	})
}

func TestUserCookieMiddleware(t *testing.T) {
	wrapped := UserCookieMiddleware("user", 1)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		value := r.Context().Value(1)
		fmt.Fprintf(w, "value is %v", value)
	}))
	t.Run("no cookie", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", nil)
		wrapped.ServeHTTP(w, r)
		if w.Code != http.StatusBadRequest {
			t.Errorf("Unexpected status code %v", w.Code)
		}
		if !strings.Contains(w.Body.String(), "received no or blank user identifier") {
			t.Errorf("Unexpected body %s", w.Body.String())
		}
	})

	t.Run("no value", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", nil)
		wrapped.ServeHTTP(w, r)
		r.AddCookie(&http.Cookie{
			Name:  "user",
			Value: "",
		})
		if w.Code != http.StatusBadRequest {
			t.Errorf("Unexpected status code %v", w.Code)
		}
		if !strings.Contains(w.Body.String(), "received no or blank user identifier") {
			t.Errorf("Unexpected body %s", w.Body.String())
		}
	})

	t.Run("ok", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", nil)
		r.AddCookie(&http.Cookie{
			Name:  "user",
			Value: "token",
		})
		wrapped.ServeHTTP(w, r)
		if w.Code != http.StatusOK {
			t.Errorf("Unexpected status code %v", w.Code)
		}
		if w.Body.String() != "value is token" {
			t.Errorf("Unexpected body %s", w.Body.String())
		}
	})
}
