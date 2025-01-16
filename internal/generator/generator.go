package generator

import (
	"fmt"
	"time"

	"github.com/peekknuf/Log-Generation/internal/utils"
)

type LogEntry struct {
	Timestamp           string
	ActivityID          string
	EventType           string
	UserID              string
	Username            string
	IPAddress           string
	Country             string
	Region              string
	City                string
	Coordinates         string
	OS                  string
	Browser             string
	DeviceType          string
	Action              string
	Status              string
	FailedAttempts      int
	LastSuccessfulLogin string
	SessionID           string
	RequestID           string
	TraceID             string
}

func GenerateLog() LogEntry {
	now := time.Now()
	action := utils.RandomAction()
	status := utils.RandomStatus()

	// Only include "failed_attempts" and "last_successful_login" for login attempts
	var failedAttempts int
	var lastSuccessfulLogin string
	if action == "LOGIN_ATTEMPT" {
		failedAttempts = utils.RandomInt(1, 10)
		lastSuccessfulLogin = utils.RandomPastTimestamp(now)
	} else {
		failedAttempts = 0
		lastSuccessfulLogin = ""
	}

	return LogEntry{
		Timestamp:           now.Format(time.RFC3339Nano), // Include nanoseconds for precision
		ActivityID:          fmt.Sprintf("ACT-%s-%03d", now.Format("20060102T150405"), utils.RandomInt(1, 999)),
		EventType:           "SUSPICIOUS_LOGIN_ATTEMPT", // You can make this dynamic too if needed
		UserID:              fmt.Sprintf("u-%06d", utils.RandomInt(1, 999999)),
		IPAddress:           utils.RandomIP(),
		Coordinates:         utils.RandomCoordinates(),
		OS:                  utils.RandomOS(),
		Browser:             utils.RandomBrowser(),
		DeviceType:          utils.RandomDeviceType(),
		Action:              action,
		Status:              status,
		FailedAttempts:      failedAttempts,
		LastSuccessfulLogin: lastSuccessfulLogin,
		SessionID:           fmt.Sprintf("sess-%06d", utils.RandomInt(1, 999999)),
		RequestID:           fmt.Sprintf("req-%09d", utils.RandomInt(1, 999999999)),
		TraceID:             fmt.Sprintf("trace-%s", utils.RandomHex(12)),
	}
}

func (l LogEntry) String() string {
	logEntry := fmt.Sprintf(
		"%s | %s | %s | user_id=%s username=%s ip_address=%s country=%s region=%s city=%s coordinates=%s os=%s browser=%s device_type=%s | action=%s status=%s",
		l.Timestamp,
		l.ActivityID,
		l.EventType,
		l.UserID,
		l.Username,
		l.IPAddress,
		l.Country,
		l.Region,
		l.City,
		l.Coordinates,
		l.OS,
		l.Browser,
		l.DeviceType,
		l.Action,
		l.Status,
	)

	// Only include "failed_attempts" and "last_successful_login" for login attempts
	if l.Action == "LOGIN_ATTEMPT" {
		logEntry += fmt.Sprintf(" failed_attempts=%d last_successful_login=%s", l.FailedAttempts, l.LastSuccessfulLogin)
	}

	logEntry += fmt.Sprintf(" | session_id=%s request_id=%s trace_id=%s", l.SessionID, l.RequestID, l.TraceID)

	return logEntry
}
