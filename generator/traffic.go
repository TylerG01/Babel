package generator

import (
	"math/rand"
)

// Generates a random session duration (1-300 seconds)
func GenerateRandomDuration() int {
	return rand.Intn(300) + 1
}

// Generates a random amount of sent bytes (500-50000)
func GenerateRandomSentBytes() int {
	return rand.Intn(49501) + 500
}

// Generates a random amount of received bytes (5000-500000)
func GenerateRandomReceivedBytes() int {
	return rand.Intn(495001) + 5000
}

// Generates a random count of sent packets (1-1000)
func GenerateRandomSentPackets() int {
	return rand.Intn(1000) + 1
}

// Generates a random count of received packets (1-2000)
func GenerateRandomReceivedPackets() int {
	return rand.Intn(2000) + 1
}

// Generates a random number of detected applications (1-10)
func GenerateRandomAppCount() int {
	return rand.Intn(10) + 1
}
