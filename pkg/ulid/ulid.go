package ulid

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid"
)

func createULID() (string, error) {
	entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
	ms := ulid.Timestamp(time.Now())
	id, err := ulid.New(ms, entropy)
	return id.String(), err
}

func CreateUserID() (string, error) {
	id, err := createULID()
	if err != nil {
		return "", err
	}
	return "user-" + id, err
}

func CreatePostID() (string, error) {
	id, err := createULID()
	if err != nil {
		return "", err
	}
	return "post-" + id, err
}
