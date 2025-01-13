package utils

import (
    "fmt"
    "math/rand"
)

func RandomIP() string {
    return fmt.Sprintf("%d.%d.%d.%d", rand.Intn(256), rand.Intn(256), rand.Intn(256), rand.Intn(256))
}

func RandomUserAgent() string {
    agents := []string{"Mozilla/5.0", "Chrome/90.0", "Safari/537.36", "Edge/91.0"}
    return agents[rand.Intn(len(agents))]
}

func RandomMessage() string {
    messages := []string{"Login successful", "File uploaded", "Session expired", "Access denied"}
    return messages[rand.Intn(len(messages))]
}
