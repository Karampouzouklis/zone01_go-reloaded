# TASK-006: Multi-word Case Transformations

## Objective
Extend case transformations to handle count parameters: `(up, N)`, `(low, N)`, `(cap, N)`.

## Acceptance Criteria
- [ ] Processes multiple words based on count parameter
- [ ] Handles count greater than available words gracefully
- [ ] Maintains existing single-word functionality
- [ ] Comprehensive test coverage for various counts

## Key Learning Concepts
- Parsing numbers from strings with `strconv.Atoi()`
- Slice manipulation and bounds checking
- Loop control with range and counters
- Edge case handling (bounds checking)
- Extending existing functionality

## Technical Requirements

### Transformation Examples
- `"this is exciting (up, 2)"` → `"this IS EXCITING"`
- `"BREAKFAST IN BED (low, 3)"` → `"breakfast in bed"`
- `"hello world test (cap, 2)"` → `"hello World Test"`

### Implementation Details
- Parse count from marker: `(up, 2)` → count = 2
- Apply transformation to N previous words
- Handle bounds checking (count > available words)
- Extend existing case transformation functions

## Test Cases
1. **Basic Count**: "this is exciting (up, 2)" → "this IS EXCITING"
2. **Full Phrase**: "BREAKFAST IN BED (low, 3)" → "breakfast in bed"
3. **Capitalize Multiple**: "hello world (cap, 2)" → "Hello World"
4. **Count Too Large**: "word (up, 5)" → "WORD" (only available words)
5. **Count Zero**: "word (up, 0)" → "word" (no change)

## Definition of Done
- [ ] All multi-word transformation tests pass
- [ ] Bounds checking prevents errors
- [ ] Backward compatibility with single-word transformations
- [ ] Count parsing robust and error-free
- [ ] Performance scales with word count

## References
- [Go by Example: Range](https://gobyexample.com/range)
- [Go Slices: usage and internals](https://blog.golang.org/slices-intro)
- [strconv.Atoi Documentation](https://pkg.go.dev/strconv#Atoi)

## Dependencies
- TASK-005: Basic Case Transformations (completed)

## Estimated Effort
**3-4 hours**

## Related Tasks
- TASK-007: Basic Punctuation Formatting (next transformation type)