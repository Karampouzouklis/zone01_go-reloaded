# TASK-004: Binary Conversion

## Objective
Implement binary to decimal conversion for `(bin)` markers.

## Acceptance Criteria
- [ ] Converts binary numbers to decimal when followed by `(bin)`
- [ ] Handles only valid binary digits (0, 1)
- [ ] Gracefully handles invalid binary input
- [ ] Reuses conversion pattern from hex implementation
- [ ] Comprehensive test coverage

## Key Learning Concepts
- Binary number system (base 2)
- Reusing `strconv.ParseInt()` with different base
- Pattern recognition for similar functionality
- Code reusability and DRY principle

## Technical Requirements

### Conversion Examples
- `"10 (bin)"` → `"2"`
- `"1010 (bin)"` → `"10"`
- `"11111111 (bin)"` → `"255"`
- `"102 (bin)"` → `"102 (bin)"` (invalid, no change)

### Implementation Details
- Use `strconv.ParseInt(s, 2, 64)` for conversion
- Follow same error handling pattern as hex conversion
- Code reuse where possible
- Maintain consistency with hex transformer

## Test Cases
1. **Valid Binary**: "10 (bin)" → "2"
2. **Longer Binary**: "1010 (bin)" → "10"
3. **All Ones**: "11111111 (bin)" → "255"
4. **Invalid Digits**: "102 (bin)" → "102 (bin)" (unchanged)
5. **Edge Cases**: "0 (bin)" → "0", "1 (bin)" → "1"

## Definition of Done
- [ ] All binary conversion tests pass
- [ ] Code follows same pattern as hex conversion
- [ ] Invalid input handled consistently
- [ ] Performance matches hex conversion
- [ ] Documentation updated

## References
- [Binary Number System](https://en.wikipedia.org/wiki/Binary_number)
- [Go strconv.ParseInt with base 2](https://pkg.go.dev/strconv#ParseInt)
- [DRY Principle](https://en.wikipedia.org/wiki/Don%27t_repeat_yourself)

## Dependencies
- TASK-003: Hexadecimal Conversion (completed)

## Estimated Effort
**1-2 hours** (reusing hex patterns)

## Related Tasks
- TASK-005: Basic Case Transformations (next transformation type)