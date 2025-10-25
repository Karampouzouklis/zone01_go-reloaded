# TASK-005: Basic Case Transformations

## Objective
Implement single-word case transformations for `(up)`, `(low)`, and `(cap)` markers.

## Acceptance Criteria
- [ ] `(up)` converts word to uppercase
- [ ] `(low)` converts word to lowercase  
- [ ] `(cap)` capitalizes first letter of word
- [ ] Handles Unicode characters correctly
- [ ] Comprehensive test coverage for all cases

## Key Learning Concepts
- String case manipulation: `strings.ToUpper()`, `strings.ToLower()`
- Unicode handling with `strings.Title()` or manual capitalization
- Switch statements for multiple conditions
- Function composition and chaining

## Technical Requirements

### Transformation Examples
- `"word (up)"` → `"WORD"`
- `"WORD (low)"` → `"word"`
- `"word (cap)"` → `"Word"`
- `"hello (cap)"` → `"Hello"`

### Implementation Details
- Use `strings.ToUpper()`, `strings.ToLower()`
- For capitalization: `strings.ToUpper(s[:1]) + strings.ToLower(s[1:])`
- Handle empty strings and single characters
- Process single words only (no count parameter yet)

## Test Cases
1. **Uppercase**: "word (up)" → "WORD"
2. **Lowercase**: "WORD (low)" → "word"
3. **Capitalize**: "word (cap)" → "Word"
4. **Mixed Case**: "WoRd (low)" → "word"
5. **Single Char**: "a (up)" → "A"
6. **Unicode**: "café (up)" → "CAFÉ"

## Definition of Done
- [ ] All case transformation tests pass
- [ ] Unicode handling verified
- [ ] Edge cases (empty, single char) handled
- [ ] Performance acceptable
- [ ] Code follows project patterns

## References
- [Go by Example: String Functions](https://gobyexample.com/string-functions)
- [strings Package Documentation](https://pkg.go.dev/strings)
- [Unicode in Go](https://blog.golang.org/strings)
- [Go by Example: Switch](https://gobyexample.com/switch)

## Dependencies
- TASK-004: Binary Conversion (completed)

## Estimated Effort
**2-3 hours**

## Related Tasks
- TASK-006: Multi-word Case Transformations (extends this functionality)