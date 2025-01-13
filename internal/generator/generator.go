package generator

import (
    "time"
    "log_gen/internal/utils"
)

type LogEntry struct {
    Timestamp string
    IP        string
    UserAgent string
    Message   string
}

func GenerateLog() LogEntry {
    return LogEntry{
        Timestamp: time.Now().Format(time.RFC3339),
        IP:        utils.RandomIP(),
        UserAgent: utils.RandomUserAgent(),
        Message:   utils.RandomMessage(),
    }
}

func (l LogEntry) String() string {
    return "[" + l.Timestamp + "] IP: " + l.IP + ", UserAgent: " + l.UserAgent + ", Message: " + l.Message
}

