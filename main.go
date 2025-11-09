package main

import (
	"fmt"
	"os"
	"strings"
	"go-reloaded/tokenizer"
	"go-reloaded/transform"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("usage: go-reloaded <input.txt> <output.txt>")
	}

	inputFile := args[0]
	outputFile := args[1]

	// Read input file
	content, err := os.ReadFile(inputFile)
	if err != nil {
		return fmt.Errorf("failed to read input file '%s': %v", inputFile, err)
	}

	// Tokenize the input text
	text := string(content)
	tokens := tokenizer.Tokenize(text)
	
	// Apply transformations
	tokens = transform.ProcessTokens(tokens)
	
	// Reconstruct text from transformed tokens using string builder
	var builder strings.Builder
	builder.Grow(len(content)) // Pre-allocate based on input size
	for _, token := range tokens {
		builder.WriteString(token.Value)
	}
	result := builder.String()
	
	// Write to output file
	err = os.WriteFile(outputFile, []byte(result), 0644)
	if err != nil {
		return fmt.Errorf("failed to write output file '%s': %v", outputFile, err)
	}

	return nil
}