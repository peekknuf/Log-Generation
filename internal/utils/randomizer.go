package utils

import (
	"fmt"
	"math/rand"
	"time"
)

// RandomIP generates a random IP address.
func RandomIP() string {
	return fmt.Sprintf("%d.%d.%d.%d", rand.Intn(256), rand.Intn(256), rand.Intn(256), rand.Intn(256))
}

// RandomUserAgent generates a random user agent string.
func RandomUserAgent() string {
	agents := []string{"Mozilla/5.0", "Chrome/90.0", "Safari/537.36", "Edge/91.0"}
	return agents[rand.Intn(len(agents))]
}

// RandomMessage generates a random log message.
func RandomMessage() string {
	messages := []string{"Login successful", "File uploaded", "Session expired", "Access denied"}
	return messages[rand.Intn(len(messages))]
}

// RandomCoordinates generates random latitude and longitude.
func RandomCoordinates() string {
	lat := fmt.Sprintf("%.4f", 90-rand.Float64()*180)
	long := fmt.Sprintf("%.4f", 180-rand.Float64()*360)
	return fmt.Sprintf("%s,%s", lat, long)
}

// RandomOS generates a random operating system.
func RandomOS() string {
	osList := []string{"Windows 10", "macOS 12", "Linux", "iOS 15", "Android 11"}
	return osList[rand.Intn(len(osList))]
}

// RandomBrowser generates a random browser.
func RandomBrowser() string {
	browsers := []string{"Chrome 117.0.0.0", "Firefox 118.0.0.0", "Safari 16.0.0.0", "Edge 117.0.0.0"}
	return browsers[rand.Intn(len(browsers))]
}

// RandomDeviceType generates a random device type.
func RandomDeviceType() string {
	devices := []string{"Desktop", "Mobile", "Tablet"}
	return devices[rand.Intn(len(devices))]
}

// RandomPastTimestamp generates a random timestamp in the past.
func RandomPastTimestamp(now time.Time) string {
	past := now.Add(-time.Duration(rand.Intn(30)) * 24 * time.Hour)
	return past.Format(time.RFC3339Nano)
}

// RandomHex generates a random hex string of the given length.
func RandomHex(length int) string {
	const charset = "abcdef0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}

// RandomInt generates a random integer between min and max.
func RandomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

// RandomAction generates a random action.
func RandomAction() string {
	actions := []string{"LOGIN_ATTEMPT", "LOGOUT", "FILE_UPLOAD", "PASSWORD_CHANGE", "ACCOUNT_CREATION"}
	return actions[rand.Intn(len(actions))]
}

// RandomStatus generates a random status.
func RandomStatus() string {
	statuses := []string{"SUCCESS", "FAILED", "PENDING"}
	return statuses[rand.Intn(len(statuses))]
}
