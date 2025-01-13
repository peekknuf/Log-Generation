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
    cfg := config.DefaultConfig()

    // Ensure logs directory exists
    if err := os.MkdirAll(cfg.LogDirectory, 0755); err != nil {
        log.Fatalf("Failed to create log directory: %v", err)
    }

    logFile := "logs" + cfg.LogFormat

    ticker := time.NewTicker(cfg.LogInterval)
    defer ticker.Stop()

    for {
        select {
        case <-ticker.C:
            logEntry := generator.GenerateLog()
            logData := logEntry.String()

            if err := utils.WriteLog(cfg.LogDirectory, logFile, logData); err != nil {
                log.Printf("Error writing log: %v", err)
            } else {
                log.Println("Generated log:", logData)
            }
        }
    }
}

