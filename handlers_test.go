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
