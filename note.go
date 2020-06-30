package tornote

import (
	"encoding/base64"
	"github.com/google/uuid"
)

type Note struct {
	UUID uuid.UUID `json:"-" pg:",pk,type:uuid"`
	Data []byte    `json:"data"`
}

func (n *Note) String() string {
	b, err := n.UUID.MarshalBinary()
	if err != nil {
		return ""
	}
	return base64.RawURLEncoding.EncodeToString(b)
}
