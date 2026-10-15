package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"tinylb/internal/server"
)

func TestHomeHandler(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	server.Handler().ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Errorf("status = %d, want %d", response.Code, http.StatusOK)
	}
	if body := response.Body.String(); body != "Hello Server\n" {
		t.Errorf("body = %q, want %q", body, "Hello Server\n")
	}
}
