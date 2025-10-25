# TASK-003: Hexadecimal Conversion

## Objective
Implement hexadecimal to decimal conversion for `(hex)` markers.

## Acceptance Criteria
- [ ] Converts hex numbers to decimal when followed by `(hex)`
- [ ] Handles both uppercase and lowercase hex digits
- [ ] Gracefully handles invalid hex input
- [ ] Maintains original text for invalid conversions
- [ ] Comprehensive test coverage including edge cases

## Key Learning Concepts
- Number base conversion (hex to decimal)
- `strconv.ParseInt()` with base parameter
- Error handling for invalid input
- String formatting with `fmt.Sprintf()`

## Technical Requirements

### Conversion Examples
- `"1E (hex)"` → `"30"`
- `"ff (hex)"` → `"255"`
- `"ABC (hex)"` → `"2748"`
- `"invalid (hex)"` → `"invalid (hex)"` (no change)

### Implementation Details
- Use `strconv.ParseInt(s, 16, 64)` for conversion
- Handle conversion errors gracefully
- Process tokens from tokenizer output
- Return modified token slice

## Test Cases
1. **Valid Hex**: "1E (hex)" → "30"
2. **Lowercase**: "ff (hex)" → "255"  
3. **Large Numbers**: "FFFF (hex)" → "65535"
4. **Invalid Input**: "xyz (hex)" → "xyz (hex)" (unchanged)
5. **Edge Cases**: "0 (hex)" → "0", "" (hex)" → "" (hex)"

## Definition of Done
- [ ] All hex conversion tests pass
- [ ] Invalid input handled without errors
- [ ] Integration with tokenizer verified
- [ ] Performance acceptable for typical inputs
- [ ] Error logging for invalid conversions

## References
- [Go by Example: Number Parsing](https://gobyexample.com/number-parsing)
- [strconv Package Documentation](https://pkg.go.dev/strconv)
- [Hexadecimal Number System](https://en.wikipedia.org/wiki/Hexadecimal)

## Dependencies
- TASK-002: Tokenizer Foundation (completed)

## Estimated Effort
**2-3 hours**

## Related Tasks
- TASK-004: Binary Conversion (similar implementation pattern)