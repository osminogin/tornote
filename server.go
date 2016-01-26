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
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
)

type server struct {
	config                             string
	dbName, dbHost, dbUser, dbPassword string
}

const (
	secretKey   = "changeme01234567" // 16 bytes
	templateDir = "/home/oc/go/src/github.com/osminogin/tornote/templates/"
	dbName      = "d71pmcmein1n6n"
	dbHost      = "ec2-54-243-245-159.compute-1.amazonaws.com"
	dbUser      = "qutntbdraptjak"
	dbPassword  = "BcqJE5qiSr3htZc_8e9mG2lM8t"
)

var templates map[string]*template.Template

func init() {
	// Load templates on application initialization.
	if templates == nil {
		templates = make(map[string]*template.Template)
	}

	includes, err := filepath.Glob(templateDir + "layouts/*.html")
	checkErr(err)
	pages, err := filepath.Glob(templateDir + "*.html")
	checkErr(err)

	for _, file := range pages {
		targets := append(includes, file)
		templates[filepath.Base(file)] = template.Must(template.ParseFiles(targets...))
	}

	// Command arguments
	// XXX:
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", frontPageHandler).Methods("GET")
	r.HandleFunc("/", saveNoteHandler).Methods("POST")
	r.HandleFunc("/n/{id}", readNoteHandler).Methods("GET")
	http.Handle("/", r)

	log.Println("Tornote server started...")
	log.Panic(http.ListenAndServe(":8081", nil))
}
