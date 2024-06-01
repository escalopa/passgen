package printer

import (
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/escalopa/passgen/internal/domain"
)

type Provider struct{}

func NewProvider() *Provider {
	return &Provider{}
}

func (p *Provider) Print(password *domain.Password, toClipboard bool) error {
	if password == nil {
		return domain.ErrPasswordNil
	}

	if toClipboard {
		err := clipboard.WriteAll(password.Hashed)
		if err != nil {
			return fmt.Errorf("write to clipboard: %v", err)
		}
		fmt.Println("Password copied to clipboard")
	}

	fmt.Printf("Original: %s\n", password.Original)
	fmt.Printf("Hashed: %s\n", password.Hashed)
	fmt.Printf("Salt: %s\n", password.Salt)

	return nil
}
