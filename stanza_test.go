package gatha

import (
	"encoding/json"
	"os"
	"testing"
)

// TestStanzaJsonMarshaling tests the JSON marshaling and unmarshaling of Stanza.
func TestStanzaJsonMarshaling(t *testing.T) {
	originalText := "The quick brown fox jumps over the lazy dog\nand in to the box, right next to the log."
	line := NewStanza(originalText)

	// Marshal to JSON
	jsonData, err := json.Marshal(line)
	if err != nil {
		t.Fatalf("Failed to marshal Stanza to JSON: %v", err)
	}

	expectedJson, err := os.ReadFile("test/expected_stanza.json")
	if err != nil {
		t.Fatalf("Failed to read test artifact file: %v", err)
	}

	if string(jsonData) != string(expectedJson) {
		t.Errorf("Expected JSON %s, got %s", expectedJson, string(jsonData))
	}

	// Unmarshal from JSON
	var unmarshaledStanza Stanza
	err = json.Unmarshal(jsonData, &unmarshaledStanza)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON to Stanza: %v", err)
	}

	if unmarshaledStanza.toString() != originalText {
		t.Errorf("Expected Stanza value %s, got %s", originalText, unmarshaledStanza.toString())
	}
}
