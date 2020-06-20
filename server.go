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
	"fmt"
	"log"
	"net/http"
	"html/template"
	"strings"

	"github.com/gorilla/mux"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	_ "github.com/mattn/go-sqlite3"
)

// XXX: Move inside server struct.
var Templates map[string]*template.Template

type server struct {
	// Listen port
	Port uint64
	// Secret used for encryption/decription
	Key  string
	// Data source name
	DSN string
	// PostgreSQL connection
	db   *pg.DB
}

type Server interface {
	Run() error
}

// Constructor for new server.
func NewServer(port uint64, dsn string) *server {
	_, err := pg.ParseURL(dsn)
	if err != nil {
		panic(err)
	}
	return &server{Port: port, DSN: dsn}
}

// Open and check database connection.
func (s *server) connectDB() error {
	opt, err := pg.ParseURL(s.DSN)
	if err != nil {
		return err
	}
	s.db = pg.Connect(opt)

	// Ping postgres connection
	if err = s.db.Ping(nil); err != nil {
		return err
	}
	return nil
}

// Creates database tables for notes if not exists.
func (s *server) createSchema() error {
	err := s.db.CreateTable(&Note{}, &orm.CreateTableOptions{
		IfNotExists: true,
	})
	if err != nil {
		return err
	}
	return nil
}

// Compiles templates from templates/ dir into global map.
func compileTemplates() (err error) {
	if Templates == nil {
		Templates = make(map[string]*template.Template)
	}
	// XXX:
	layout := "templates/base.html"
	pages := []string{
		"templates/index.html",
		"templates/note.html",
	}
	for _, file := range pages {
		baseName := strings.TrimLeft(file, "templates/")
		Templates[baseName], err = template.New("").ParseFiles(file, layout)
		if err != nil {
			return err
		}
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

	// Connecting to database
	if err := s.connectDB(); err != nil {
		return err
	}
	defer s.db.Close()

	// Bootstrap database if not exists
	if err := s.createSchema(); err != nil {
		return err
	}

	// Pre-compile templates
	if err := compileTemplates(); err != nil {
		return err
	}

	// Listen server on specified port
	log.Printf("Starting server on :%d", s.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", s.Port), r))
	return nil
}
