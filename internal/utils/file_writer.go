package utils

import (
	"os"
	"path/filepath"
)

// WriteLog writes a log entry to a file in the specified directory.
// If the directory does not exist, it creates it.
func WriteLog(directory string, fileName string, logData string) error {
	// Ensure the directory exists
	if err := os.MkdirAll(directory, 0755); err != nil {
		return err
	}

	// Create or open the file in append mode
	path := filepath.Join(directory, fileName)
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the log entry with a newline
	_, err = file.WriteString(logData + "\n")
	return err
}
