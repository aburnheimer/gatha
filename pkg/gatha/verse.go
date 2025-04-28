package gatha

import (
	"strings"
)

// Verse represents a verse of text split into stanzas.
type Verse struct {
	OriginalText string   `json:"-"`
	Stanzas      []Stanza `json:"stanzas"`
}

// NewVerse creates a new Verse instance, parsing the OriginalText into Stanzas.
func NewVerse(originalText string) Verse {
	stanzas := []Stanza{}
	for _, stanzaText := range strings.Split(originalText, "\n\n") {
		stanzas = append(stanzas, NewStanza(stanzaText))
	}
	return Verse{
		OriginalText: originalText,
		Stanzas:      stanzas,
	}
}

// toString returns the text of the Verse, as composed by its Stanzas.
func (v Verse) toString() string {
	return strings.Join(v.StanzasToStrings(), "\n\n")
}

func (v Verse) StanzasToStrings() []string {
	stanzaStrings := make([]string, len(v.Stanzas))
	for i, stanza := range v.Stanzas {
		stanzaStrings[i] = stanza.toString()
	}
	return stanzaStrings
}
