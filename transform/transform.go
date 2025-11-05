package transform

import (
	"fmt"
	"go-reloaded/tokenizer"
	"strconv"
	"strings"
)

// ProcessTokens applies all transformations in pipeline order
func ProcessTokens(tokens []tokenizer.Token) []tokenizer.Token {
	tokens = processNumberConversion(tokens)
	tokens = processCaseTransformations(tokens)
	// Future transformations will be added here:
	// tokens = processPunctuation(tokens)
	// tokens = processQuotes(tokens)
	// tokens = processArticles(tokens)
	return tokens
}

// processNumberConversion converts hex and binary numbers to decimal
func processNumberConversion(tokens []tokenizer.Token) []tokenizer.Token {
	result := make([]tokenizer.Token, 0, len(tokens))
	
	i := 0
	for i < len(tokens) {
		token := tokens[i]
		
		if token.Type == tokenizer.Marker && (token.Marker == "hex" || token.Marker == "bin") {
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
				// Determine base: hex=16, bin=2
				base := 16
				if token.Marker == "bin" {
					base = 2
				}
				
				// Convert to decimal
				if decimal, err := strconv.ParseInt(tokens[wordIndex].Value, base, 64); err == nil {
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
					
					// Skip the marker
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

// processCaseTransformations handles (up), (low), and (cap) markers
func processCaseTransformations(tokens []tokenizer.Token) []tokenizer.Token {
	result := make([]tokenizer.Token, 0, len(tokens))
	
	i := 0
	for i < len(tokens) {
		token := tokens[i]
		
		if token.Type == tokenizer.Marker && (token.Marker == "up" || token.Marker == "low" || token.Marker == "cap") {
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
				// Determine how many words to transform
				count := token.Count
				if count == 0 {
					count = 1
				}
				
				// Remove whitespace before marker from result
				for len(result) > 0 && result[len(result)-1].Type == tokenizer.Whitespace {
					result = result[:len(result)-1]
				}
				
				// Transform the specified number of words backward
				wordsTransformed := 0
				for k := len(result) - 1; k >= 0 && wordsTransformed < count; k-- {
					if result[k].Type == tokenizer.Word {
						word := result[k].Value
						var transformed string
						
						switch token.Marker {
						case "up":
							transformed = strings.ToUpper(word)
						case "low":
							transformed = strings.ToLower(word)
						case "cap":
							if len(word) > 0 {
								transformed = strings.ToUpper(word[:1]) + strings.ToLower(word[1:])
							} else {
								transformed = word
							}
						}
						
						result[k] = tokenizer.Token{
							Type:  tokenizer.Word,
							Value: transformed,
						}
						wordsTransformed++
					}
				}
				
				// Skip the marker
				i++
				continue
			}
		}
		
		result = append(result, token)
		i++
	}
	
	return result
}