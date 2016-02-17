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

package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/osminogin/tornote"
)

var (
	addr    = flag.String("addr", ":8000", "The address to bind to")
	db      = flag.String("db", "./db.sqlite3", "Path to sqlite3 database")
	version = flag.Bool("version", false, "Print server version")
)

var GitCommit string

func main() {
	flag.Parse()

	if *version {
		fmt.Printf("Version: git%s\n", GitCommit)
		os.Exit(0)
	}

	server := &tornote.Server{Host: *addr}

	// Connecting to database
	if err := server.OpenDB(*db); err != nil {
		log.Fatal(err)
	}
	defer server.DB.Close()

	// Starting
	server.Run()
}
