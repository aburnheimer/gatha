// Package gatha provides objects and methods for working with verses of poetry
package gatha

import "strings"

// Line represents a line of text split into words.
type Line struct {
	OriginalText string `json:"-"`
	Words        []Word `json:"words"`
}

// NewLine creates a new Line instance, parsing the OriginalText into Words.
func NewLine(originalText string) Line {
	words := []Word{}
	for _, word := range strings.Fields(originalText) {
		words = append(words, Word{Value: word})
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
		wordStrings[i] = word.Value
	}
	return wordStrings
}
