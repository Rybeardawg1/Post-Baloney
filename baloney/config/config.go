package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const appName = "baloney"

// Config structure
type Config struct {
	DefaultPath string `json:"default_path"`
}

// GetConfigPath returns the full path to the config file
func GetConfigPath() (string, error) {
	appData, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	configDir := filepath.Join(appData, appName)
	err = os.MkdirAll(configDir, os.ModePerm)
	if err != nil {
		return "", err
	}
	return filepath.Join(configDir, "config.json"), nil
}

// LoadConfig reads and loads the config file
func LoadConfig() Config {
	configFile, err := GetConfigPath()
	if err != nil {
		fmt.Println("Error determining config path:", err)
		return Config{}
	}

	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		return Config{DefaultPath: filepath.Join(os.UserHomeDir(), "Downloads")}
	}

	data, err := os.ReadFile(configFile)
	if err != nil {
		fmt.Println("Error reading config file:", err)
		return Config{}
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		fmt.Println("Error parsing config file:", err)
		return Config{}
	}

	return config
}

// SaveConfig writes the config to file
func SaveConfig(config Config) {
	configFile, err := GetConfigPath()
	if err != nil {
		fmt.Println("Error determining config path:", err)
		return
	}

	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		fmt.Println("Error encoding config:", err)
		return
	}

	if err := os.WriteFile(configFile, data, 0644); err != nil {
		fmt.Println("Error writing config file:", err)
	}
}
