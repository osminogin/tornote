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
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/google/uuid"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/go-pg/pg/v10"
)

// frontPageHandler render home page.
func frontPageHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html", nil)
}

// publicFileHandler get file from bindata or return not found error.
func publicFileHandler(w http.ResponseWriter, r *http.Request) {
	uri := r.URL.Path[1:]
	log.Print(uri)
	http.ServeFile(w, r, uri)

	//data, err := Asset()
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusNotFound)
	//	return
	//}

	// Set headers by file extension
	//switch filepath.Ext(r.URL.Path[1:]) {
	//case ".js":
	//	w.Header().Set("Content-Type", "application/javascript; charset=utf-8")
	//case ".css":
	//	w.Header().Set("Content-Type", "text/css")
	//}
	//
	//w.Write(data)
}

// readNoteHandler print encrypted data for client-side decrypt and destroy note.
func readNoteHandler(db *pg.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		// Valid UUID required.
		id, err := uuid.Parse(vars["id"])
		if err != nil {
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		note := &Note{ID: id}

		// Get encrypted note or return 404
		err = db.Select(note)
		if err != nil {
			http.NotFound(w, r)
			return
		}
		// Deferred note deletion
		defer func() {
			db.Delete(note)
		}()

		// Print encrypted note to user
		renderTemplate(w, "note.html", note.Data)
	})
}

// saveNoteHandler save secret note to persistent datastore and return note ID.
func saveNoteHandler(db *pg.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		encrypted := r.FormValue("body")
		secret := make([]byte, 11)

		// Generate random data for note id
		_, err := rand.Read(secret)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Encode note id with URL safe format
		id := base64.RawURLEncoding.EncodeToString(secret)

		// Save data to database
		_, err = db.Exec("INSERT INTO notes (id, encrypted) VALUES (?, ?)", id, encrypted)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Fprint(w, id)
	})
}
