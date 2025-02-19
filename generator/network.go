package generator

import (
	"fmt"
	"math/rand"
)

// Predefined network protocols
var protocols = []int{6, 17, 1}                                   // 6 = TCP, 17 = UDP, 1 = ICMP
var actions = []string{"start", "deny", "close"}                  // Firewall actions
var policyTypes = []string{"policy", "firewall", "vpn", "proxy"}  // Policy types
var translations = []string{"snat", "dnat", "none"}               //
var firewallpolices = []string{"learn", "monitor", "enforce"}     // Short list of firewall policy types
var servicetype = []string{"HTTP", "HTTPS", "TCP", "UDP", "SMTP"} // Short list of network protocl service types

// Selects a random network protocol (TCP, UDP, ICMP) from protocols variable
func GenerateRandomProtocol() int {
	return protocols[rand.Intn(len(protocols))]
}

// Selects a random string from actions variable
func GenerateRandomAction() string {
	return actions[rand.Intn(len(actions))]
}

// Generates a random firewall policy ID number between 1-999
func GenerateRandomPolicyID() int {
	return rand.Intn(999) + 1
}

// Selects a random policy type from politcyTypes variable
func GenerateRandomPolicyType() string {
	return policyTypes[rand.Intn(len(policyTypes))]
}

// Selects a random transaltion from translations variable
func GenerateRandomTranslationType() string {
	return translations[rand.Intn(len(translations))]
}

// Generates a random translated IP (for SNAT/DNAT cases)
func GenerateRandomTranslatedIP() string {
	return fmt.Sprintf("172.%d.%d.%d",
		rand.Intn(256), rand.Intn(256), rand.Intn(256))
}

// Generates a NAT Translation (transport) port (valid range: 1024-65535)
func GenerateRandomTranslatedPort() int {
	return rand.Intn(64511) + 1024
}

// Generates a random firewall policy from list
func GenerateRandomFirewallPolicyMode() string {
	return firewallpolices[rand.Intn(len(firewallpolices))]
}

// Generates a radnom network protocol from list
func GenerateRandomServiceType() string {
	return servicetype[rand.Intn(len(servicetype))]
}
