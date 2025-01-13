package utils

import (
    "os"
    "path/filepath"
)

func WriteLog(directory string, fileName string, logData string) error {
    path := filepath.Join(directory, fileName)
    file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return err
    }
    defer file.Close()

    _, err = file.WriteString(logData + "\n")
    return err
}

