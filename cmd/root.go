package cmd

import (
	"os"

	"github.com/escalopa/passgen/internal/config"
	"github.com/spf13/cobra"
)

var (
	configPath string
)

var rootCmd = &cobra.Command{
	Use:          "passgen",
	Short:        "Password Generator CLI",
	Long:         `A Command Line Tool to generate secure passwords`,
	SilenceUsage: true,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		err := config.Parse(configPath)
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&configPath, "file", "f", "", "Path to the config file. If not provided, the default values will be used.")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
