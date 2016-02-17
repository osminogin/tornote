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
	"errors"
	"html/template"
	"log"
	"net/http"
)

var templates map[string]*template.Template

// Load and compile templates files from bindata.
func initTemplates() error {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}

	layout, err := Asset("templates/layout/base.html")
	if err != nil {
		log.Fatal(err)
	}

	pages, err := AssetDir("templates")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range pages {
		// Skip layout dir
		if file == "layout" {
			continue
		}
		// Get template data from bindata
		templateData, err := Asset("templates/" + file)
		if err != nil {
			return err
		}
		// Compile layout
		target, err := template.New(file).Parse(string(layout))
		if err != nil {
			return err
		}
		// Compile target template
		templates[file], err = target.Parse(string(templateData))
		if err != nil {
			return err
		}
	}

	return nil
}

// Wrapper around template.ExecuteTemplate method.
func renderTemplate(w http.ResponseWriter, name string, data interface{}) error {
	tmpl, ok := templates[name]
	if !ok {
		return errors.New("This template doesn't exist")
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl.ExecuteTemplate(w, "base", data)

	return nil
}
