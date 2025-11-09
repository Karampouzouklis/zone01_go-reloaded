package transform

import (
	"go-reloaded/tokenizer"
	"strings"
	"testing"
)

// BenchmarkStringBuilder tests string builder vs string concatenation
func BenchmarkStringBuilder(b *testing.B) {
	tokens := tokenizer.Tokenize(largeTextSample)
	processedTokens := ProcessTokens(tokens)
	
	b.Run("StringConcatenation", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			output := ""
			for _, token := range processedTokens {
				output += token.Value
			}
			_ = output
		}
	})
	
	b.Run("StringBuilder", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			var builder strings.Builder
			builder.Grow(len(largeTextSample)) // Pre-allocate
			for _, token := range processedTokens {
				builder.WriteString(token.Value)
			}
			_ = builder.String()
		}
	})
}