// Copyright 2020 Vladimir Osintsev <osintsev@gmail.com>
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
	"encoding/base64"
	"github.com/google/uuid"
)

const (
	UntilRead = iota
	OneHour
	OneDay
	OneWeek
	OneMonth
)

type Note struct {
	UUID uuid.UUID `json:"-" pg:",pk,type:uuid"`
	Data []byte    `json:"data"`
	// XXX
	Lifetime uint8
}

func (n *Note) String() string {
	b, err := n.UUID.MarshalBinary()
	if err != nil {
		return ""
	}
	return base64.RawURLEncoding.EncodeToString(b)
}
