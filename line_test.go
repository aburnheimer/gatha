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

	expectedJson, err := os.ReadFile("test/expected_line.json")
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

// Replace a normal Word with one that has rhyme data annotation (RhymeWord)
func TestLineReplacingWordWithRhymeWord(t *testing.T) {
	originalText := "The quick brown fox jumps over the lazy dog"
	line := NewLine(originalText)

	word := line.GetWord(8)
	if reflect.TypeOf(word).String() != "*gatha.NormalWord" {
		t.Errorf("Expected Word to be of type NormalWord, got this instead %s", reflect.TypeOf(word).String())
	}

	line.AnnotateRhymeToWord(8, "A1")

	word = line.GetWord(8)
	if reflect.TypeOf(word).String() != "*gatha.RhymeWord" {
		t.Errorf("Expected Word to be of type RhymeWord, got this instead %s", reflect.TypeOf(word).String())
	}
}

// TestAnnotateRhymeToWordValidIndex tests annotating a word with a valid index.
func TestAnnotateRhymeToWordValidIndex(t *testing.T) {
	line := NewLine("The quick brown fox jumps over the lazy dog")
	line.AnnotateRhymeToWord(3, "A1")

	word := line.GetWord(3)
	if rhymeWord, ok := word.(*RhymeWord); ok {
		if rhymeWord.Rhyme != "A1" {
			t.Errorf("Expected rhyme 'A1', got '%s'", rhymeWord.Rhyme)
		}
	} else {
		t.Errorf("Expected word to be of type RhymeWord, got %T", word)
	}
}

// TestAnnotateRhymeToWordInvalidNegativeIndex tests annotating a word with a negative index.
func TestAnnotateRhymeToWordInvalidNegativeIndex(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for negative index, but no panic occurred")
		}
	}()

	line := NewLine("The quick brown fox jumps over the lazy dog")
	line.AnnotateRhymeToWord(-1, "A1")
}

// TestAnnotateRhymeToWordInvalidOutOfBoundsIndex tests annotating a word with an out-of-bounds index.
func TestAnnotateRhymeToWordInvalidOutOfBoundsIndex(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for out-of-bounds index, but no panic occurred")
		}
	}()

	line := NewLine("The quick brown fox jumps over the lazy dog")
	line.AnnotateRhymeToWord(10, "A1")
}

// TestAnnotateRhymeToWordExistingRhymeWord tests annotating a word that is already a RhymeWord.
func TestAnnotateRhymeToWordExistingRhymeWord(t *testing.T) {
	line := NewLine("The quick brown fox jumps over the lazy dog")
	line.AnnotateRhymeToWord(3, "A1")
	line.AnnotateRhymeToWord(3, "B2")

	word := line.GetWord(3)
	if rhymeWord, ok := word.(*RhymeWord); ok {
		if rhymeWord.Rhyme != "B2" {
			t.Errorf("Expected updated rhyme 'B2', got '%s'", rhymeWord.Rhyme)
		}
	} else {
		t.Errorf("Expected word to be of type RhymeWord, got %T", word)
	}
}
