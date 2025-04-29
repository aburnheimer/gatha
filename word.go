// Package gatha provides objects and methods for working with verses of poetry
package gatha

import (
	"encoding/json"
	"fmt"
)

// Word represents an abstract word in a Line of poetry, either rhyming or otherwise.
type Word interface {
	GetValue() string
	UnmarshalJSON(data []byte) error
}

// NormalWord represents a single word in a Line of poetry.
type NormalWord struct {
	Value string `json:"value"`
}

// NewWord creates a new Word instance.
func NewNormalWord(value string) NormalWord {
	return NormalWord{Value: value}
}

// GetValue gets the Value for the RhymeWord.
func (nw *NormalWord) GetValue() string {
	return nw.Value
}

func (nw *NormalWord) UnmarshalJSON(data []byte) error {
	var nwObj map[string]interface{}
	if err := json.Unmarshal(data, &nwObj); err != nil {
		return err
	}
	if value, ok := nwObj["value"].(string); ok {
		nw.Value = value
	} else {
		return fmt.Errorf("value is not a string")
	}
	return nil
}

// RhymeWord is a Word that lands as a rhyme in a meter of poetry.
type RhymeWord struct {
	NormalWord
	Rhyme string `json:"rhyme"`
}

// NewRhymeWord creates a new RhymeWord instance.
func NewRhymeWord(value string) RhymeWord {
	return RhymeWord{NormalWord: NormalWord{Value: value}}
}

// GetValue gets the Value for the RhymeWord.
func (rw *RhymeWord) GetValue() string {
	return rw.Value
}

// GetRhyme gets the Rhyme for the RhymeWord.
func (rw *RhymeWord) GetRhyme() string {
	return rw.Rhyme
}

// SetRhyme sets the Rhyme for the RhymeWord.
func (rw *RhymeWord) SetRhyme(rhyme string) {
	rw.Rhyme = rhyme
}

func (rw *RhymeWord) UnmarshalJSON(data []byte) error {
	var rwObj map[string]interface{}
	if err := json.Unmarshal(data, &rwObj); err != nil {
		return err
	}
	if value, ok := rwObj["value"].(string); ok {
		rw.Value = value
	} else {
		return fmt.Errorf("value is not a string")
	}

	if rhyme, ok := rwObj["rhyme"].(string); ok {
		rw.Rhyme = rhyme
	} else {
		return fmt.Errorf("rhyme is not a string")
	}
	return nil
}
