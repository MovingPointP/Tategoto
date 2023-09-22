package ulid

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid"
)

func CreateULID() (string, error) {
	entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
	ms := ulid.Timestamp(time.Now())
	id, err := ulid.New(ms, entropy)
	return id.String(), err
}
