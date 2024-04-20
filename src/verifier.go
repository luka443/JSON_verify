package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// Function to read data from a file
func readFromFile(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err) // Stops the program in case of file read error
		return "", err
	}
	return string(content), nil
}

// Function to validate JSON data
func validateJsonData(jsonStr string) (bool, error) {
	var jsonData Data
	err := json.Unmarshal([]byte(jsonStr), &jsonData)
	if err != nil {
		return false, err
	}

	if jsonData.PolicyName == "" || jsonData.PolicyDocument.Version == "" {
		return false, fmt.Errorf("required fields: PolicyName or PolicyDocument.Version")
	}

	if len(jsonData.PolicyDocument.Statement) == 0 {
		return false, fmt.Errorf("required field: PolicyDocument.Statement")
	}

	if jsonData.PolicyDocument.Statement[0].Resource == "" {
		return false, fmt.Errorf("required field: PolicyDocument.Statement[0].Resource")
	}

	return !(jsonData.PolicyDocument.Statement[0].Resource == "pppp"), nil
}

// Data struct to hold JSON data

type Data struct {
	PolicyName     string `json:"PolicyName"`
	PolicyDocument struct {
		Version   string `json:"Version"`
		Statement []struct {
			Sid      string   `json:"Sid"`
			Effect   string   `json:"Effect"`
			Action   []string `json:"Action"`
			Resource string   `json:"Resource"`
		} `json:"Statement"`
	} `json:"PolicyDocument"`
}
