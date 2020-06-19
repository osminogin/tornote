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
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

type server struct {
	// Listen port
	Port uint64
	// Secret used for encryption/decription
	Key  string
	// Data source name
	DSN string
	// Database connection
	db   *sql.DB
}

type Server interface {
	Run() error
}

// Constructor for new server.
func NewServer(port uint64, dbPath string) *server {
	return &server{Port: port, DSN: dbPath}
}

// Open and check database connection.
func (s *server) ConnectDB() (err error) {
	if s.db, err = sql.Open("sqlite3", s.DSN); err != nil {
		return err
	}

	// Ping DB connection
	if err = s.db.Ping(); err != nil {
		return err
	}
	return nil
}

// Running daemon process.
func (s *server) Run() error {
	r := mux.NewRouter().StrictSlash(true)

	// HTTP handlers
	r.HandleFunc("/", frontPageHandler).Methods("GET")
	r.PathPrefix("/public/").HandlerFunc(publicFileHandler).Methods("GET")
	r.Handle("/note", saveNoteHandler(s.db)).Methods("POST")
	r.Handle("/{id}", readNoteHandler(s.db)).Methods("GET")

	// Prebuild templates
	if err := initTemplates(); err != nil {
		return err
	}

	// Connecting to database
	if err := s.ConnectDB(); err != nil {
		return err
	}
	defer s.db.Close()

	// Listen server on specified port
	log.Printf("Starting server on :%s", s.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", s.Port), r))
}
