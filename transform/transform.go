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
	tokens = processArticles(tokens)
	tokens = processQuotes(tokens)
	tokens = processCaseTransformations(tokens)
	tokens = processPunctuation(tokens)
	return tokens
}

// processNumberConversion converts hex and binary numbers to decimal
func processNumberConversion(tokens []tokenizer.Token) []tokenizer.Token {
	result := make([]tokenizer.Token, 0, len(tokens))

	i := 0
	for i < len(tokens) {
		token := tokens[i]

		if token.Type == tokenizer.Command && (token.Command == "hex" || token.Command == "bin") {
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
				if token.Command == "bin" {
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

		if token.Type == tokenizer.Command && (token.Command == "up" || token.Command == "low" || token.Command == "cap") {
			// Find the previous word token (skip whitespace and quotes)
			wordIndex := -1
			for j := i - 1; j >= 0; j-- {
				if tokens[j].Type == tokenizer.Word {
					wordIndex = j
					break
				} else if tokens[j].Type != tokenizer.Whitespace && tokens[j].Type != tokenizer.Quote {
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

				// Transform the specified number of words backward (skip quotes and whitespace)
				wordsTransformed := 0
				for k := len(result) - 1; k >= 0 && wordsTransformed < count; k-- {
					if result[k].Type == tokenizer.Word {
						word := result[k].Value
						var transformed string

						switch token.Command {
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
					// Skip quotes and whitespace when counting backward
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

// processPunctuation handles spacing for punctuation marks .,!?:; and groups
func processPunctuation(tokens []tokenizer.Token) []tokenizer.Token {
	result := make([]tokenizer.Token, 0, len(tokens))

	i := 0
	for i < len(tokens) {
		token := tokens[i]

		if token.Type == tokenizer.Punctuation {
			// Remove whitespace before punctuation
			for len(result) > 0 && result[len(result)-1].Type == tokenizer.Whitespace {
				result = result[:len(result)-1]
			}

			// Collect consecutive punctuation marks (skip whitespace between them)
			punctGroup := token.Value
			j := i + 1
			for j < len(tokens) {
				if tokens[j].Type == tokenizer.Punctuation {
					punctGroup += tokens[j].Value
					j++
				} else if tokens[j].Type == tokenizer.Whitespace {
					// Skip whitespace and check if next is punctuation
					if j+1 < len(tokens) && tokens[j+1].Type == tokenizer.Punctuation {
						j++
					} else {
						break
					}
				} else {
					break
				}
			}

			// Add the punctuation group as single token
			result = append(result, tokenizer.Token{Type: tokenizer.Punctuation, Value: punctGroup})

			// Add single space after punctuation group if there's a following word
			if j < len(tokens) && tokens[j].Type != tokenizer.Whitespace {
				result = append(result, tokenizer.Token{Type: tokenizer.Whitespace, Value: " "})
			}

			// Skip processed tokens
			i = j
		} else {
			result = append(result, token)
			i++
		}
	}

	return result
}

// processQuotes handles single quote positioning around words and phrases
func processQuotes(tokens []tokenizer.Token) []tokenizer.Token {
	result := make([]tokenizer.Token, 0, len(tokens))

	i := 0
	for i < len(tokens) {
		token := tokens[i]

		if token.Type == tokenizer.Quote {
			// Find matching closing quote
			closeIndex := -1
			for j := i + 1; j < len(tokens); j++ {
				if tokens[j].Type == tokenizer.Quote {
					closeIndex = j
					break
				}
			}

			if closeIndex != -1 {
				// Collect content between quotes
				quoteContent := make([]tokenizer.Token, 0)
				for k := i + 1; k < closeIndex; k++ {
					quoteContent = append(quoteContent, tokens[k])
				}

				// Remove leading/trailing whitespace from content
				start := 0
				end := len(quoteContent)

				for start < end && quoteContent[start].Type == tokenizer.Whitespace {
					start++
				}

				for end > start && quoteContent[end-1].Type == tokenizer.Whitespace {
					end--
				}

				// Add opening quote
				result = append(result, token)

				// Add content without leading/trailing whitespace
				for k := start; k < end; k++ {
					result = append(result, quoteContent[k])
				}

				// Add closing quote
				result = append(result, tokenizer.Token{Type: tokenizer.Quote, Value: "'"})

				// Skip to after closing quote
				i = closeIndex + 1
				continue
			}
		}

		result = append(result, token)
		i++
	}

	return result
}

// processArticles converts 'a' to 'an' before vowels (a,e,i,o,u) and 'h'
func processArticles(tokens []tokenizer.Token) []tokenizer.Token {
	result := make([]tokenizer.Token, 0, len(tokens))

	for i, token := range tokens {
		if token.Type == tokenizer.Word && (token.Value == "a" || token.Value == "A") {
			// Look for next word (skip whitespace)
			nextWordIndex := -1
			for j := i + 1; j < len(tokens); j++ {
				if tokens[j].Type == tokenizer.Word {
					nextWordIndex = j
					break
				} else if tokens[j].Type != tokenizer.Whitespace {
					break
				}
			}

			if nextWordIndex != -1 {
				nextWord := tokens[nextWordIndex].Value
				if len(nextWord) > 0 {
					firstChar := strings.ToLower(string(nextWord[0]))
					// Check if starts with vowel or h
					if firstChar == "a" || firstChar == "e" || firstChar == "i" || firstChar == "o" || firstChar == "u" || firstChar == "h" {
						// Convert a->an, A->An
						if token.Value == "a" {
							token.Value = "an"
						} else {
							token.Value = "An"
						}
					}
				}
			}
		}

		result = append(result, token)
	}

	return result
}
