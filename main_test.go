package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHomeRouteRendersHelloWorld(t *testing.T) {
	router := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}

	body := w.Body.String()
	if !strings.Contains(body, "Hello World") {
		t.Fatalf("expected body to contain 'Hello World', got:\n%s", body)
	}
}
