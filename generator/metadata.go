package generator

import (
	"fmt"
	"math/rand"
	"time"
)

// Predefine metadate variables
var logLevels = []string{"notice", "warning", "critical", "alert"}      // Predefined log severity levels
var logSubTypes = []string{"forward", "local-in", "multicast", "sniff"} // Predefined log sub-types
var logIDCounter int = 1                                                // Global counter for log IDs

// Generate the current date in YYYY-MM-DD format
func GenerateCurrentDate() string {
	return time.Now().Format("2006-01-02")
}

// Returns the current time with randomly adjusted seconds
func GenerateRandomTime() string {
	currentTime := time.Now()
	randomSeconds := rand.Intn(121) - 60 // Range: -60 to +60 seconds
	randomTime := currentTime.Add(time.Duration(randomSeconds) * time.Second)
	return randomTime.Format("15:04:05")
}

// Generates a sequential 10-digit log ID
func GenerateLogID() string {
	logID := fmt.Sprintf("%010d", logIDCounter)
	logIDCounter++ // Increment counter for next log
	return logID
}

// Generates a UNIX timestamp for event time
func GenerateEventTime() int64 {
	return time.Now().Unix()
}

// Selects a random log level from predefined values in the log level variable
func GenerateLogLevel() string {
	return logLevels[rand.Intn(len(logLevels))]
}

// Selects a random log level from predefined values in the log sub-type variable
func GenerateSubType() string {
	return logSubTypes[rand.Intn(len(logSubTypes))]
}

// Generates a random Virtual Domain (vdom1 through vdom25)
func GenerateRandomVD() string {
	return fmt.Sprintf("vdom%d", rand.Intn(25)+1)
}
