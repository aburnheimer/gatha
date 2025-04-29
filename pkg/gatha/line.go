// Package gatha provides objects and methods for working with verses of poetry
package gatha

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Line represents a line of text split into words.
type Line struct {
	OriginalText string `json:"-"`
	Words        []Word `json:"words"`
}

type MarshaledWords struct {
	Words []struct {
		Value string `json:"value"`
		Rhyme string `json:"rhyme,omitempty"`
	} `json:"words"`
}

// NewLine creates a new Line instance, parsing the OriginalText into Words.
func NewLine(originalText string) Line {
	words := []Word{}
	for _, word := range strings.Fields(originalText) {
		words = append(words, &NormalWord{Value: word})
	}
	return Line{
		OriginalText: originalText,
		Words:        words,
	}
}

// toString returns the text of the Line, as composed by the Words in the Line.
func (l Line) toString() string {
	return strings.Join(l.WordsToStrings(), " ")
}

func (l Line) WordsToStrings() []string {
	wordStrings := make([]string, len(l.Words))
	for i, word := range l.Words {
		wordStrings[i] = word.GetValue()
	}
	return wordStrings
}

// UnmarshalJSON customizes the JSON unmarshaling for Line.
func (l *Line) UnmarshalJSON(data []byte) error {
	var words []Word
	var rawWordsObj MarshaledWords
	if err := json.Unmarshal([]byte(data), &rawWordsObj); err != nil {
		return err
	}

	rawWords := rawWordsObj.Words

	wordType := "NormalWord"
	for _, rawWord := range rawWords {
		rhyme := rawWord.Rhyme
		if rhyme != "" {
			wordType = "RhymeWord"
		}

		value := rawWord.Value
		if value == "" {
			return fmt.Errorf("missing or invalid word value")
		}

		switch wordType {
		case "RhymeWord":
			words = append(words, &RhymeWord{NormalWord: NormalWord{Value: value}, Rhyme: rhyme})
		default:
			words = append(words, &NormalWord{Value: value})
		}
	}

	l.Words = words
	return nil
}

// GetWord retrieves a word from the Words slice by its index.
func (l Line) GetWord(index int) Word {
	if index < 0 || index >= len(l.Words) {
		return &NormalWord{} // Return an empty NormalWord pointer if the index is out of bounds
	}
	return l.Words[index]
}
