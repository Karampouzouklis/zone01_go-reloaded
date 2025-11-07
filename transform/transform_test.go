package transform

import (
	"testing"

	"go-reloaded/tokenizer"
)

func TestMulti(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"functional case 1", "If I make you BREAKFAST IN BED (low, 3) just say thank you instead of: how (cap) did you get in my house (up, 2) ?", "If I make you breakfast in bed just say thank you instead of: How did you get in MY HOUSE?"},
		{"functional case 2", "I have to pack 101 (bin) outfits. Packed 1a (hex) just to be sure", "I have to pack 5 outfits. Packed 26 just to be sure"},
		{"functional case 3", "Don not be sad ,because sad backwards is das . And das not good", "Don not be sad, because sad backwards is das. And das not good"},
		// TODO: Enable when quote handling and article correction are implemented
		// {"functional case 4", "harold wilson (cap, 2) : ' I am a optimist ,but a optimist who carries a raincoat . '", "Harold Wilson: 'I am an optimist, but an optimist who carries a raincoat.'"},
		{"punctuation groups", "Wait ... What !? No way !! Are you serious ... ?", "Wait... What!? No way!! Are you serious...?"},
		{"overlapping transformations", "Make these THREE WORDS (low, 2) very (up, 3) important", "Make these THREE WORDS VERY important"},
		{"adjacent marker", "Make it loud(up)!", "Make it LOUD!"},
		{"mixed transformations", "Simply add 42 (hex) and 10 (bin) and you get 68.", "Simply add 66 and 2 and you get 68."},
		{"tricky 1 - quotes with transformations", "He said ' I want 1a (hex) apples (up, 2) ' today .", "He said 'I want 26 APPLES' today."},
		// TODO: Enable when quote handling and article correction are implemented
		// {"tricky 4 - article with quoted transformation", "It was a 'honest (cap)' decision .", "It was an 'Honest' decision."},
		// TODO: Enable when article correction is implemented
		// {"comprehensive text", "When developing a text processing application ,you must understand that ff (hex) different algorithms exist for parsing . A experienced (cap) programmer knows that 1010 (bin) main approaches (up, 2) can be used: streaming parsers and batch processors .", "When developing a text processing application, you must understand that 255 different algorithms exist for parsing. An Experienced programmer knows that 10 MAIN APPROACHES can be used: streaming parsers and batch processors."},
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

func TestNumberConversion(t *testing.T) {
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

func TestCaseTransformations(t *testing.T) {
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
		{"multi-word up", "this is exciting (up, 2)", "this IS EXCITING"},
		{"multi-word low", "BREAKFAST IN BED (low, 3)", "breakfast in bed"},
		{"multi-word cap", "hello world test (cap, 2)", "hello World Test"},
		{"count too large", "word (up, 5)", "WORD"},
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

func TestPunctuation(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"comma spacing", "word ,next", "word, next"},
		{"exclamation", "word !", "word!"},
		{"question", "really ?", "really?"},
		{"colon", "note :", "note:"},
		{"multiple punct", "hello , world !", "hello, world!"},
		{"end period", "done .", "done."},
		{"ellipsis", "word ... next", "word... next"},
		{"question-exclamation", "What !?", "What!?"},
		{"double exclamation", "word !!", "word!!"},
		{"mixed groups", "Wait ... What !?", "Wait... What!?"},
		{"complex groups", "thinking ... really !?", "thinking... really!?"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokens := tokenizer.Tokenize(tt.input)
			result := processPunctuation(tokens)

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
func TestQuotes(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"simple quotes", "I am ' awesome '", "I am 'awesome'"},
		{"multi-word quotes", "As he said: ' I am the best '", "As he said: 'I am the best'"},
		{"quotes with punctuation", "He said ' Hello , world ! '", "He said 'Hello, world!'"},
		{"quotes with transformations", "He said ' word (up) '", "He said 'WORD'"},
		{"multiple quote pairs", "' first ' and ' second '", "'first' and 'second'"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokens := tokenizer.Tokenize(tt.input)
			result := processQuotes(tokens)
			
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