package transform

import (
	"go-reloaded/tokenizer"
	"reflect"
	"testing"
)

func TestProcessArticles(t *testing.T) {
	tests := []struct {
		name     string
		input    []tokenizer.Token
		expected []tokenizer.Token
	}{
		{
			name: "a before vowel a",
			input: []tokenizer.Token{
				{Type: tokenizer.Word, Value: "a"},
				{Type: tokenizer.Whitespace, Value: " "},
				{Type: tokenizer.Word, Value: "amazing"},
			},
			expected: []tokenizer.Token{
				{Type: tokenizer.Word, Value: "an"},
				{Type: tokenizer.Whitespace, Value: " "},
				{Type: tokenizer.Word, Value: "amazing"},
			},
		},
		{
			name: "a before vowel e",
			input: []tokenizer.Token{
				{Type: tokenizer.Word, Value: "a"},
				{Type: tokenizer.Whitespace, Value: " "},
				{Type: tokenizer.Word, Value: "elephant"},
			},
			expected: []tokenizer.Token{
				{Type: tokenizer.Word, Value: "an"},
				{Type: tokenizer.Whitespace, Value: " "},
				{Type: tokenizer.Word, Value: "elephant"},
			},
		},
		{
			name: "a before h",
			input: []tokenizer.Token{
				{Type: tokenizer.Word, Value: "a"},
				{Type: tokenizer.Whitespace, Value: " "},
				{Type: tokenizer.Word, Value: "honest"},
			},
			expected: []tokenizer.Token{
				{Type: tokenizer.Word, Value: "an"},
				{Type: tokenizer.Whitespace, Value: " "},
				{Type: tokenizer.Word, Value: "honest"},
			},
		},
		{
			name: "uppercase A",
			input: []tokenizer.Token{
				{Type: tokenizer.Word, Value: "A"},
				{Type: tokenizer.Whitespace, Value: " "},
				{Type: tokenizer.Word, Value: "amazing"},
			},
			expected: []tokenizer.Token{
				{Type: tokenizer.Word, Value: "An"},
				{Type: tokenizer.Whitespace, Value: " "},
				{Type: tokenizer.Word, Value: "amazing"},
			},
		},
		{
			name: "no change for consonant",
			input: []tokenizer.Token{
				{Type: tokenizer.Word, Value: "a"},
				{Type: tokenizer.Whitespace, Value: " "},
				{Type: tokenizer.Word, Value: "car"},
			},
			expected: []tokenizer.Token{
				{Type: tokenizer.Word, Value: "a"},
				{Type: tokenizer.Whitespace, Value: " "},
				{Type: tokenizer.Word, Value: "car"},
			},
		},
		{
			name: "all vowels",
			input: []tokenizer.Token{
				{Type: tokenizer.Word, Value: "a"},
				{Type: tokenizer.Whitespace, Value: " "},
				{Type: tokenizer.Word, Value: "umbrella"},
			},
			expected: []tokenizer.Token{
				{Type: tokenizer.Word, Value: "an"},
				{Type: tokenizer.Whitespace, Value: " "},
				{Type: tokenizer.Word, Value: "umbrella"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := processArticles(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("processArticles() = %+v, want %+v", result, tt.expected)
			}
		})
	}
}