package output

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
)

// Prints logs in JSON format
func PrintLogAsJSON(logEntry interface{}) {
	jsonData, err := json.MarshalIndent(logEntry, "", "  ")
	if err != nil {
		log.Fatal("Error generating JSON:", err)
	}
	fmt.Println(string(jsonData))
}

// Writes a single log entry to an individual JSON file
func WriteLogAsJSONFile(filename string, logEntry interface{}) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Error creating JSON file: %v", err)
	}
	defer file.Close()

	jsonData, err := json.MarshalIndent(logEntry, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling JSON data: %v", err)
	}

	_, err = file.Write(jsonData)
	if err != nil {
		log.Fatalf("Error writing JSON to file: %v", err)
	}
}

// Writes logs to a CSV file
func WriteLogAsCSV(filename string, logEntries []interface{}) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal("Error creating CSV file:", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if len(logEntries) == 0 {
		log.Fatal("No data to write to CSV.")
	}

	// Extract headers from struct field names
	headers := getStructFields(logEntries[0])
	writer.Write(headers)

	// Extract values for each log entry
	for _, entry := range logEntries {
		values := getStructValues(entry)
		writer.Write(values)
	}
	fmt.Println("CSV file created:", filename)
}

// Prints logs in RAW format (key=value format)
func PrintLogAsRaw(logEntry interface{}) {
	val := reflect.ValueOf(logEntry)
	typ := val.Type()

	var rawLog []string
	for i := 0; i < val.NumField(); i++ {
		fieldName := typ.Field(i).Tag.Get("json")                 // Get JSON field name
		fieldValue := fmt.Sprintf("%v", val.Field(i).Interface()) // Convert value to string
		rawLog = append(rawLog, fmt.Sprintf("%s=%s", fieldName, fieldValue))
	}
	fmt.Println(strings.Join(rawLog, " ")) // Print key=value pairs separated by spaces
}

// Extracts field names from a struct
func getStructFields(obj interface{}) []string {
	var fields []string
	val := reflect.ValueOf(obj)
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		fields = append(fields, typ.Field(i).Tag.Get("csv")) // Extract CSV tag
	}
	return fields
}

// Extracts field values from a struct as strings
func getStructValues(obj interface{}) []string {
	var values []string
	val := reflect.ValueOf(obj)

	for i := 0; i < val.NumField(); i++ {
		values = append(values, fmt.Sprintf("%v", val.Field(i).Interface()))
	}
	return values
}
