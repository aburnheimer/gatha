// Package gatha provides objects and methods for working with verses of poetry
package gatha

// Word represents a single word in a Line of poetry.
type Word struct {
	Value string `json:"value"`
}

// NewWord creates a new Word instance.
func NewWord(value string) Word {
	return Word{Value: value}
}

type Marshalable interface {
	ToJson()
	FromJson()
}

// RhymeWord is a Word that lands as a rhyme in a meter of poetry.
type RhymeWord struct {
	Word
	Rhyme string `json:"rhyme"`
}

// NewRhymeWord creates a new RhymeWord instance.
func NewRhymeWord(value string) RhymeWord {
	return RhymeWord{Word: Word{Value: value}}
}

// SetRhyme sets the Rhyme for the RhymeWord.
func (rw *RhymeWord) SetRhyme(rhyme string) {
	rw.Rhyme = rhyme
}
