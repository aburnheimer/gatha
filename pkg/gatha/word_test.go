package gatha

import (
	"encoding/json"
	"testing"
)

// TestWordJsonMarshaling tests the JSON marshaling and unmarshaling of Word.
func TestWordJsonMarshaling(t *testing.T) {
	word := NewWord("Cellar")

	// Marshal to JSON
	jsonData, err := json.Marshal(word)
	if err != nil {
		t.Fatalf("Failed to marshal Word to JSON: %v", err)
	}

	expectedJson := `{"value":"Cellar"}`
	if string(jsonData) != expectedJson {
		t.Errorf("Expected JSON %s, got %s", expectedJson, string(jsonData))
	}

	// Unmarshal from JSON
	var unmarshaledWord Word
	err = json.Unmarshal(jsonData, &unmarshaledWord)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON to Word: %v", err)
	}

	if unmarshaledWord.Value != word.Value {
		t.Errorf("Expected Word value %s, got %s", word.Value, unmarshaledWord.Value)
	}
}

// TestRhymeWordJsonMarshaling tests the JSON marshaling and unmarshaling of RhymeWord.
func TestRhymeWordJsonMarshaling(t *testing.T) {
	rhymeWord := NewRhymeWord("Door")
	rhymeWord.SetRhyme("A1")

	// Marshal to JSON
	jsonData, err := json.Marshal(rhymeWord)
	if err != nil {
		t.Fatalf("Failed to marshal RhymeWord to JSON: %v", err)
	}

	expectedJson := `{"value":"Door","rhyme":"A1"}`
	if string(jsonData) != expectedJson {
		t.Errorf("Expected JSON %s, got %s", expectedJson, string(jsonData))
	}

	// Unmarshal from JSON
	var unmarshaledRhymeWord RhymeWord
	err = json.Unmarshal(jsonData, &unmarshaledRhymeWord)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON to RhymeWord: %v", err)
	}

	if unmarshaledRhymeWord.Value != rhymeWord.Value || unmarshaledRhymeWord.Rhyme != rhymeWord.Rhyme {
		t.Errorf("Expected RhymeWord value %s and rhyme %s, got value %s and rhyme %s",
			rhymeWord.Value, rhymeWord.Rhyme, unmarshaledRhymeWord.Value, unmarshaledRhymeWord.Rhyme)
	}
}
