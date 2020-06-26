package tornote

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMainFormHandler(t *testing.T) {
	w := httptest.NewRecorder()
	s := stubServer()

	req, _ := http.NewRequest("GET", "/", nil)
	s.router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("return code %d instead %d", w.Code, http.StatusOK)
	}
	// TODO: Check substring in the page
}

func TestHealthStatusHandler(t *testing.T) {
	w := httptest.NewRecorder()
	s := stubServer()

	req, _ := http.NewRequest("GET", "/healthz", nil)
	s.router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("return code %d instead %d", w.Code, http.StatusOK)
	}
}
