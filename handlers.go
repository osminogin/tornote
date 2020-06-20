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
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// frontPageHandler render home page.
func frontPageHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html", nil)
}

// publicFileHandler get file from bindata or return not found error.
func publicFileHandler(w http.ResponseWriter, r *http.Request) {
	uri := r.URL.Path[1:]
	http.ServeFile(w, r, uri)
}

// readNoteHandler print encrypted data for client-side decrypt and destroy note.
func readNoteHandler(s *server) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		raw, _ := base64.RawURLEncoding.DecodeString(vars["id"])
		id, err := uuid.FromBytes(raw)
		if err != nil {
			http.NotFound(w, r)
			return
		}

		n := &Note{UUID: id}

		// Get encrypted n or return 404
		err = s.db.Select(n)
		if err != nil {
			http.NotFound(w, r)
			return
		}
		// Deferred n deletion
		defer func() {
			s.db.Delete(n)
		}()

		// Print encrypted n to user
		renderTemplate(w, "note.html", string(n.Data))
	})
}

// createNoteHandler save secret note to persistent datastore and return note ID.
func createNoteHandler(s *server) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		n := &Note{
			UUID: uuid.New(),
			Data: []byte(r.FormValue("body")),
		}

		err := s.db.Insert(n)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		b, _ := n.UUID.MarshalBinary()
		marshalled := base64.RawURLEncoding.EncodeToString(b)
		fmt.Fprint(w, marshalled)
	})
}
