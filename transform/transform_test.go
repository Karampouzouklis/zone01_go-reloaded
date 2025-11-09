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
		{"Functional Case 1",
			"If I make you BREAKFAST IN BED (low, 3) just say thank you instead of: how (cap) did you get in my house (up, 2) ?",
			"If I make you breakfast in bed just say thank you instead of: How did you get in MY HOUSE?"},
		{"Functional Case 2",
			"I have to pack 101 (bin) outfits. Packed 1a (hex) just to be sure",
			"I have to pack 5 outfits. Packed 26 just to be sure"},
		{"Functional Case 3",
			"Don not be sad ,because sad backwards is das . And das not good",
			"Don not be sad, because sad backwards is das. And das not good"},
		{"Functional Case 4",
			"harold wilson (cap, 2) : ' I am a optimist ,but a optimist who carries a raincoat . '",
			"Harold Wilson: 'I am an optimist, but an optimist who carries a raincoat.'"},
		{"Tricky 2 - punctuation groups",
			"Wait ... What !? No way !! Are you serious ... ?",
			"Wait... What!? No way!! Are you serious...?"},
		{"Tricky 3 - overlapping transformations",
			"Make these THREE WORDS (low, 2) very (up, 3) important",
			"Make these THREE WORDS VERY important"},
		{"adjacent marker",
			"Make it loud(up)!",
			"Make it LOUD!"},
		{"mixed transformations",
			"Simply add 42 (hex) and 10 (bin) and you get 68.",
			"Simply add 66 and 2 and you get 68."},
		{"Simple Quotes",
			"I am ' awesome '",
			"I am 'awesome'"},
		{"Tricky 1 - quotes with transformations",
			"He said ' I want 1a (hex) apples (up, 2) ' today .",
			"He said 'I want 26 APPLES' today."},
		{"Tricky 4 - article with quoted transformation",
			"It was a 'honest (cap)' decision .",
			"It was an 'Honest' decision."},
		{"comprehensive text",
			"When developing a text processing application ,you must understand that ff (hex) different algorithms exist for parsing . A experienced (cap) programmer knows that 1010 (bin) main approaches (up, 2) can be used: streaming parsers and batch processors .",
			"When developing a text processing application, you must understand that 255 different algorithms exist for parsing. An Experienced programmer knows that 10 MAIN APPROACHES can be used: streaming parsers and batch processors."},
		{"quotes with punctuation full pipeline",
			"He said ' Hello , world ! '",
			"He said 'Hello, world!'"},
		{"quotes with transformations full pipeline",
			"He said ' word (up) '",
			"He said 'WORD'"},
		{"articles with punctuation full pipeline",
			"a elephant !",
			"an elephant!"},
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
		{"invalid hex", "xyz (hex)", "xyz"},
		{"simple binary", "10 (bin)", "2"},
		{"longer binary", "1010 (bin)", "10"},
		{"all ones", "11111111 (bin)", "255"},
		{"zero binary", "0 (bin)", "0"},
		{"invalid binary", "102 (bin)", "102"},
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

func TestArticles(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"a before vowel a", "a amazing", "an amazing"},
		{"a before vowel e", "a elephant", "an elephant"},
		{"a before vowel i", "a interesting", "an interesting"},
		{"a before vowel o", "a orange", "an orange"},
		{"a before vowel u", "a umbrella", "an umbrella"},
		{"a before h", "a honest", "an honest"},
		{"uppercase A", "A amazing", "An amazing"},
		{"no change for consonant", "a car", "a car"},
		{"multiple articles", "a apple and a orange", "an apple and an orange"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokens := tokenizer.Tokenize(tt.input)
			result := processArticles(tokens)

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
