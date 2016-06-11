package tornote

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFrontPage(t *testing.T) {
	//req, _ := http.NewRequest("GET", "", nil)
	w := httptest.NewRecorder()
	if w.Code != http.StatusOK {
		t.Errorf("Home page didn't return %v", http.StatusOK)
	}
}
