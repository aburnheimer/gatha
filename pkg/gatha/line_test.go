package gatha

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"
)

// TestLineJsonMarshaling tests the JSON marshaling and unmarshaling of Line.
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

// TestLineUnmarshalJSONNormalAndRhymeWords tests that the set of Words contain both Normal- and RhymeWords.
func TestLineUnmarshalJSONNormalAndRhymeWords(t *testing.T) {
	mixedJson := `{"words":[{"value":"dog","rhyme":""}, {"value":"puppy","rhyme":"A1"}]}`

	var unmarshaledLine Line
	err := json.Unmarshal([]byte(mixedJson), &unmarshaledLine)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON to Line: %v", err)
	}

	if reflect.TypeOf(unmarshaledLine.GetWord(0)).String() != "*gatha.NormalWord" {
		t.Errorf("Expected Word to be of type NormalWord, got this instead %s", reflect.TypeOf(unmarshaledLine.GetWord(0)).String())
	}

	if reflect.TypeOf(unmarshaledLine.GetWord(1)).String() != "*gatha.RhymeWord" {
		t.Errorf("Expected Word to be of type RhymeWord, got this instead %s", reflect.TypeOf(unmarshaledLine.GetWord(1)).String())
	}
}

// TestLineUnmarshalJSONMissingValue tests that an error is raised when a word has a missing or invalid value during unmarshaling.
func TestLineUnmarshalJSONMissingValue(t *testing.T) {
	invalidJson := `{"words":[{"value":"","rhyme":"cat"}]}`

	var unmarshaledLine Line
	err := json.Unmarshal([]byte(invalidJson), &unmarshaledLine)
	if err == nil {
		t.Errorf("Expected an error due to missing or invalid word value, but got nil")
	}
}

// TestLineUnmarshalJSONInvalidFormat tests that an error is raised when the JSON format is invalid.
func TestLineUnmarshalJSONInvalidFormat(t *testing.T) {
	invalidJson := `{"words":[{"value":"dog","rhyme":123}]}`

	var unmarshaledLine Line
	err := json.Unmarshal([]byte(invalidJson), &unmarshaledLine)
	if err == nil {
		t.Errorf("Expected an error due to invalid JSON format, but got nil")
	}
}

// TestLineGetWordValidIndex tests retrieving a word with a valid index.
func TestLineGetWordValidIndex(t *testing.T) {
	line := NewLine("The quick brown fox")
	word := line.GetWord(2)

	if word.GetValue() != "brown" {
		t.Errorf("Expected word 'brown', got '%s'", word.GetValue())
	}
}

// TestLineGetWordInvalidNegativeIndex tests retrieving a word with a negative index.
func TestLineGetWordInvalidNegativeIndex(t *testing.T) {
	line := NewLine("The quick brown fox")
	word := line.GetWord(-1)

	if word.GetValue() != "" {
		t.Errorf("Expected empty word for negative index, got '%s'", word.GetValue())
	}
}

// TestLineGetWordInvalidOutOfBoundsIndex tests retrieving a word with an out-of-bounds index.
func TestLineGetWordInvalidOutOfBoundsIndex(t *testing.T) {
	line := NewLine("The quick brown fox")
	word := line.GetWord(10)

	if word.GetValue() != "" {
		t.Errorf("Expected empty word for out-of-bounds index, got '%s'", word.GetValue())
	}
}

// TestLineGetWordEmptyLine tests retrieving a word from an empty Line.
func TestLineGetWordEmptyLine(t *testing.T) {
	line := NewLine("")
	word := line.GetWord(0)

	if word.GetValue() != "" {
		t.Errorf("Expected empty word for empty Line, got '%s'", word.GetValue())
	}
}
