package gatha

import (
	"encoding/json"
	"os"
	"testing"
)

// TestVerseJsonMarshaling tests the JSON marshaling and unmarshaling of Verse.
func TestVerseJsonMarshaling(t *testing.T) {
	originalText := "The quick brown fox jumps over the lazy dog\nand in to the box, " +
		"right next to the log.\n\nHer coat of reddish hues,\nthe most lovely of views."
	line := NewVerse(originalText)

	// Marshal to JSON
	jsonData, err := json.Marshal(line)
	if err != nil {
		t.Fatalf("Failed to marshal Verse to JSON: %v", err)
	}

	expectedJson, err := os.ReadFile("testdata/expected_verse.json")
	if err != nil {
		t.Fatalf("Failed to read test artifact file: %v", err)
	}

	if string(jsonData) != string(expectedJson) {
		t.Errorf("Expected JSON %s, got %s", expectedJson, string(jsonData))
	}

	// Unmarshal from JSON
	var unmarshaledVerse Verse
	err = json.Unmarshal(jsonData, &unmarshaledVerse)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON to Verse: %v", err)
	}

	if unmarshaledVerse.toString() != originalText {
		t.Errorf("Expected Verse value %s, got %s", originalText, unmarshaledVerse.toString())
	}
}
