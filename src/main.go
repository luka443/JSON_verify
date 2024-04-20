package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fileName := flag.String("file", "", "flag file")
	flag.Parse()

	if *fileName == "" {
		fmt.Println("Usage: ./parser --file FILE")
		os.Exit(1)
	}

	validation, err := readAndValidateFromFile(*fileName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println(validation)
}

func readAndValidateFromFile(fileName string) (bool, error) {
	jsonData, err := readFromFile(fileName)
	if err != nil {
		return false, fmt.Errorf("error reading from file: %v", err)
	}

	validation, err := validateJsonData(jsonData)
	if err != nil {
		return false, fmt.Errorf("error validating json: %v", err)
	}

	return validation, nil
}
