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

package main

import (
	"log"

	"github.com/osminogin/tornote"
	"github.com/spf13/viper"
)

var (
	GitCommit string
)

func main() {
	var srv *tornote.Server
	var opts tornote.ServerOpts

	// Set default settings
	v := viper.New()
	v.SetDefault("PORT", 8000)
	v.SetDefault("DATABASE_URL", "postgres://postgres:postgres@localhost/postgres")
	v.SetDefault("HTTPS_ONLY", false)
	v.SetDefault("VERSION", GitCommit)

	// Read configuration file and environment
	v.SetConfigName(".env")
	v.SetConfigType("dotenv")
	v.AddConfigPath(".")
	v.ReadInConfig()
	v.AutomaticEnv()

	// Server initialization and start
	opts.Port = v.GetUint64("PORT")
	opts.DSN = v.GetString("DATABASE_URL")
	opts.Secret = v.GetString("SECRET_KEY")
	opts.HTTPSOnly = v.GetBool("HTTPS_ONLY")

	srv = tornote.NewServer(opts)
	srv.Init()
	log.Fatal(srv.Listen())
}
