package utils

import (
	"fmt"
	"os"
)

// IsValidPath checks if the provided path exists and is a directory
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
