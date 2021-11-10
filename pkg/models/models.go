package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("no matching record ")
var ErrDuplicateEmail = errors.New("email already exist")
var ErrInvalidCredentials = errors.New("invalid email or password")

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
}
