package transform

import (
	"testing"

	"go-reloaded/tokenizer"
)

func TestSampleFile(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"adjacent hex no space",
			"1f3(hex)",
			"499"},
		{"large hex value",
			"F6bC (hex)",
			"63164"},
		{"binary 101",
			"101(bin)",
			"5"},
		{"binary 5",
			"5(bin)",
			"5"},
		{"complex punctuation groups",
			"here ! !!",
			"here!!!"},
		{"ellipsis with spaces",
			"be 1f3(hex) ..' weird errors ' here",
			"be 499..'weird errors' here"},
		{"quotes with case transformation",
			"' a upper case ' (up, 3)",
			"'AN UPPER CASE'"},
		{"article with ugly",
			"in A Ugly (low,2) basket",
			"in an ugly basket"},
		{"orange with cap and up",
			"' a orange (cap, 2)'(up, 2)",
			"'AN ORANGE'"},
		{"multiple spaces and punctuation",
			"away !  !",
			"away!!"},
		{"command pattern with can up",
			"' command pattern ' can (up, 3) be",
			"'COMMAND PATTERN' CAN be"},
		{"harold wilson quote",
			"harold wilson (cap, 2) : ' I am a optimist ,but a optimist who carries a raincoat . '",
			"Harold Wilson: 'I am an optimist, but an optimist who carries a raincoat.'"},
		{"complex with parentheses",
			"This (is not) what you 2 would (up,3) expect",
			"This is not what YOU 2 WOULD expect"},
		{"wizards in land",
			"101(bin)wizards in the land (cap, 5)of lizards",
			"5 Wizards In The Land of lizards"},
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