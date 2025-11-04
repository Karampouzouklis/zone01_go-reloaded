package transform

import (
	"go-reloaded/tokenizer"
	"testing"
)

func TestProcessTokens(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"simple hex", "1E (hex)", "30"},
		{"lowercase hex", "ff (hex)", "255"},
		{"uppercase hex", "ABC (hex)", "2748"},
		{"zero", "0 (hex)", "0"},
		{"invalid hex", "xyz (hex)", "xyz (hex)"},
		{"mixed text", "Simply add 42 (hex) and you get 66.", "Simply add 66 and you get 66."},
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

func TestProcessHexConversion(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"simple hex", "1E (hex)", "30"},
		{"lowercase hex", "ff (hex)", "255"},
		{"uppercase hex", "ABC (hex)", "2748"},
		{"zero", "0 (hex)", "0"},
		{"invalid hex", "xyz (hex)", "xyz (hex)"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokens := tokenizer.Tokenize(tt.input)
			result := processHexConversion(tokens)
			
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