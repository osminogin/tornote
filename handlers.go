// Copyright 2016 Vladimir Osintsev <osintsev@gmail.com>
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
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

// frontPageHandler render home page.
func frontPageHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html", nil)
}

// readNoteHandler print encrypted data for client-side decrypt and destroy note.
func readNoteHandler(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var encrypted string
		vars := mux.Vars(r)

		// Get encrypted note or return error
		err := db.QueryRow("SELECT encrypted from notes where id = ?", vars["id"]).Scan(&encrypted)
		switch {
		case err == sql.ErrNoRows:
			http.Error(w, err.Error(), http.StatusNotFound)
		case err != nil:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		// Deferred destroy note
		defer func() {
			db.QueryRow("DELETE from notes where id = ?", vars["id"])
		}()

		// Print encrypted note to user
		renderTemplate(w, "note.html", encrypted)
	})
}

// saveNoteHandler save secret note to persistent datastore and show ID.
func saveNoteHandler(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		encrypted := r.FormValue("encrypted")
		result, err := db.Exec("INSERT INTO notes VALUES (encrypted, ?)", encrypted)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		// XXX: Show encryption key for this note to user
		renderTemplate(w, "done.html", result.LastInsertId())
	})
}
