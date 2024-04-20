package main

import (
	"testing"
)

func TestValidateJsonStr(t *testing.T) {

	testW, _ := readFromFile("../TESTS/test1.json")
	_, err := validateJsonData(testW)

	if err == nil {
		t.Errorf("Expected invalid json")
	}

	testTrueStr, _ := readFromFile("../TESTS/test2.json")
	testTrue, _ := validateJsonData(testTrueStr)

	if !testTrue {
		t.Errorf("Expected: true but gotfalse")
	}

}
