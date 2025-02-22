package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

const version = "1.0.0"
const appName = "baloney"

type Config struct {
	DefaultPath string `json:"default_path"`
}

func GetDownloadsFolder() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		return "C:\\Downloads"
	}
	return filepath.Join(homeDir, "Downloads")
}

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

func LoadConfig() Config {
	configFile, err := GetConfigPath()
	if err != nil {
		fmt.Println("Error determining config path:", err)
		return Config{}
	}

	// Check if config file exists
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		return Config{DefaultPath: GetDownloadsFolder()}
	}

	// Read the file
	data, err := os.ReadFile(configFile)
	if err != nil {
		fmt.Println("Error reading config file:", err)
		return Config{}
	}

	// Parse JSON
	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		fmt.Println("Error parsing config file:", err)
		return Config{}
	}

	return config
}

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

	// Write to file
	if err := os.WriteFile(configFile, data, 0644); err != nil {
		fmt.Println("Error writing config file:", err)
	}
}

func IsValidPath(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		fmt.Println("Error: Path does not exist:", path)
		return false
	}
	if err != nil {
		fmt.Println("Error checking path:", err)
		return false
	}
	if !info.IsDir() {
		fmt.Println("Error: Specified path is not a directory:", path)
		return false
	}
	return true
}

func main() {
	config := LoadConfig()
	showVersion := flag.Bool("v", false, "Show version")
	torrentFile := flag.String("t", "", ".torrent file path")
	magnetLink := flag.String("m", "", "Magnet link")
	downloadPath := flag.String("p", "", "Download path")
	setPath := flag.String("d", "", "Set a new default path")
	showPath := flag.Bool("show", false, "Show current default path")

	flag.Usage = func() {
		fmt.Println("Usage: baloney <flag> [arguments]")
		fmt.Println("\nFlags:")
		fmt.Println("  -v \t\t Show version")
		fmt.Println("  -t \t\t .torrent file path")
		fmt.Println("  -m \t\t Magnet link")
		fmt.Println("  -p \t\t Specify Download path")
		fmt.Println("  -d \t\t Set Default Download path")
		fmt.Println("  -show \t Show current default path")
		fmt.Println()
	}

	flag.Parse()

	if *showPath {
		fmt.Println("Current default path:", config.DefaultPath)
		return
	}

	if *setPath != "" {
		if IsValidPath(*setPath) {
			config.DefaultPath = *setPath
			SaveConfig(config)
			fmt.Println("New default path set to:", *setPath)
		} else {
			fmt.Println("Failed to set new default path. Please provide a valid directory.")
		}
		return
	}

	if *setPath != "" {
		config.DefaultPath = *setPath
		SaveConfig(config)
		fmt.Println("New default path set to:", *setPath)
		return
	}

	if *showVersion {
		fmt.Printf("baloney version %s\n", version)
		os.Exit(0)
	}

	if *torrentFile == "" && *magnetLink == "" {
		fmt.Println("Error: You must provide either a .torrent file (-t) or a magnet link (-m).")
		flag.Usage()
		os.Exit(1)
	}

	if *torrentFile != "" {
		fmt.Println("Torrent file:", *torrentFile)
	}
	if *magnetLink != "" {
		fmt.Println("Magnet link:", *magnetLink)
	}
	if *downloadPath != "" {
		fmt.Println("Download path:", *downloadPath)
	}

}
