# AGENTS.md

## Project Overview
Text manipulation tool that processes input files and applies transformation rules to produce formatted output using a multi-stage pipeline architecture.

## Setup Commands
- Build: `go build -o go-reloaded`
- Run: `./go-reloaded input.txt output.txt`
- Run directly: `go run . input.txt output.txt`
- Test: `go test ./...`
- Test with verbose: `go test -v ./...`

## Architecture
Use multi-stage pipeline approach:
1. **Tokenizer**: Convert raw text into tokens (words, punctuation, transformation markers)
2. **Transformation stages**: Each handles single responsibility (hex/bin conversion, case changes, article correction, quote formatting)
3. **Formatter**: Serialize tokens back to output text

## Code Style
- Follow standard Go conventions
- Use descriptive function and variable names
- Keep functions small and focused (single responsibility)
- Prefer composition over inheritance
- Write self-documenting code with minimal comments

## TDD Requirements
- **Always write tests first** before implementation
- Each transformation rule must have comprehensive unit tests
- Test all examples from README.md
- Test edge cases and error conditions
- Ensure 100% test coverage for transformation logic

## Transformation Rules Implementation
Process in this order:
1. Hex/binary number conversion: `(hex)`, `(bin)`
2. Case transformations: `(up)`, `(low)`, `(cap)` with optional counts
3. Punctuation formatting: proper spacing and grouping
4. Quote handling: `'text'` formatting
5. Article correction: `a` → `an` before vowels/h

## Testing Instructions
- Test each transformation rule independently
- Test combinations of multiple rules
- Test all provided functional test cases
- Test tricky edge cases (transformations inside quotes, overlapping ranges)
- Use table-driven tests for multiple test cases
- Run tests before every commit

## File Structure
```
go-reloaded/
├── main.go              # Entry point and file I/O
├── tokenizer/           # Text tokenization
├── transformers/        # Individual transformation stages
├── formatter/           # Output formatting
└── *_test.go           # Test files alongside implementation
```

## Error Handling
- Gracefully handle malformed input
- Continue processing when encountering invalid hex/binary numbers
- Provide meaningful error messages for file I/O issues
- Never panic on user input

## Performance Notes
- Pipeline approach prioritizes maintainability over performance
- Multiple passes acceptable for small input files
- Focus on correctness first, optimize later if needed