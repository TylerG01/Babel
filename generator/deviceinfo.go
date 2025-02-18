package generator

import (
	"math/rand"
)

// Map of device types to corresponding OS options
var osOptions = map[string][]string{
	"Linux PC":       {"Arch", "Debian", "Ubuntu", "Mint", "Manjaro"},
	"Linux Server":   {"Arch Linux", "CentOS", "ClearOS", "Debian", "Oracle", "RHEL", "Slackware Linux", "SUSE", "Ubuntu"},
	"Windows PC":     {"Windows 7", "Windows 8", "Windows 10", "Windows 11", "Windows Vista"},
	"Windows Server": {"Windows Server 2008", "Windows Server 2012", "Windows Server 2016", "Windows Server 2019"},
	"Mac":            {"macOS Monterey", "macOS Big Sur", "macOS Ventura", "macOS Catalina"},
}

// List of device types
var deviceTypes = []string{"Linux PC", "Windows PC", "Linux Server", "Windows Server", "Mac"}

// Selects a random device type from deviceTypes variable above
func GenerateRandomDeviceType() string {
	return deviceTypes[rand.Intn(len(deviceTypes))]
}

// Returns a random OS based on device type
func GenerateRandomOSName(deviceType string) string {
	if osList, exists := osOptions[deviceType]; exists {
		return osList[rand.Intn(len(osList))] // Pick a random OS from the list
	}
	return "Unknown OS"
}
