package domain

import "fmt"

var (
	ErrPasswordNil = fmt.Errorf("password is nil")
)

// Password struct to represent a generated password
type Password struct {
	Original string `json:"original"`
	Hashed   string `json:"hashed"`
	Salt     string `json:"salt"`
}
