package tokenizer

import (
	"reflect"
	"testing"
)

func TestTokenize(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []Token
	}{
		{
			name:  "simple words",
			input: "hello world",
			expected: []Token{
				{Type: Word, Value: "hello"},
				{Type: Whitespace, Value: " "},
				{Type: Word, Value: "world"},
			},
		},
		{
			name:  "with punctuation",
			input: "hello, world!",
			expected: []Token{
				{Type: Word, Value: "hello"},
				{Type: Punctuation, Value: ","},
				{Type: Whitespace, Value: " "},
				{Type: Word, Value: "world"},
				{Type: Punctuation, Value: "!"},
			},
		},
		{
			name:  "transformation commands",
			input: "word (up)",
			expected: []Token{
				{Type: Word, Value: "word"},
				{Type: Whitespace, Value: " "},
				{Type: Command, Value: "(up)", Command: "up"},
			},
		},
		{
			name:  "counted commands",
			input: "words (up, 2)",
			expected: []Token{
				{Type: Word, Value: "words"},
				{Type: Whitespace, Value: " "},
				{Type: Command, Value: "(up, 2)", Command: "up", Count: 2},
			},
		},
		{
			name:  "hex command",
			input: "1E (hex)",
			expected: []Token{
				{Type: Word, Value: "1E"},
				{Type: Whitespace, Value: " "},
				{Type: Command, Value: "(hex)", Command: "hex"},
			},
		},
		{
			name:  "multiple punctuation",
			input: "hello... world!?",
			expected: []Token{
				{Type: Word, Value: "hello"},
				{Type: Punctuation, Value: "..."},
				{Type: Whitespace, Value: " "},
				{Type: Word, Value: "world"},
				{Type: Punctuation, Value: "!?"},
			},
		},
		{
			name:  "quotes",
			input: "I am ' awesome '",
			expected: []Token{
				{Type: Word, Value: "I"},
				{Type: Whitespace, Value: " "},
				{Type: Word, Value: "am"},
				{Type: Whitespace, Value: " "},
				{Type: Quote, Value: "'"},
				{Type: Whitespace, Value: " "},
				{Type: Word, Value: "awesome"},
				{Type: Whitespace, Value: " "},
				{Type: Quote, Value: "'"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Tokenize(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Tokenize() = %+v, want %+v", result, tt.expected)
			}
		})
	}
}