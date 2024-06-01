package cmd

import (
	"fmt"

	"github.com/escalopa/passgen/internal/config"
	"github.com/escalopa/passgen/internal/password"
	"github.com/escalopa/passgen/internal/printer"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	passProvider  = password.NewProvider()
	printProvider = printer.NewProvider()
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate password(s)",
	Long:  `Generate password(s) with the specified length and options.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg := config.Get()

		// Used for debugging
		// fmt.Printf("%+v\n", cfg)

		passwords, err := passProvider.Generate(
			cfg.Generate.Length,
			cfg.Generate.Iterations,
			cfg.Generate.Characters,
		)
		if err != nil {
			return err
		}

		err = printProvider.Print(passwords, cfg.Generate.Clipboard)
		if err != nil {
			return fmt.Errorf("print output: %v", err)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().IntP("length", "l", 12, "Length of the password")
	generateCmd.Flags().IntP("iterations", "i", 1, "Number of times to hash the password")
	generateCmd.Flags().StringP("characters", "c", "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789", "Characters to use for the password")
	generateCmd.Flags().BoolP("clipboard", "b", false, "Copy the password to the clipboard")

	viper.BindPFlag("generate.length", generateCmd.Flags().Lookup("length"))
	viper.BindPFlag("generate.iterations", generateCmd.Flags().Lookup("iterations"))
	viper.BindPFlag("generate.characters", generateCmd.Flags().Lookup("characters"))
	viper.BindPFlag("generate.clipboard", generateCmd.Flags().Lookup("clipboard"))

	viper.BindEnv("generate.length", "PASSGEN_GENERATE_LENGTH")
	viper.BindEnv("generate.iterations", "PASSGEN_GENERATE_ITERATIONS")
	viper.BindEnv("generate.characters", "PASSGEN_GENERATE_CHARACTERS")
	viper.BindEnv("generate.clipboard", "PASSGEN_GENERATE_CLIPBOARD")
}
