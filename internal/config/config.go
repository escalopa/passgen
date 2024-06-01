package config

import (
	_ "embed"
	"fmt"
	"os"
	"path"
	"sync"

	"github.com/spf13/viper"
)

const (
	defaultConfigName = ".passgen"
)

var (
	//go:embed default_config.yml
	defaultConfigContent string

	allowedExtensions = []string{".toml", ".json", ".yaml", ".yml"}
)

type Config struct {
	Generate GenerateConfig `yaml:"generate" json:"generate" toml:"generate"`
}

type GenerateConfig struct {
	Length     int    `yaml:"length" json:"length" toml:"length"`
	Iterations int    `yaml:"iterations" json:"iterations" toml:"iterations"`
	Characters string `yaml:"characters" json:"characters" toml:"characters"`
	Clipboard  bool   `yaml:"clipboard" json:"clipboard" toml:"clipboard"`
}

var (
	cfg       = Config{}
	parseOnce sync.Once
)

func Get() Config {
	return cfg
}

// Parse reads the config file and sets the values
// in the Config struct.
func Parse(path string) error {
	var err error

	parseOnce.Do(func() {
		parseErr := parseConfig(path)
		if parseErr != nil {
			err = fmt.Errorf("parse config: %v", parseErr)
		}
	})

	return err
}

func parseConfig(configPath string) error {
	configPath, err := getConfigPath(configPath)
	if err != nil {
		return err
	}

	viper.SetConfigFile(configPath)
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("read config file: %v", err)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		return fmt.Errorf("unmarshal config: %v", err)
	}

	return nil
}

// getConfigPath returns the path of the config file.
// If the configPath is empty, it will check the user's
// home directory for the default config file.
// If the config file doesn't exist, it will create it.
func getConfigPath(configPath string) (string, error) {
	if configPath != "" {
		if fileExists(configPath) {
			return configPath, nil
		}
		return "", fmt.Errorf("config file not found: %s", configPath)
	}

	// Get the user's home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	// Check for the default config file in the home directory
	// with the allowed extensions (.toml, .json, .yaml, .yml)
	for _, ext := range allowedExtensions {
		configPath = path.Join(homeDir, defaultConfigName+ext)
		if fileExists(configPath) {
			return configPath, nil
		}
	}

	// Config file doesn't exist, create it.
	configPath = path.Join(homeDir, defaultConfigName+".yml")
	err = os.WriteFile(configPath, []byte(defaultConfigContent), 0644)
	if err != nil {
		return "", fmt.Errorf("create default config file: %v", err)
	}

	return configPath, nil
}

// fileExists checks if a file exists
func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
