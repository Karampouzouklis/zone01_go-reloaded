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
		{"zero hex", "0 (hex)", "0"},
		{"invalid hex", "xyz (hex)", "xyz (hex)"},
		{"simple binary", "10 (bin)", "2"},
		{"longer binary", "1010 (bin)", "10"},
		{"all ones", "11111111 (bin)", "255"},
		{"zero binary", "0 (bin)", "0"},
		{"invalid binary", "102 (bin)", "102 (bin)"},
		{"mixed text", "Simply add 42 (hex) and 10 (bin) and you get 68.", "Simply add 66 and 2 and you get 68."},
		{"uppercase", "word (up)", "WORD"},
		{"lowercase", "WORD (low)", "word"},
		{"capitalize", "word (cap)", "Word"},
		{"mixed case low", "WoRd (low)", "word"},
		{"single char up", "a (up)", "A"},
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

func TestProcessNumberConversion(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"simple hex", "1E (hex)", "30"},
		{"lowercase hex", "ff (hex)", "255"},
		{"uppercase hex", "ABC (hex)", "2748"},
		{"zero hex", "0 (hex)", "0"},
		{"invalid hex", "xyz (hex)", "xyz (hex)"},
		{"simple binary", "10 (bin)", "2"},
		{"longer binary", "1010 (bin)", "10"},
		{"all ones", "11111111 (bin)", "255"},
		{"zero binary", "0 (bin)", "0"},
		{"invalid binary", "102 (bin)", "102 (bin)"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokens := tokenizer.Tokenize(tt.input)
			result := processNumberConversion(tokens)
			
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
func TestProcessCaseTransformations(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"uppercase", "word (up)", "WORD"},
		{"lowercase", "WORD (low)", "word"},
		{"capitalize", "word (cap)", "Word"},
		{"mixed case low", "WoRd (low)", "word"},
		{"single char up", "a (up)", "A"},
		{"empty string", " (cap)", " (cap)"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokens := tokenizer.Tokenize(tt.input)
			result := processCaseTransformations(tokens)
			
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