# Contributing to Go-Reloaded

## Getting Started

1. Fork the repository
2. Clone your fork: `git clone https://github.com/yourusername/go-reloaded.git`
3. Create a feature branch: `git checkout -b feature/your-feature-name`
4. Make your changes following our guidelines below
5. Test your changes: `go test ./...`
6. Commit with clear messages: `git commit -m "Add: feature description"`
7. Push and create a Pull Request

## Development Guidelines

### Code Style
- Follow standard Go conventions (`gofmt`, `golint`)
- Use descriptive variable and function names
- Keep functions small and focused (single responsibility)
- Write self-documenting code with minimal comments

### Testing Requirements
- **Always write tests first** (TDD approach)
- Maintain 100% test coverage for transformation logic
- Test all edge cases and error conditions
- Use table-driven tests for multiple scenarios

### Commit Message Format
```
Type: Brief description

Detailed explanation if needed

- Add specific changes
- Fix specific issues
```

Types: `Add`, `Fix`, `Update`, `Remove`, `Refactor`

### Pull Request Guidelines
- Title format: `[Component] Brief description`
- Include test results in PR description
- Reference any related issues
- Ensure all CI checks pass

## Project Structure

Follow the established pipeline architecture:
- `tokenizer/` - Text tokenization logic
- `transformers/` - Individual transformation stages
- `formatter/` - Output formatting
- Tests alongside implementation files

## Questions?

Open an issue for discussion before starting major changes.