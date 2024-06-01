package password

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"math/big"

	"github.com/escalopa/passgen/internal/domain"
	"golang.org/x/crypto/argon2"
)

type Provider struct{}

func NewProvider() *Provider {
	return &Provider{}
}

func (p *Provider) Generate(length int, iterations int, characters string) (*domain.Password, error) {
	// Generate password
	password, err := generatePassword(length, characters)
	if err != nil {
		return nil, fmt.Errorf("generate password: %v", err)
	}

	// Hash password
	hash, salt, err := hashPassword(password, iterations)
	if err != nil {
		return nil, fmt.Errorf("hash password: %v", err)
	}

	return &domain.Password{
		Original: password,
		Hashed:   hash,
		Salt:     salt,
	}, nil
}

// generatePassword generates a random password using the letters provided
func generatePassword(length int, letters string) (string, error) {
	password := make([]byte, length)

	// Generate a random password using
	// the letters provided and the length
	for i := range password {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		password[i] = letters[num.Int64()]
	}

	return string(password), nil
}

// hashPassword generates a random salt and hashes the password
// multiple times using Argon2 algorithm.
func hashPassword(password string, times int) (string, string, error) {
	// Generate a random salt
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return "", "", fmt.Errorf("generate salt: %v", err)
	}

	saltStr := base64.RawStdEncoding.EncodeToString(salt)
	hashedPassword := hashPasswordWithSalt(password, saltStr, times)

	return hashedPassword, saltStr, nil
}

func hashPasswordWithSalt(password string, salt string, times int) string {
	hashedPassword := password

	// Hash the password multiple times
	for i := 0; i < times; i++ {
		hash := argon2.IDKey([]byte(hashedPassword), []byte(salt), 1, 64*1024, 4, 32)
		hashedPassword = base64.RawStdEncoding.EncodeToString(hash)
	}

	return hashedPassword
}
