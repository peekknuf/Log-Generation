package main

import (
	"log"
	"os"
	"time"

	"github.com/peekknuf/Log-Generation/config"
	"github.com/peekknuf/Log-Generation/internal/generator"
	"github.com/peekknuf/Log-Generation/internal/utils"
)

func main() {
	// Load configuration
	cfg := config.DefaultConfig()

	// Ensure logs directory exists
	if err := os.MkdirAll(cfg.LogDirectory, 0755); err != nil {
		log.Fatalf("Failed to create log directory: %v", err)
	}

	// Define the log file name
	logFile := "logs" + cfg.LogFormat

	// Create a ticker to generate logs at the specified interval
	ticker := time.NewTicker(cfg.LogInterval)
	defer ticker.Stop()

	// Log generation loop using for range
	for range ticker.C {
		// Generate a log entry
		logEntry := generator.GenerateLog()
		logData := logEntry.String()

		// Write the log entry to the file
		if err := utils.WriteLog(cfg.LogDirectory, logFile, logData); err != nil {
			log.Printf("Error writing log: %v", err)
		} else {
			log.Println("Generated log:", logData)
		}
	}
}
