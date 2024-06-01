package printer

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/escalopa/passgen/internal/domain"
	"github.com/olekukonko/tablewriter"
)

type Provider struct{}

func NewProvider() *Provider {
	return &Provider{}
}

func (p *Provider) Print(password *domain.Password, outputFormat string) error {
	if password == nil {
		return domain.ErrPasswordNil
	}

	switch outputFormat {
	case "json":
		return p.printJSON(password)
	case "table":
		return p.printTable(password)
	case "raw":
		return p.printPlain(password)
	default:
		return fmt.Errorf("unknown output format: %s", outputFormat)
	}
}

// printJSON prints the passwords in a JSON format.
func (p *Provider) printJSON(password *domain.Password) error {
	data, err := json.MarshalIndent(password, "", "    ")
	if err != nil {
		return err
	}

	fmt.Println(string(data))
	return nil
}

// printTable prints the passwords in a table format.
func (p *Provider) printTable(password *domain.Password) error {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Password", "Hashed", "Salt"})
	table.Append([]string{password.Original, password.Hashed, password.Salt})
	table.Render()
	return nil
}

// printPlain prints the passwords in a raw format.
func (p *Provider) printPlain(password *domain.Password) error {
	fmt.Println(password.Hashed)
	return nil
}
