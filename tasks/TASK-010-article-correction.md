# TASK-010: Article Correction

## Objective
Convert `a` to `an` before words beginning with vowels (`a`, `e`, `i`, `o`, `u`) or `h`.

## Acceptance Criteria
- [ ] `a` becomes `an` before vowel-starting words
- [ ] `a` becomes `an` before h-starting words
- [ ] Case-insensitive matching (handles `A` and `a`)
- [ ] No change when not followed by vowel/h words
- [ ] Works with transformed words (after other transformations)

## Key Learning Concepts
- Lookahead processing (checking next token)
- Character classification (vowels vs consonants)
- Case-insensitive string comparison
- Conditional text replacement

## Technical Requirements

### Correction Examples
- `"a amazing"` → `"an amazing"`
- `"a honest"` → `"an honest"`
- `"A elephant"` → `"An elephant"`
- `"a car"` → `"a car"` (no change)

### Implementation Details
- Check next word's first character
- Handle case-insensitive vowel detection
- Process after other transformations (order matters)
- Maintain original case of article

## Test Cases
1. **Vowel A**: "a amazing" → "an amazing"
2. **Vowel E**: "a elephant" → "an elephant"
3. **Vowel I**: "a interesting" → "an interesting"
4. **Vowel O**: "a orange" → "an orange"
5. **Vowel U**: "a umbrella" → "an umbrella"
6. **Letter H**: "a honest" → "an honest"
7. **Uppercase**: "A amazing" → "An amazing"
8. **No Change**: "a car" → "a car"
9. **With Quotes**: "a 'honest (cap)'" → "an 'Honest'"

## Definition of Done
- [ ] All article correction tests pass
- [ ] Case handling works correctly
- [ ] Integration with other transformations verified
- [ ] Vowel and h detection accurate
- [ ] Edge cases handled

## References
- [Go by Example: String Functions](https://gobyexample.com/string-functions)
- [English Grammar: Articles](https://en.wikipedia.org/wiki/Article_(grammar))
- [Unicode Character Classification](https://golang.org/pkg/unicode/)

## Dependencies
- TASK-009: Quote Handling (completed)

## Estimated Effort
**2-3 hours**

## Related Tasks
- TASK-011: Pipeline Integration (combines all transformations)