package main

import (
	"fmt"
	"os"
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

	// For now, just copy input to output (we'll add transformations later)
	text := string(content)
	
	// Write to output file
	err = os.WriteFile(outputFile, []byte(text), 0644)
	if err != nil {
		return fmt.Errorf("failed to write output file '%s': %v", outputFile, err)
	}

	return nil
}