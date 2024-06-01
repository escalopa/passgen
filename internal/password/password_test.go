package password

import (
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/argon2"
)

func TestProviderGenerate(t *testing.T) {
	t.Parallel()

	p := NewProvider()

	const (
		length     = 12
		iterations = 10
		count      = 10
	)

	// Generate a password
	passwords, err := p.Generate(length, false, "", iterations, count)

	require.NoError(t, err)
	require.Len(t, passwords, count)

	// Verify the password length & hash
	for _, pass := range passwords {
		require.Len(t, pass.Original, length)
		require.NotEmpty(t, pass.Hashed)
		require.NotEmpty(t, pass.Salt)

		// Verify the password hash
		hash := hashPasswordWithSalt(pass.Original, pass.Salt, iterations)
		require.Equal(t, hash, pass.Hashed)
	}
}

func TestHashPassword(t *testing.T) {
	t.Parallel()

	const (
		password   = "password"
		iterations = 10
	)

	hash, salt, err := hashPassword(password, iterations)
	require.NoError(t, err)
	require.NotEmpty(t, hash)
	require.NotEmpty(t, salt)

	// Verify the hash
	hashedPassword := password
	saltStr := salt

	for i := 0; i < iterations; i++ {
		hash := argon2.IDKey([]byte(hashedPassword), []byte(saltStr), 1, 64*1024, 4, 32)
		hashedPassword = base64.RawStdEncoding.EncodeToString(hash)
	}

	require.Equal(t, hashedPassword, hash)
}
