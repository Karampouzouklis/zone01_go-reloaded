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
- Test all examples from docs/TEST_CASES.md
- Test edge cases and error conditions
- Ensure 100% test coverage for transformation logic
- Follow testing procedures in docs/ops/testing.md

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

## Project Documentation
- `docs/REQUIREMENTS.md` - Original exercise specifications
- `docs/ARCHITECTURE.md` - Technical design decisions and patterns
- `docs/TEST_CASES.md` - Comprehensive test examples and edge cases
- `docs/ops/testing.md` - Testing procedures and quality gates
- `tasks/` - Implementation roadmap with learning objectives
- `CONTRIBUTING.md` - Development workflow and guidelines

## File Structure
```
go-reloaded/
├── main.go              # Entry point and file I/O
├── tokenizer/           # Text tokenization
├── transformers/        # Individual transformation stages
├── formatter/           # Output formatting
├── docs/                # Project documentation
├── tasks/               # Implementation tasks
└── *_test.go           # Test files alongside implementation
```

## Error Handling
- Gracefully handle malformed input
- Continue processing when encountering invalid hex/binary numbers
- Provide meaningful error messages for file I/O issues
- Never panic on user input

## Implementation Guidance
- Follow task-based development approach in tasks/ directory
- Each task includes learning objectives and references
- Refer to docs/ARCHITECTURE.md for design patterns
- Use docs/TEST_CASES.md for validation examples
- Follow TDD cycle: Red → Green → Refactor

## Performance Notes
- Pipeline approach prioritizes maintainability over performance
- Multiple passes acceptable for small input files
- Focus on correctness first, optimize later if needed