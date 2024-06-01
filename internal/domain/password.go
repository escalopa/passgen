package domain

// Password struct to represent a generated password
type Password struct {
	Original string `json:"original"`
	Hashed   string `json:"hashed"`
	Salt     string `json:"salt"`
}
