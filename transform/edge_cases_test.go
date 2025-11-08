package transform

import (
	"go-reloaded/tokenizer"
	"testing"
)

// TestEdgeCases tests defensive programming and error handling
func TestEdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty string", "", ""},
		{"only whitespace", "   \n\t  ", "   \n\t  "},
		{"unmatched quote", "This has ' only one quote", "This has ' only one quote"},
		{"empty quotes", "Empty ' ' quotes", "Empty '' quotes"},
		{"nested quotes", "This ' has ' nested ' quotes", "This 'has' nested ' quotes"},
		{"marker without word", "(up) nothing before", "(up) nothing before"},
		{"marker at start", "(cap) start", "(cap) start"},
		{"multiple spaces", "word    (up)    next", "WORD    next"},
		{"marker after punctuation", "word! (up)", "WORD!"},
		{"invalid count", "word (up, abc)", "word up, abc"},
		{"zero count", "word (up, 0)", "WORD"},
		{"negative count", "word (up, -1)", "word up, 1"},
		{"very large count", "word (up, 999)", "WORD"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokens := tokenizer.Tokenize(tt.input)
			result := ProcessTokens(tokens)

			// Reconstruct text from tokens
			output := ""
			for _, token := range result {
				output += token.Value
			}

			if output != tt.expected {
				t.Errorf("Expected '%s', got '%s'", tt.expected, output)
			}
		})
	}
}