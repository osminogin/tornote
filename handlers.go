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
	"log"
	"net/http"

	"github.com/NebulousLabs/Sia/crypto"
	"github.com/gorilla/mux"
)

// frontPageHandler render home page.
func frontPageHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html", nil)
}

// readNoteHandler show warn screen and destroy note.
func readNoteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// XXX:
	data := vars["id"]
	// XXX: Get ciphertext from persistent storage by note id.
	// XXX: Decrypt plaintext.
	// c.Text, err = key.DecryptBytes(ciphertext)
	renderTemplate(w, "note.html", data)
}

// saveNoteHandler save secret note to persistent datastore.
func saveNoteHandler(w http.ResponseWriter, r *http.Request) {
	// XXX:
	data := "123213213213"

	key, _ := crypto.GenerateTwofishKey()
	plaintext := r.FormValue("body")
	ciphertext, _ := key.EncryptBytes([]byte(plaintext))

	// XXX: Save ciphertext in database, show key to user
	log.Printf("key = %v\n", key)
	log.Printf("ciphertext = %v\n", ciphertext)

	// XXX: Save key and note id to securecookie and redirect
	renderTemplate(w, "done.html", data)
}
