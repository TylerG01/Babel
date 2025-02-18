package generator

import (
	"fmt"
	"math/rand"
)

// Predefined variables
var dstInterfaces = []string{"port10", "port11", "port12", "eth1", "wlan1"} // Predefined destination interfaces
// Predefined desitnation countries
var dstCountries = []string{"United States", "Canada", "Germany", "United Kingdom", "France", "Japan", "Australia",
	"South Korea", "Italy", "Spain", "Netherlands", "Sweden", "Switzerland", "Norway", "Denmark", "Finland", "Belgium",
	"Austria", "New Zealand", "Singapore", "Ireland", "Portugal", "Luxembourg", "Iceland", "Israel", "Russia", "Bulgaria",
							"Ukraine", "Romania", "China", "Taiwan", "North Korea", "Thailand"}
var dstintfrole = []string{"underfined", "unknown"} // Predefined destination role

// Generates a random IPv4 address
func GenerateRandomDestinationIP() string {
	return fmt.Sprintf("%d.%d.%d.%d",
		rand.Intn(256), rand.Intn(256), rand.Intn(256), rand.Intn(256))
}

// Generates a random MAC address
func GenerateRandomDestinationMAC() string {
	mac := []byte{
		0x00, 0x16, 0x3e, // Vendor prefix (can be randomized)
		byte(rand.Intn(256)),
		byte(rand.Intn(256)),
		byte(rand.Intn(256)),
	}
	return fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x", mac[0], mac[1], mac[2], mac[3], mac[4], mac[5])
}

// Generates a random port (1024-65535 for dynamic ports)
func GenerateRandomDestinationPort() int {
	return rand.Intn(64511) + 1024 // 1024 to 65535
}

// Selects a random destination interface name from dstInterfaces varaiable above
func GenerateRandomDestinationInterface() string {
	return dstInterfaces[rand.Intn(len(dstInterfaces))]
}

// Selects a random country from dstCountries variable above
func GenerateRandomDestinationCountry() string {
	return dstCountries[rand.Intn(len(dstCountries))]
}

// Creates a hostname like "server55", "web42"
func GenerateRandomDestinationName() string {
	names := []string{"server", "database", "web", "api", "proxy", "gateway", "firewall"}
	randomNum := rand.Intn(999) + 1 // Random number between 1 and 999
	return fmt.Sprintf("%s%d", names[rand.Intn(len(names))], randomNum)
}

// Selects a destination role from dstinfrole variable above
func GenerateDestingationRole() string {
	return dstintfrole[rand.Intn(len(dstintfrole))]
}
