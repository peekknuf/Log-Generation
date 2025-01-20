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

	// Metrics tracking variables
	var recordCount int64 = 0
	var totalBytesWritten int64 = 0
	startTime := time.Now()

	// Periodic reporting ticker (e.g., every 5 seconds)
	reportTicker := time.NewTicker(1 * time.Second)
	defer reportTicker.Stop()

	logger := log.New(os.Stdout, "", log.LstdFlags)
	metricsLogger := log.New(os.Stdout, "METRICS: ", log.LstdFlags)

	// Log generation loop using for-select to handle multiple tickers
	for {
		select {
		case <-ticker.C:
			logEntry := generator.GenerateLog()
			logData := logEntry.String()
			err := utils.WriteLog(cfg.LogDirectory, logFile, logData)
			if err != nil {
				logger.Printf("Error writing log: %v", err)
			} else {
				recordCount++
				totalBytesWritten += int64(len(logData))
			}
		case <-reportTicker.C:
			elapsedTime := time.Since(startTime).Seconds()
			rps := float64(recordCount) / elapsedTime
			mbWritten := float64(totalBytesWritten) / 1024 / 1024
			metricsLogger.Printf("RPS: %.2f | Total Records: %d | Data Written: %.2f MB",
				rps, recordCount, mbWritten)
		}
	}
}
