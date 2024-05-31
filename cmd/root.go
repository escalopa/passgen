package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	err := loadConfig()
	if err != nil {
		fmt.Printf("load config: %v\n", err)
	}
}

var rootCmd = &cobra.Command{
	Use:   "passgen",
	Short: "Password Generator CLI",
	Long:  `A Command Line Tool to generate secure passwords`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
