package cmd

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

const defaultConfigContent = `# Password Generator Config
length: 12
useSigns: false
hashTimes: 1
numPasswords: 1
customChars: ""
outputFormat: "json"
`

type Config struct {
	Length       int    `yaml:"length"`
	UseSigns     bool   `yaml:"useSigns"`
	HashTimes    int    `yaml:"hashTimes"`
	NumPasswords int    `yaml:"numPasswords"`
	CustomChars  string `yaml:"customChars"`
	OutputFormat string `yaml:"outputFormat"`
}

var defaultConfig = Config{
	Length:       12,
	UseSigns:     false,
	HashTimes:    1,
	NumPasswords: 1,
	CustomChars:  "",
	OutputFormat: "json",
}

func loadConfig() error {
	configPath, err := getConfigPath()
	if err != nil {
		return fmt.Errorf("get config path: %v", err)
	}

	// Read the config file
	configData, err := os.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("read config file: %v", err)
	}

	// Unmarshal the YAML data into Config struct
	var config Config
	err = yaml.Unmarshal(configData, &config)
	if err != nil {
		return fmt.Errorf("unmarshal config data: %v", err)
	}

	// Update default values with values from config file
	defaultConfig = config

	return nil
}

func getConfigPath() (string, error) {
	// Get the user's home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	// Construct the path to the config file
	configPath := homeDir + "/.passgen.yaml"
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		// Check if the config file exists with suffix .yml
		configPath = homeDir + "/.passgen.yml"
		if _, err := os.Stat(configPath); err == nil {
			return configPath, nil
		}

		// Config file doesn't exist, create it
		err = os.WriteFile(configPath, []byte(defaultConfigContent), 0644)
		if err != nil {
			return "", err
		}
	}

	return configPath, nil
}
