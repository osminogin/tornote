// Copyright 2016-2020 Vladimir Osintsev <osintsev@gmail.com>
//
// This file is part of Tornote.
//
// Tornote is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Tornote is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package tornote

import (
	"crypto/sha256"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
)

type Server struct {
	// Server options
	opts ServerOpts
	// PostgreSQL connection
	db *pg.DB
	// Mux router
	router *mux.Router
	// Compiled templates
	templates map[string]*template.Template
}

type ServerOpts struct {
	// Listen port
	Port uint64
	// Data source name
	DSN string
	// HTTPS only traffic allowed
	HTTPSOnly bool
	// Server Secret key
	Secret string
}

// Constructor for new Server.
func NewServer(opts ServerOpts) *Server {
	_, err := pg.ParseURL(opts.DSN)
	if err != nil {
		panic(err)
	}
	return &Server{opts: opts}
}

// Open and check database connection.
func (s *Server) connectDB() error {
	o, err := pg.ParseURL(s.opts.DSN)
	if err != nil {
		return err
	}
	s.db = pg.Connect(o)

	// XXX: Ping postgres connection
	//if err = s.db.Ping(); err != nil {
	//	return err
	//}
	return nil
}

// Creates database tables for notes if not exists.
func (s *Server) createSchema() error {
	err := s.db.CreateTable(&Note{}, &orm.CreateTableOptions{
		IfNotExists: true,
	})
	if err != nil {
		return err
	}
	return nil
}

// Generates a hash with a static length suitable for CSRF middleware.
func (s *Server) genHashFromSecret() []byte {
	h := sha256.New()
	h.Write([]byte(s.opts.Secret))
	return h.Sum(nil)
}

// Compiles templates from templates/ dir into global map.
func (s *Server) compileTemplates() (err error) {
	if s.templates == nil {
		s.templates = make(map[string]*template.Template)
	}
	// XXX:
	layout := "templates/base.html"
	pages := []string{
		"templates/index.html",
		"templates/note.html",
	}
	for _, file := range pages {
		baseName := strings.TrimLeft(file, "templates/")
		s.templates[baseName], err = template.New("").ParseFiles(file, layout)
		if err != nil {
			return err
		}
	}
	return nil
}

// Wrapper around template.ExecuteTemplate method.
func (s *Server) renderTemplate(w http.ResponseWriter, name string, data interface{}) {
	// XXX: data is context may be...
	tmpl, ok := s.templates[name]
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatalf("%s template file doesn't exists", name)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := tmpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Initialize server.
func (s *Server) Init() {
	s.router = mux.NewRouter().StrictSlash(true)

	// Setup middlewares
	csrfMiddleware := csrf.Protect(
		s.genHashFromSecret(),
		csrf.FieldName("csrf_token"),
		csrf.Secure(s.opts.HTTPSOnly),
	)
	s.router.Use(csrfMiddleware)
	if s.opts.HTTPSOnly {
		s.router.Use(RedirectToHTTPSMiddleware)
	}

	// HTTP handlers
	s.router.Handle("/", MainFormHandler(s)).Methods("GET")
	s.router.HandleFunc("/healthz", HealthStatusHandler).Methods("GET")
	s.router.PathPrefix("/public/").HandlerFunc(PublicFileHandler).Methods("GET")
	s.router.Handle("/note", CreateNoteHandler(s)).Methods("POST")
	s.router.Handle("/read/{id}", ReadRawNoteHandler(s)).Methods("GET")
	s.router.Handle("/{id}", ReadNoteHandler(s)).Methods("GET")

	// Pre-compile templates
	if err := s.compileTemplates(); err != nil {
		panic(err)
	}

	// Connecting to database
	if err := s.connectDB(); err != nil {
		panic(err)
	}

	// Bootstrap tables if not exists
	if err := s.createSchema(); err != nil {
		panic(err)
	}
}

// Listen server on specified port with opened database connection.
func (s *Server) Listen() error {
	defer s.db.Close()

	// Start the server
	if s.opts.HTTPSOnly {
		log.Println("HTTPS only traffic allowed")
	}
	log.Printf("Starting server on :%d", s.opts.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", s.opts.Port), s.router))
	return nil
}
