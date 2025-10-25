# TASK-008: Punctuation Groups

## Objective
Handle punctuation groups like `...`, `!?`, `!!` with no internal spaces.

## Acceptance Criteria
- [ ] Groups like `...` and `!?` formatted without internal spaces
- [ ] Groups attached to previous word (no space before)
- [ ] Single space after groups (before next word)
- [ ] Maintains basic punctuation formatting from TASK-007

## Key Learning Concepts
- Pattern recognition in sequences
- Lookahead/lookbehind logic
- State tracking across multiple tokens
- Complex conditional logic

## Technical Requirements

### Formatting Examples
- `"word ... next"` → `"word... next"`
- `"What !?"` → `"What!?"`
- `"word !!"` → `"word!!"`
- `"thinking ... really ?"` → `"thinking... really?"`

### Implementation Details
- Identify consecutive punctuation marks
- Remove spaces within punctuation groups
- Apply group formatting rules
- Extend punctuation formatter from TASK-007

## Test Cases
1. **Ellipsis**: "word ... next" → "word... next"
2. **Question-Exclamation**: "What !?" → "What!?"
3. **Double Exclamation**: "word !!" → "word!!"
4. **Mixed Groups**: "Wait ... What !?" → "Wait... What!?"
5. **Complex**: "thinking ... really !?" → "thinking... really!?"

## Definition of Done
- [ ] All punctuation group tests pass
- [ ] Groups formatted without internal spaces
- [ ] Backward compatibility with basic punctuation
- [ ] Complex scenarios handled correctly
- [ ] Performance acceptable

## References
- [Finite State Machines](https://en.wikipedia.org/wiki/Finite-state_machine)
- [Go by Example: If/Else](https://gobyexample.com/if-else)
- [Pattern Matching Techniques](https://en.wikipedia.org/wiki/Pattern_matching)

## Dependencies
- TASK-007: Basic Punctuation Formatting (completed)

## Estimated Effort
**3-4 hours**

## Related Tasks
- TASK-009: Quote Handling (different punctuation type)