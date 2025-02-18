package generator

import (
	"fmt"
	"math/rand"
)

// Predefined application categories
var appCategories = []string{"Web.Client", "Email", "Social Media", "Gaming", "Business", "Security", "Streaming"} // Predefined application categories
var appNames = []string{"HTTPS.BROWSER", "Gmail", "Facebook", "FortiGuard", "YouTube", "Zoom", "Slack"}            // Predefined application names
var appRisks = []string{"low", "medium", "high", "critical"}                                                       // Predefined application risks
var utmActions = []string{"allow", "block", "monitor", "quarantine"}                                               // Predefined UTM actions

// Generates a random App ID (10000-99999)
func GenerateRandomAppID() int {
	return rand.Intn(90000) + 10000
}

// Geneates a random application name based on options in appName variable
func GenerateRandomAppName() string {
	return appNames[rand.Intn(len(appNames))]
}

// Generates a random application category based on options in appCategories variable
func GenerateRandomAppCategory() string {
	return appCategories[rand.Intn(len(appCategories))]
}

// Generates a random application risk level based on options in appRisks variable
func GenerateRandomAppRisk() string {
	return appRisks[rand.Intn(len(appRisks))]
}

// Generates a random UTM Action based on options in utmActions variable
func GenerateRandomUTMAction() string {
	return utmActions[rand.Intn(len(utmActions))]
}

// Generates returns a random UTM reference ID int (0-999999)
func GenerateRandomUTMReference() string {
	return fmt.Sprintf("%d-%06d", rand.Intn(10), rand.Intn(1000000))
}
