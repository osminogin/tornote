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
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

type Server struct {
	DB   *sql.DB
	Host string
	Key  string
}

// Open and check database connection.
func (s *Server) OpenDB(path string) (err error) {
	if s.DB, err = sql.Open("sqlite3", path); err != nil {
		return err
	}

	// Ping DB connection
	if err = s.DB.Ping(); err != nil {
		return err
	}

	return nil
}

// Running daemon process.
func (s *Server) Run() {
	r := mux.NewRouter().StrictSlash(true)

	// HTTP handlers
	r.HandleFunc("/", frontPageHandler).Methods("GET")
	r.PathPrefix("/public/").HandlerFunc(publicFileHandler).Methods("GET")
	r.Handle("/note", saveNoteHandler(s.DB)).Methods("POST")
	r.Handle("/{id}", readNoteHandler(s.DB)).Methods("GET")

	// Prebuild templates
	if err := initTemplates(); err != nil {
		log.Fatal(err)
	}

	// Listen server on port 8080
	log.Printf("Starting tornote server on %s", s.Host)
	log.Fatal(http.ListenAndServe(s.Host, r))
}
