package main

import (
	"fmt"
	"os"
	"go-reloaded/tokenizer"
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
	
	// For now, reconstruct text from tokens (demonstrating tokenizer works)
	result := ""
	for _, token := range tokens {
		result += token.Value
	}
	
	// Write to output file
	err = os.WriteFile(outputFile, []byte(result), 0644)
	if err != nil {
		return fmt.Errorf("failed to write output file '%s': %v", outputFile, err)
	}

	return nil
}