package generator

import (
	"fmt"
	"math/rand"
	"strconv"
)

// Predefined srouce interface names, roles and countries
var srcInterfaces = []string{"port1", "port2", "port3", "port12", "eth0", "wlan0"}
var srcIntfRoles = []string{"undefined", "trusted", "untrusted"}
var srccountries = []string{"United States", "Canada", "Germany", "United Kingdom", "France", "Japan", "Australia",
	"South Korea", "Italy", "Spain", "Netherlands", "Sweden", "Switzerland", "Norway", "Denmark", "Finland", "Belgium",
	"Austria", "New Zealand", "Singapore", "Ireland", "Portugal", "Luxembourg", "Iceland", "Israel", "Russia", "Bulgaria",
	"Ukraine", "Romania", "China", "Taiwan", "North Korea", "Thailand"}

// Generates a random IPv4 address
func GenerateRandomIP() string {
	return fmt.Sprintf("%d.%d.%d.%d",
		rand.Intn(256), rand.Intn(256), rand.Intn(256), rand.Intn(256))
}

// Generatess a hostname like "pc13", "sensor25", "server999"
func GenerateRandomSourceName() string {
	names := []string{"pc", "laptop", "server", "host"} // List of names to pick from at random
	randomNum := rand.Intn(999) + 1                     // Random number between 1 and 999
	return fmt.Sprintf("%s%d", names[rand.Intn(len(names))], randomNum)
}

// Function to generate a random port number within range
func GenerateRandomPort() string {
	port := rand.Intn(65536) + 1
	return strconv.Itoa(port)
}

// Generatess a random MAC address
func GenerateRandomMAC() string {
	mac := []byte{
		0x00, 0x16, 0x3e, // Vendor prefix can be randomized here
		byte(rand.Intn(256)),
		byte(rand.Intn(256)),
		byte(rand.Intn(256)),
	}
	return fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x", mac[0], mac[1], mac[2], mac[3], mac[4], mac[5])
}

// Selects a random country from srccountries variable
func GenerateRandomCountry() string {
	return srccountries[rand.Intn(len(srccountries))]
}

// Selects a random source interface from srcInterface variable
func GenerateRandomSourceInterface() string {
	return srcInterfaces[rand.Intn(len(srcInterfaces))]
}

// Selects a random source role from srcIntfRoles variable
func GenerateRandomSourceRole() string {
	return srcIntfRoles[rand.Intn(len(srcIntfRoles))]
}
