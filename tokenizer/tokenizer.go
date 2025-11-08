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
	Command
	Quote
	Whitespace
)

// Token represents a single token in the text
type Token struct {
	Type   TokenType
	Value  string // Original text value
	Command string // For transformation commands (hex, bin, up, etc.)
	Count  int    // For multi-word transformations
}

// Tokenize splits input text into tokens
func Tokenize(text string) []Token {
	var tokens []Token
	
	// Regex patterns
	commandPattern := regexp.MustCompile(`\([a-z]+(?:,\s*\d+)?\)`)
	quotePattern := regexp.MustCompile(`'`)
	punctPattern := regexp.MustCompile(`[.,:;!?]+`)
	wordPattern := regexp.MustCompile(`[a-zA-Z0-9]+`)
	spacePattern := regexp.MustCompile(`\s+`)
	
	i := 0
	for i < len(text) {
		// Check for commands first (highest priority)
		if commandMatch := commandPattern.FindStringIndex(text[i:]); commandMatch != nil && commandMatch[0] == 0 {
			commandText := text[i : i+commandMatch[1]]
			token := parseCommand(commandText)
			tokens = append(tokens, token)
			i += commandMatch[1]
			continue
		}
		
		// Check for quotes
		if quoteMatch := quotePattern.FindStringIndex(text[i:]); quoteMatch != nil && quoteMatch[0] == 0 {
			quoteText := text[i : i+quoteMatch[1]]
			tokens = append(tokens, Token{Type: Quote, Value: quoteText})
			i += quoteMatch[1]
			continue
		}
		
		// Check for other punctuation
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

// parseCommand extracts command type and count from command text
func parseCommand(commandText string) Token {
	// Remove parentheses
	inner := strings.Trim(commandText, "()")
	
	// Check for count (e.g., "up, 2")
	if strings.Contains(inner, ",") {
		parts := strings.Split(inner, ",")
		commandType := strings.TrimSpace(parts[0])
		countStr := strings.TrimSpace(parts[1])
		count, _ := strconv.Atoi(countStr)
		return Token{
			Type:    Command,
			Value:   commandText,
			Command: commandType,
			Count:   count,
		}
	}
	
	// Simple command without count
	return Token{
		Type:    Command,
		Value:   commandText,
		Command: inner,
		Count:   0,
	}
}