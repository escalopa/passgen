package cmd

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/argon2"
)

var (
	length       int
	useSigns     bool
	hashTimes    int
	numPasswords int
	customChars  string
	outputFormat string
)

// Password struct to represent a generated password
type Password struct {
	Original string `json:"original"`
	Hashed   string `json:"hashed"`
	Salt     string `json:"salt"`
}

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate password(s)",
	Long:  `Generate password(s) with the specified length and options.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		passwords := make([]Password, numPasswords)
		for i := 0; i < numPasswords; i++ {
			password, err := generatePassword()
			if err != nil {
				return fmt.Errorf("generate password: %v", err)
			}

			hashedPassword, salt, err := hashPasswordMultipleTimes(password, hashTimes)
			if err != nil {
				return fmt.Errorf("hash password: %v", err)
			}

			passwords[i] = Password{
				Original: password,
				Hashed:   hashedPassword,
				Salt:     salt,
			}
		}

		err := printOutput(passwords)
		if err != nil {
			return fmt.Errorf("print output: %v", err)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().IntVarP(&length, "length", "l", defaultConfig.Length, "Length of the password")
	generateCmd.Flags().BoolVarP(&useSigns, "use-signs", "s", defaultConfig.UseSigns, "Include special characters in the password")
	generateCmd.Flags().IntVarP(&hashTimes, "hash-times", "t", defaultConfig.HashTimes, "Number of times to hash the password with Argon2")
	generateCmd.Flags().IntVarP(&numPasswords, "num-passwords", "n", defaultConfig.NumPasswords, "Number of passwords to generate")
	generateCmd.Flags().StringVarP(&customChars, "custom-chars", "c", defaultConfig.CustomChars, "Custom characters to use for password generation")
	generateCmd.Flags().StringVarP(&outputFormat, "output", "o", defaultConfig.OutputFormat, "Output format (json, table, raw)")

}

func generatePassword() (string, error) {
	letters := getChars()

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

func hashPasswordMultipleTimes(password string, times int) (string, string, error) {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return "", "", fmt.Errorf("generate salt: %v", err)
	}

	hashedPassword := password
	for i := 0; i < times; i++ {
		hash := argon2.IDKey([]byte(hashedPassword), salt, 1, 64*1024, 4, 32)
		hashedPassword = base64.RawStdEncoding.EncodeToString(hash)
	}

	return hashedPassword, base64.RawStdEncoding.EncodeToString(salt), nil
}

func getChars() string {
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	if useSigns {
		letters += "!@#$%^&*()-_=+[]{}|;:,.<>?/~"
	}

	if customChars != "" {
		letters += customChars
	}

	return letters
}

func printOutput(passwords []Password) error {
	switch outputFormat {
	case "json":
		jsonOutput, err := json.MarshalIndent(passwords, "", "  ")
		if err != nil {
			return fmt.Errorf("marshal JSON output: %v", err)
		}
		fmt.Println(string(jsonOutput))
	case "table":
		// TODO: Implement table output formatting
		fmt.Println("Table output is not yet implemented")
	case "raw":
		// TODO: Implement raw output formatting
		fmt.Println("Raw output is not yet implemented")
	default:
		return fmt.Errorf("unsupported output format: %s", outputFormat)
	}
	return nil
}
