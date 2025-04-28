package gatha

import (
	"strings"
)

// Stanza represents a stanza of text split into lines.
type Stanza struct {
	OriginalText string `json:"-"`
	Lines        []Line `json:"lines"`
}

// NewStanza creates a new Stanza instance, parsing the OriginalText into Lines.
func NewStanza(originalText string) Stanza {
	lines := []Line{}
	for _, lineText := range strings.Split(originalText, "\n") {
		lines = append(lines, NewLine(lineText))
	}
	return Stanza{
		OriginalText: originalText,
		Lines:        lines,
	}
}

// toString returns the text of the Stanza, as composed by its Lines.
func (s Stanza) toString() string {
	return strings.Join(s.LinesToStrings(), "\n")
}

func (s Stanza) LinesToStrings() []string {
	lineStrings := make([]string, len(s.Lines))
	for i, line := range s.Lines {
		lineStrings[i] = line.toString()
	}
	return lineStrings
}
