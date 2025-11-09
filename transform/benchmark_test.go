package transform

import (
	"go-reloaded/tokenizer"
	"strings"
	"testing"
)

// Large text sample from docs/TEST_CASES.md for performance testing
const largeTextSample = `When developing a text processing application ,you must understand that ff (hex) different algorithms exist for parsing . A experienced (cap) programmer knows that 1010 (bin) main approaches (up, 2) can be used: streaming parsers and batch processors . The streaming approach processes data ' character by character ' while batch processing loads everything into memory first ... However ,both methods have trade-offs ! Streaming uses less memory but batch processing is often faster ? The choice depends on your specific requirements : memory constraints ,performance needs ,and data size . For instance ,if you need to process a HUGE file containing 1a (hex) million records ,streaming might be BETTER (low, 2) for memory usage . Remember that a efficient algorithm should handle edge cases gracefully ... What happens when you encounter malformed input !? Your parser should be ROBUST ENOUGH (low, 2) to continue processing . As donald knuth (cap, 2) said : ' premature optimization is the root of all evil ' in programming . Focus on correctness first ,then optimize . Whether you choose 101 (bin) different data structures or stick to basic arrays ,make sure your code is readable and maintainable . Additionally ,consider that 1f (hex) different optimization techniques (up, 2) exist for performance tuning ... Some developers prefer a object-oriented (cap, 3) approach while others favor functional programming !? The key is understanding your data flow : input validation ,transformation logic ,and output formatting . Modern compilers can optimize 1100 (bin) percent of simple operations automatically . However ,complex algorithms still require careful design ... Remember that a elegant (up) solution often beats a brute-force approach ! When working with large datasets containing abc (hex) thousand entries ,consider using parallel processing . As Linus Torvalds once said : ' talk is cheap ,show me the code ' . Whether you implement 11 (bin) different parsing strategies or just one robust solution ,always test thoroughly with edge cases .`

func BenchmarkProcessTokens(b *testing.B) {
	tokens := tokenizer.Tokenize(largeTextSample)
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ProcessTokens(tokens)
	}
}

func BenchmarkTokenizeAndProcess(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tokens := tokenizer.Tokenize(largeTextSample)
		ProcessTokens(tokens)
	}
}

func BenchmarkLargeInput(b *testing.B) {
	// Create 100x larger input (~200KB)
	largeInput := strings.Repeat(largeTextSample+" ", 100)
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tokens := tokenizer.Tokenize(largeInput)
		ProcessTokens(tokens)
	}
}

func BenchmarkMemoryUsage(b *testing.B) {
	// Test memory allocation patterns
	b.ReportAllocs()
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tokens := tokenizer.Tokenize(largeTextSample)
		ProcessTokens(tokens)
	}
}