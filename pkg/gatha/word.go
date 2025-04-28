// Package gatha provides objects and methods for working with verses of poetry
package gatha

// Word represents a single word in a Line of poetry.
type Word struct {
	Value string `json:"value"`
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
