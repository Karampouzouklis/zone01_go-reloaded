package transform

import (
	"fmt"
	"go-reloaded/tokenizer"
	"strconv"
)

// ProcessTokens applies all transformations in pipeline order
func ProcessTokens(tokens []tokenizer.Token) []tokenizer.Token {
	tokens = processHexConversion(tokens)
	// Future transformations will be added here:
	// tokens = processBinaryConversion(tokens)
	// tokens = processCaseTransformations(tokens)
	// tokens = processPunctuation(tokens)
	// tokens = processQuotes(tokens)
	// tokens = processArticles(tokens)
	return tokens
}

// processHexConversion converts hex numbers to decimal
func processHexConversion(tokens []tokenizer.Token) []tokenizer.Token {
	result := make([]tokenizer.Token, 0, len(tokens))
	
	i := 0
	for i < len(tokens) {
		token := tokens[i]
		
		if token.Type == tokenizer.Marker && token.Marker == "hex" {
			// Find the previous word token (skip whitespace)
			wordIndex := -1
			for j := i - 1; j >= 0; j-- {
				if tokens[j].Type == tokenizer.Word {
					wordIndex = j
					break
				} else if tokens[j].Type != tokenizer.Whitespace {
					break
				}
			}
			
			if wordIndex >= 0 {
				// Convert hex to decimal
				if decimal, err := strconv.ParseInt(tokens[wordIndex].Value, 16, 64); err == nil {
					// Remove whitespace before marker from result
					for len(result) > 0 && result[len(result)-1].Type == tokenizer.Whitespace {
						result = result[:len(result)-1]
					}
					
					// Find and replace the word in result
					for k := len(result) - 1; k >= 0; k-- {
						if result[k].Type == tokenizer.Word && result[k].Value == tokens[wordIndex].Value {
							result[k] = tokenizer.Token{
								Type:  tokenizer.Word,
								Value: fmt.Sprintf("%d", decimal),
							}
							break
						}
					}
					
					// Skip the hex marker
					i++
					continue
				}
			}
		}
		
		result = append(result, token)
		i++
	}
	
	return result
}