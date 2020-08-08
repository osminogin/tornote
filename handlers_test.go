// Copyright 2016-2020 Vladimir Osintsev <osintsev@gmail.com>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
// See the COPYING file in the main directory for details.

package tornote

import (
	"bytes"
	"golang.org/x/net/html"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

const TestRandomString = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func requestMainForm(s *Server) (w *httptest.ResponseRecorder) {
	w = httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	s.router.ServeHTTP(w, req)
	return
}

func getCSRFToken(body *bytes.Buffer) (token string) {
	root, err := html.Parse(body)
	if err != nil {
		return ""
	}
	if n, ok := getElementByName("csrf_token", root); ok {
		token = getAttributeByKey("value", n)
	}
	return
}

func getElementByName(name string, n *html.Node) (el *html.Node, ok bool) {
	for _, a := range n.Attr {
		if a.Key == "name" && a.Val == name {
			return n, true
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if el, ok = getElementByName(name, c); ok {
			return
		}
	}
	return
}

func getAttributeByKey(key string, n *html.Node) string {
	for _, attr := range n.Attr {
		if attr.Key == key {
			return attr.Val
		}
	}
	return ""
}

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

func TestCreateNoteHandler(t *testing.T) {
	srv := stubServer()
	w := requestMainForm(srv)

	data := url.Values{}
	cookies := w.Result().Cookies()
	data.Set("body", TestRandomString)
	data.Set("csrf_token", getCSRFToken(w.Body))

	w = httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/note", bytes.NewBufferString(data.Encode()))
	req.AddCookie(cookies[0])
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	srv.router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("return code %d instead %d", w.Code, http.StatusCreated)
	}
}

// Checks cross site request forgery (CSRF) protection works.
func TestCreateNoteForbidden(t *testing.T) {
	srv := stubServer()
	w := requestMainForm(srv)

	data := url.Values{}
	cookies := w.Result().Cookies()
	data.Set("body", TestRandomString)
	data.Set("csrf_token", "WRONG_TOKEN")

	w = httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/note", bytes.NewBufferString(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.AddCookie(cookies[0])
	srv.router.ServeHTTP(w, req)

	if w.Code != http.StatusForbidden {
		t.Errorf("return code %d instead %d", w.Code, http.StatusForbidden)
	}
}
