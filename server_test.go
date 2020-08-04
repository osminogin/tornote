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

const (
	TestDSN    = "postgres://postgres:postgres@localhost/testdb"
	TestPort   = 31337
	TestSecret = "0123456789"
)

func stubServer() *server {
	s := NewServer(TestPort, TestDSN, TestSecret)
	s.Init()
	return s
}

//func TestNewServer(t *testing.T) {
//	s := stubServer()
//	if s.DSN != TestDSN && s.Port != TestPort {
//		t.Fatal("can not initialize server")
//	}
//}
//
//func TestServer_Run(t *testing.T) {
//	s := stubServer()
//
//}
