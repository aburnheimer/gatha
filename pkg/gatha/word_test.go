package gatha

import (
	"encoding/json"
	"testing"
)

// TestWordJsonMarshaling tests the JSON marshaling and unmarshaling of Word.
func TestWordJsonMarshaling(t *testing.T) {
	word := NewNormalWord("Cellar")

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
	var unmarshaledWord NormalWord
	err = json.Unmarshal(jsonData, &unmarshaledWord)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON to Word: %v", err)
	}

	if unmarshaledWord.GetValue() != word.GetValue() {
		t.Errorf("Expected Word value %s, got %s", word.GetValue(), unmarshaledWord.GetValue())
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

	if unmarshaledRhymeWord.GetValue() != rhymeWord.GetValue() || unmarshaledRhymeWord.GetRhyme() != rhymeWord.GetRhyme() {
		t.Errorf("Expected RhymeWord value %s and rhyme %s, got value %s and rhyme %s",
			rhymeWord.GetValue(), rhymeWord.GetRhyme(), unmarshaledRhymeWord.Value, unmarshaledRhymeWord.Rhyme)
	}
}

// TestNormalWordUnmarshalJSONInvalidFormat tests that an error is raised when the JSON format is invalid.
func TestNormalWordUnmarshalJSONInvalidFormat(t *testing.T) {
	invalidJson := `{"value":123}`

	var unmarshaledNormalWord NormalWord
	err := json.Unmarshal([]byte(invalidJson), &unmarshaledNormalWord)
	if err == nil {
		t.Errorf("Expected an error due to invalid JSON format, but got nil")
	}
}

// TestRhymeWordUnmarshalJSONInvalidFormat tests that an error is raised when the JSON format is invalid.
func TestRhymeWordUnmarshalJSONInvalidFormat(t *testing.T) {
	invalidJson := `{"value":"dog","rhyme":123}`

	var unmarshaledRhymeWord RhymeWord
	err := json.Unmarshal([]byte(invalidJson), &unmarshaledRhymeWord)
	if err == nil {
		t.Errorf("Expected an error due to invalid JSON format, but got nil")
	}

	invalidJson = `{"value":123,"rhyme":"A1"}`
	err = json.Unmarshal([]byte(invalidJson), &unmarshaledRhymeWord)
	if err == nil {
		t.Errorf("Expected an error due to invalid JSON format, but got nil")
	}
}

// TestUnmarshalJSONInvalidForNormalWordStruct tests that an error is raised when the JSON is in a bad format.
func TestUnmarshalJSONInvalidForNormalWordStruct(t *testing.T) {
	invalidJson := `[ 123, 456 ]`

	var unmarshaledNormalWord NormalWord
	err := json.Unmarshal([]byte(invalidJson), &unmarshaledNormalWord)
	if err == nil {
		t.Errorf("No error was raised, but it should have been")
	}
}

// TestUnmarshalJSONInvalidForRhymeWordStruct tests that an error is raised when the JSON is in a bad format.
func TestUnmarshalJSONInvalidForRhymeWordStruct(t *testing.T) {
	invalidJson := `[ 123, 456 ]`

	var unmarshaledRhymeWord RhymeWord
	err := json.Unmarshal([]byte(invalidJson), &unmarshaledRhymeWord)
	if err == nil {
		t.Errorf("No error was raised, but it should have been")
	}
}
