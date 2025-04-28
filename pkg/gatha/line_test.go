package gatha

import (
	"encoding/json"
	"os"
	"testing"
)

// TestWordJsonMarshaling tests the JSON marshaling and unmarshaling of Line.
func TestLineJsonMarshaling(t *testing.T) {
	originalText := "The quick brown fox jumps over the lazy dog"
	line := NewLine(originalText)

	// Marshal to JSON
	jsonData, err := json.Marshal(line)
	if err != nil {
		t.Fatalf("Failed to marshal Line to JSON: %v", err)
	}

	expectedJson, err := os.ReadFile("testdata/expected_line.json")
	if err != nil {
		t.Fatalf("Failed to read test artifact file: %v", err)
	}

	if string(jsonData) != string(expectedJson) {
		t.Errorf("Expected JSON %s, got %s", expectedJson, string(jsonData))
	}

	// Unmarshal from JSON
	var unmarshaledLine Line
	err = json.Unmarshal(jsonData, &unmarshaledLine)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON to Line: %v", err)
	}

	if unmarshaledLine.toString() != originalText {
		t.Errorf("Expected Line value %s, got %s", originalText, unmarshaledLine.toString())
	}
}
