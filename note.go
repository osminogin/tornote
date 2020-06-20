package tornote

import (
	"fmt"

	"github.com/google/uuid"
)

type Note struct {
	ID	uuid.UUID	`json:"-" sql:",type:uuid"`
	Data []byte		`json:"data"`
}

type Encrypter interface {
	Encrypt(msg []byte, secret string) (data []byte, err error)
}

type Decrypter interface {
	Decrypt(data []byte, secret string) (msg []byte, err error)
}

func (u *Note) String() string {
	return fmt.Sprintf("%v", u.ID)
}
