package printer

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/escalopa/passgen/internal/domain"
	"github.com/olekukonko/tablewriter"
)

type Provider struct{}

func NewProvider() *Provider {
	return &Provider{}
}

func (p *Provider) Print(passwords []domain.Password, outputFormat string) error {
	switch outputFormat {
	case "json":
		return p.printJSON(passwords)
	case "table":
		return p.printTable(passwords)
	case "raw":
		return p.printRaw(passwords)
	default:
		return fmt.Errorf("unknown output format: %s", outputFormat)
	}
}

// printJSON prints the passwords in a JSON format.
func (p *Provider) printJSON(passwords []domain.Password) error {
	data, err := json.MarshalIndent(passwords, "", "    ")
	if err != nil {
		return err
	}

	fmt.Println(string(data))
	return nil
}

// printTable prints the passwords in a table format.
func (p *Provider) printTable(passwords []domain.Password) error {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Index", "Password", "Hashed", "Salt"})
	for i, p := range passwords {
		idx := strconv.Itoa(i + 1)
		table.Append([]string{idx, p.Original, p.Hashed, p.Salt})
	}
	table.Render()
	return nil
}

// printRaw prints the passwords in a raw format.
func (p *Provider) printRaw(passwords []domain.Password) error {
	for _, pwd := range passwords {
		fmt.Println(pwd)
	}
	return nil
}
