# Architecture Documentation

## Overview

Go-Reloaded implements a multi-stage pipeline architecture for text transformation processing.

## Design Principles

### Pipeline Pattern
- **Separation of Concerns**: Each stage handles one responsibility
- **Testability**: Individual stages can be unit tested in isolation
- **Maintainability**: Easy to add, remove, or modify transformation rules
- **Readability**: Clear data flow from input to output

### Architecture Diagram

```
Input Text → Tokenizer → Transformers → Formatter → Output Text
                ↓           ↓            ↓
            [Token[]]   [Token[]]   [string]
```

## Components

### 1. Tokenizer
**Purpose**: Convert raw text into structured tokens
- Identifies words, punctuation, and transformation markers
- Returns slice of Token structs
- Handles whitespace and special characters

### 2. Transformers (Pipeline Stages)
**Purpose**: Apply transformation rules to tokens

**Processing Order**:
1. **Number Conversion**: `(hex)`, `(bin)` → decimal
2. **Case Transformation**: `(up)`, `(low)`, `(cap)` with optional counts
3. **Punctuation Formatting**: Spacing and grouping rules
4. **Quote Handling**: `'text'` formatting
5. **Article Correction**: `a` → `an` before vowels/h

### 3. Formatter
**Purpose**: Convert processed tokens back to formatted text
- Handles spacing between tokens
- Applies final formatting rules
- Returns clean output string

## Data Structures

### Token
```go
type Token struct {
    Type    TokenType
    Value   string
    Marker  string  // For transformation markers
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

## Design Decisions

### Why Pipeline over FSM?
- **Simplicity**: Easier to understand and debug
- **Modularity**: Each stage is independent
- **Testing**: Individual components can be tested separately
- **Maintenance**: Adding new rules doesn't affect existing ones

### Performance Considerations
- Multiple passes over data (acceptable for small inputs)
- Memory usage: O(n) where n is input size
- Time complexity: O(n) per transformation stage

## Extension Points

### Adding New Transformations
1. Create new transformer function
2. Add to pipeline in correct order
3. Write comprehensive tests
4. Update documentation

### Error Handling Strategy
- Graceful degradation for invalid input
- Continue processing on non-critical errors
- Log warnings for malformed markers
- Never panic on user input