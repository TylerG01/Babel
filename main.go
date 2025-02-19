package main

import (
	"flag"
	"fmt"
	"babel/generator"
	"babel/output"
	"os"
	"path/filepath"
)

// LogEntry struct to store all generated log fields
type LogEntry struct {
	//Metadata Information
	Date      string `json:"date" csv:"date"`
	Time      string `json:"time" csv:"time"`
	LogID     string `json:"logid" csv:"logid"`
	EventTime int64  `json:"eventtime" csv:"eventtime"`
	Level     string `json:"level" csv:"level"`
	SubType   string `json:"subtype" csv:"subtype"`
	VD        string `json:"vd" csv:"vd"`
	// Source Information
	SrcIP        string `json:"srcip" csv:"srcip"`
	SrcName      string `json:"srcname" csv:"srcname"`
	SrcPort      string `json:"srcport" csv:"srcport"`
	SrcMAC       string `json:"srcmac" csv:"srcmac"`
	SrcCountry   string `json:"scrcountry" csv:"srccountry"`
	MasterSrcMac string `json:"mastersrcmac" csv:"mastersrcmac"`
	SrcIntf      string `json:"srcintf" csv:"srcintf"`
	SrcIntfRole  string `json:"srcintfrole" csv:"srcintfrole"`
	// Destination Information
	DstIP       string `json:"dstip" csv:"dstip"`
	DstName     string `json:"dstname" csv:"dstname"`
	DstPort     int    `json:"dstport" csv:"dstport"`
	DstIntf     string `json:"dstintf" csv:"dstintf"`
	DstMAC      string `json:"dstmac" csv:"dstmac"`
	DstIntfRole string `json:"dstintfrole" csv:"dstinftrole"`
	DstCountry  string `json:"dstcountry" csv:"dstcountry"`
	// Network & Policy Information
	Protocol       int    `json:"proto" csv:"proto"`
	Action         string `json:"action" csv:"action"`
	PolicyID       int    `json:"policyid" csv:"policyid"`
	PolicyType     string `json:"policytype" csv:"policytype"`
	Translation    string `json:"trandisp" csv:"trandisp"`
	TranslatedIP   string `json:"transip" csv:"transip"`
	TranslatedPort int    `json:"transport" csv:"transport"`
	PolicyMode     string `json:"policymode" csv:"policymode"`
	Service        string `json:"service" csv:"service"`
	// NAT & Translation Information
	NATType string `json:"nat_type" csv:"nat_type"`
	NATIP   string `json:"nat_ip" csv:"nat_ip"`
	NATPort int    `json:"nat_port" csv:"nat_port"`
	// Application & Security Information
	AppID       int    `json:"appid" csv:"appid"`
	AppName     string `json:"app" csv:"app"`
	AppCategory string `json:"appcat" csv:"appcat"`
	AppRisk     string `json:"apprisk" csv:"apprisk"`
	UTMAction   string `json:"utmaction" csv:"utmaction"`
	UTMRef      string `json:"utmref" csv:"utmref"`
	// Traffic Statistics
	Duration        int `json:"duration" csv:"duration"`
	SentBytes       int `json:"sentbyte" csv:"sentbyte"`
	ReceivedBytes   int `json:"rcvdbyte" csv:"rcvdbyte"`
	SentPackets     int `json:"sentpkt" csv:"sentpkt"`
	ReceivedPackets int `json:"rcvdpkt" csv:"rcvdpkt"`
	AppCount        int `json:"countapp" csv:"countapp"`
	// Device Information
	DeviceType string `json:"devtype" csv:"devtype"`
	OSName     string `json:"osname" csv:"osname"`
}

// generateLogEntry creates a new log entry with randomly generated metadata
func generateLogEntry() LogEntry {
	deviceType := generator.GenerateRandomDeviceType()
	return LogEntry{
		// Metadata information
		Date:      generator.GenerateCurrentDate(),
		Time:      generator.GenerateRandomTime(),
		LogID:     generator.GenerateLogID(),
		EventTime: generator.GenerateEventTime(),
		Level:     generator.GenerateLogLevel(),
		SubType:   generator.GenerateSubType(),
		VD:        generator.GenerateRandomVD(),
		// Source Information
		SrcIP:        generator.GenerateRandomIP(),
		SrcName:      generator.GenerateRandomSourceName(),
		SrcPort:      generator.GenerateRandomPort(),
		SrcMAC:       generator.GenerateRandomMAC(),
		SrcCountry:   generator.GenerateRandomCountry(),
		MasterSrcMac: generator.GenerateRandomMAC(),
		SrcIntf:      generator.GenerateRandomSourceInterface(),
		SrcIntfRole:  generator.GenerateRandomSourceRole(),
		// Destination Information
		DstIP:       generator.GenerateRandomDestinationIP(),
		DstMAC:      generator.GenerateRandomDestinationMAC(),
		DstPort:     generator.GenerateRandomDestinationPort(),
		DstName:     generator.GenerateRandomDestinationName(),
		DstIntf:     generator.GenerateRandomDestinationInterface(),
		DstIntfRole: generator.GenerateDestingationRole(),
		DstCountry:  generator.GenerateRandomDestinationCountry(),
		// Network & Policy Information
		Protocol:       generator.GenerateRandomProtocol(),
		Action:         generator.GenerateRandomAction(),
		PolicyID:       generator.GenerateRandomPolicyID(),
		PolicyType:     generator.GenerateRandomPolicyType(),
		Translation:    generator.GenerateRandomTranslationType(),
		TranslatedIP:   generator.GenerateRandomTranslatedIP(),
		TranslatedPort: generator.GenerateRandomTranslatedPort(),
		PolicyMode:     generator.GenerateRandomFirewallPolicyMode(),
		Service:        generator.GenerateRandomServiceType(),
		// NAT & Translation Information
		NATType: generator.GenerateRandomNATType(),
		NATIP:   generator.GenerateRandomNATIP(),
		NATPort: generator.GenerateRandomNATPort(),
		// Application & Security Information
		AppID:       generator.GenerateRandomAppID(),
		AppName:     generator.GenerateRandomAppName(),
		AppCategory: generator.GenerateRandomAppCategory(),
		AppRisk:     generator.GenerateRandomAppRisk(),
		UTMAction:   generator.GenerateRandomUTMAction(),
		UTMRef:      generator.GenerateRandomUTMReference(),
		// Traffic Statistics
		Duration:        generator.GenerateRandomDuration(),
		SentBytes:       generator.GenerateRandomSentBytes(),
		ReceivedBytes:   generator.GenerateRandomReceivedBytes(),
		SentPackets:     generator.GenerateRandomSentPackets(),
		ReceivedPackets: generator.GenerateRandomReceivedPackets(),
		AppCount:        generator.GenerateRandomAppCount(),
		// Device Information
		DeviceType: deviceType,
		OSName:     generator.GenerateRandomOSName(deviceType),
	}
}

func printHelp() {
	fmt.Println("Usage: Babel [OPTIONS]")
	fmt.Println("\nOptions:")
	fmt.Println("  -n       <int>    Number of log events to generate (default: 10)")
	fmt.Println("  -fmt  <string> Output format: csv, json, raw (default: csv)")
	fmt.Println("  -f    <string> Output filename for CSV (default: logs.csv)")
	// Rate limiting currently in development
	// 	fmt.Println("  -r    <int>    Maximum events per second (0 = unlimited, default: 0)")
	fmt.Println("  -vB      <int>    Maximum volume in bytes before stopping (overrides -n)")
	fmt.Println("  -vT      <float>  Maximum volume in terabytes before stopping (overrides -n)")
	fmt.Println("  --help            Show this help message and exit")
	fmt.Println("\nExamples:")
	fmt.Println("  ./babel -n 100 -f json")
	fmt.Println("  ./babel -vB 1000000 -r 5 -f csv")
	fmt.Println("  ./babel -vT 1 -f raw")
	os.Exit(0)
}

func main() {
	// Define command-line flags
	numLogs := flag.Int("n", 0, "Number of log events to generate (ignored if -vB or -vT is set)")
	outputFormat := flag.String("fmt", "csv", "Output format: csv, json, raw")
	outputFile := flag.String("f", "logs.csv", "Output filename for CSV")
	outputDir := flag.String("dir", "", "Directory to store JSON files (for -format json)")
	volumeBytes := flag.Int64("vB", 0, "Maximum volume in bytes before stopping (overrides -n)")
	volumeTB := flag.Float64("vT", 0, "Maximum volume in terabytes before stopping (overrides -n)")
	helpFlag := flag.Bool("help", false, "Show help message")

	// Parse flags
	flag.Parse()

	// If --help is provided, display help message
	if *helpFlag {
		printHelp()
	}

	// Convert terabytes to bytes if -vT is used
	if *volumeTB > 0 {
		*volumeBytes = int64(*volumeTB * 1e12) // 1TB = 10^12 bytes
	}

	// Calculate bytes per log based on assumption: 1000 logs ≈ 10,000 bytes
	bytesPerLog := 10.0
	maxLogs := *numLogs

	if *volumeBytes > 0 {
		maxLogs = int(float64(*volumeBytes) / bytesPerLog)
	}

	// Generate logs
	var logs []interface{}
	for i := 0; i < maxLogs; i++ {
		logEntry := generateLogEntry()
		logs = append(logs, logEntry)

		// Save individual JSON files if JSON mode & directory specified
		if *outputFormat == "json" && *outputDir != "" {
			filename := filepath.Join(*outputDir, fmt.Sprintf("log_%04d.json", i+1))
			output.WriteLogAsJSONFile(filename, logEntry)
		}
	}

	// Handle other output formats
	if *outputFormat == "csv" {
		output.WriteLogAsCSV(*outputFile, logs)
		fmt.Printf("\nLog generation complete. Logs saved to %s\n", *outputFile)
	} else if *outputFormat == "json" && *outputDir == "" {
		fmt.Println("\nGenerated Log Entries in JSON:")
		for _, log := range logs {
			output.PrintLogAsJSON(log)
		}
	} else if *outputFormat == "raw" {
		fmt.Println("\nGenerated Log Entries in RAW format:")
		for _, log := range logs {
			output.PrintLogAsRaw(log)
		}
	}
}
