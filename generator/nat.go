package generator

import (
	"fmt"
	"math/rand"
)

// Predefined NAT types
var natTypes = []string{"snat", "dnat", "none"}

// Seelects are NAT type at randome from the natTypes variable
func GenerateRandomNATType() string {
	return natTypes[rand.Intn(len(natTypes))]
}

// Generates a random NAT-translated IP
func GenerateRandomNATIP() string {
	return fmt.Sprintf("172.%d.%d.%d", rand.Intn(256), rand.Intn(256), rand.Intn(256))
}

// Generates a random translated port (valid range: 1024-65535)
func GenerateRandomNATPort() int {
	return rand.Intn(64511) + 1024
}
