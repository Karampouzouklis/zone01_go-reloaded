package tokenizer

import (
	"regexp"
	"strconv"
	"strings"
)

// TokenType represents the type of token
type TokenType int

const (
	Word TokenType = iota
	Punctuation
	Marker
	Whitespace
)

// Token represents a single token in the text
type Token struct {
	Type   TokenType
	Value  string // Original text value
	Marker string // For transformation markers (hex, bin, up, etc.)
	Count  int    // For multi-word transformations
}

// Tokenize splits input text into tokens
func Tokenize(text string) []Token {
	var tokens []Token
	
	// Regex patterns
	markerPattern := regexp.MustCompile(`\([a-z]+(?:,\s*\d+)?\)`)
	punctPattern := regexp.MustCompile(`[.,:;!?']+`)
	wordPattern := regexp.MustCompile(`[a-zA-Z0-9]+`)
	spacePattern := regexp.MustCompile(`\s+`)
	
	i := 0
	for i < len(text) {
		// Check for markers first (highest priority)
		if markerMatch := markerPattern.FindStringIndex(text[i:]); markerMatch != nil && markerMatch[0] == 0 {
			markerText := text[i : i+markerMatch[1]]
			token := parseMarker(markerText)
			tokens = append(tokens, token)
			i += markerMatch[1]
			continue
		}
		
		// Check for punctuation
		if punctMatch := punctPattern.FindStringIndex(text[i:]); punctMatch != nil && punctMatch[0] == 0 {
			punctText := text[i : i+punctMatch[1]]
			tokens = append(tokens, Token{Type: Punctuation, Value: punctText})
			i += punctMatch[1]
			continue
		}
		
		// Check for words
		if wordMatch := wordPattern.FindStringIndex(text[i:]); wordMatch != nil && wordMatch[0] == 0 {
			wordText := text[i : i+wordMatch[1]]
			tokens = append(tokens, Token{Type: Word, Value: wordText})
			i += wordMatch[1]
			continue
		}
		
		// Check for whitespace
		if spaceMatch := spacePattern.FindStringIndex(text[i:]); spaceMatch != nil && spaceMatch[0] == 0 {
			spaceText := text[i : i+spaceMatch[1]]
			tokens = append(tokens, Token{Type: Whitespace, Value: spaceText})
			i += spaceMatch[1]
			continue
		}
		
		// Skip unknown characters
		i++
	}
	
	return tokens
}

// parseMarker extracts marker type and count from marker text
func parseMarker(markerText string) Token {
	// Remove parentheses
	inner := strings.Trim(markerText, "()")
	
	// Check for count (e.g., "up, 2")
	if strings.Contains(inner, ",") {
		parts := strings.Split(inner, ",")
		markerType := strings.TrimSpace(parts[0])
		countStr := strings.TrimSpace(parts[1])
		count, _ := strconv.Atoi(countStr)
		return Token{
			Type:   Marker,
			Value:  markerText,
			Marker: markerType,
			Count:  count,
		}
	}
	
	// Simple marker without count
	return Token{
		Type:   Marker,
		Value:  markerText,
		Marker: inner,
		Count:  0,
	}
}