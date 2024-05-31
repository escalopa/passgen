package cmd

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"math/big"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/argon2"
)

var length int
var useSigns bool
var hashTimes int

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a password",
	Long:  `Generate a password with the specified length and options.`,
	Run: func(cmd *cobra.Command, args []string) {
		password, err := generatePassword(length, useSigns)
		if err != nil {
			log.Fatalf("Failed to generate password: %v", err)
		}

		hashedPassword := hashPasswordMultipleTimes(password, hashTimes)
		fmt.Println("Generated Password:", password)
		fmt.Println("Hashed Password:", hashedPassword)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().IntVarP(&length, "length", "l", 12, "Length of the password")
	generateCmd.Flags().BoolVarP(&useSigns, "use-signs", "s", false, "Include special characters in the password")
	generateCmd.Flags().IntVarP(&hashTimes, "hash-times", "t", 1, "Number of times to hash the password with Argon2")
}

func generatePassword(length int, useSigns bool) (string, error) {
	var letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var signs = "!@#$%^&*()-_=+[]{}|;:,.<>/?"

	if useSigns {
		letters += signs
	}

	password := make([]byte, length)
	for i := range password {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		password[i] = letters[num.Int64()]
	}

	return string(password), nil
}

func hashPasswordMultipleTimes(password string, times int) string {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		log.Fatalf("Failed to generate salt: %v", err)
	}

	hashedPassword := password
	for i := 0; i < times; i++ {
		hash := argon2.IDKey([]byte(hashedPassword), salt, 1, 64*1024, 4, 32)
		hashedPassword = base64.RawStdEncoding.EncodeToString(hash)
	}

	return hashedPassword
}
