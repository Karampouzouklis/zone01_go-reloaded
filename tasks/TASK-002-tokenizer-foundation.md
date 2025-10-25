# TASK-002: Tokenizer Foundation

## Objective
Implement text tokenization to split input into structured tokens (words, punctuation, markers).

## Acceptance Criteria
- [ ] Tokenizer package created with clear API
- [ ] Identifies words, punctuation, and transformation markers
- [ ] Returns slice of Token structs
- [ ] Handles whitespace and special characters correctly
- [ ] Comprehensive test coverage for all token types

## Key Learning Concepts
- String manipulation with `strings` package
- Regular expressions with `regexp` package
- Slice operations and append
- Struct design for token representation
- Package organization in Go

## Technical Requirements

### Token Structure
```go
type Token struct {
    Type    TokenType
    Value   string
    Marker  string  // For transformation markers like "(hex)"
    Count   int     // For multi-word transformations
}

type TokenType int
const (
    Word TokenType = iota
    Punctuation
    Marker
    Whitespace
)
```

### Implementation Details
- Use regex or string parsing for tokenization
- Recognize transformation markers: `(hex)`, `(bin)`, `(up)`, `(low)`, `(cap)`
- Handle markers with counts: `(up, 2)`, `(low, 3)`
- Preserve original spacing information

## Test Cases
1. **Simple Text**: "hello world" → [Word("hello"), Whitespace(" "), Word("world")]
2. **With Punctuation**: "hello, world!" → [Word("hello"), Punctuation(","), ...]
3. **Transformation Markers**: "word (up)" → [Word("word"), Whitespace(" "), Marker("(up)")]
4. **Counted Markers**: "words (up, 2)" → [..., Marker("(up, 2)") with Count=2]

## Definition of Done
- [ ] All tokenizer tests pass
- [ ] Handles all transformation marker types
- [ ] Preserves necessary whitespace information
- [ ] Code documented with examples
- [ ] Integration with main.go verified

## References
- [Go by Example: Regular Expressions](https://gobyexample.com/regular-expressions)
- [Go by Example: Slices](https://gobyexample.com/slices)
- [Effective Go: Package Names](https://golang.org/doc/effective_go#package-names)
- [Go by Example: Structs](https://gobyexample.com/structs)

## Dependencies
- TASK-001: Project Setup (completed)

## Estimated Effort
**4-6 hours**

## Related Tasks
- TASK-003: Hexadecimal Conversion (uses tokenizer output)