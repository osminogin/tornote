package tornote

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestFrontPage(t *testing.T) {
	w := httptest.NewRecorder()
	r := mux.NewRouter()

	req, _ := http.NewRequest("GET", "/", nil)
	r.HandleFunc("/", mainFormHandler).Methods("GET")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Frontpage return code %v instead %v", w.Code, http.StatusOK)
	}
}
