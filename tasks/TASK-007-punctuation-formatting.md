# TASK-007: Basic Punctuation Formatting

## Objective
Implement proper spacing for punctuation marks: `.`, `,`, `!`, `?`, `:`, `;`.

## Acceptance Criteria
- [ ] Punctuation attached to previous word (no space before)
- [ ] Single space after punctuation (before next word)
- [ ] Handles multiple punctuation marks
- [ ] Does not affect punctuation groups (handled in TASK-008)

## Key Learning Concepts
- String building and concatenation
- Character classification and sets
- State machine concepts (tracking previous/next tokens)
- String trimming with `strings.TrimSpace()`

## Technical Requirements

### Formatting Examples
- `"word ,next"` → `"word, next"`
- `"word !"` → `"word!"`
- `"hello , world !"` → `"hello, world!"`

### Implementation Details
- Remove spaces before punctuation
- Ensure single space after punctuation
- Handle sentence endings and mid-sentence punctuation
- Process during formatting stage

## Test Cases
1. **Comma**: "word ,next" → "word, next"
2. **Exclamation**: "word !" → "word!"
3. **Question**: "really ?" → "really?"
4. **Colon**: "note :" → "note:"
5. **Multiple**: "hello , world !" → "hello, world!"
6. **End of Text**: "done ." → "done."

## Definition of Done
- [ ] All punctuation formatting tests pass
- [ ] Handles all specified punctuation types
- [ ] No regression in existing functionality
- [ ] Proper spacing maintained
- [ ] Edge cases handled

## References
- [Go by Example: String Formatting](https://gobyexample.com/string-formatting)
- [strings.Builder for efficient string building](https://pkg.go.dev/strings#Builder)
- [Character Sets in Go](https://golang.org/pkg/unicode/)

## Dependencies
- TASK-006: Multi-word Case Transformations (completed)

## Estimated Effort
**3-4 hours**

## Related Tasks
- TASK-008: Punctuation Groups (extends this functionality)