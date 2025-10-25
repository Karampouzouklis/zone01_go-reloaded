# Test Cases Documentation

## FUNCTIONAL TEST CASES

### Input and expected output

```
Input: If I make you BREAKFAST IN BED (low, 3) just say thank you instead of: how (cap) did you get in my house (up, 2) ?
Expected: If I make you breakfast in bed just say thank you instead of: How did you get in MY HOUSE?
```

```
Input: I have to pack 101 (bin) outfits. Packed 1a (hex) just to be sure
Expected: I have to pack 5 outfits. Packed 26 just to be sure
```

```
Input: Don not be sad ,because sad backwards is das . And das not good
Expected: Don not be sad, because sad backwards is das. And das not good
```

```
Input: harold wilson (cap, 2) : ' I am a optimist ,but a optimist who carries a raincoat . '
Expected: Harold Wilson: 'I am an optimist, but an optimist who carries a raincoat.'
```

## ORIGINAL TRICKY EXAMPLES

### 1. Transformation markers inside quotes
```
Input: He said ' I want 1a (hex) apples (up, 2) ' today .
Expected: He said 'I want 26 APPLES' today.
```
**Why tricky:** Tests whether transformation markers `(hex)` and `(up, 2)` work when they appear inside quotes. The implementation must handle quote pairing while still processing transformations within the quoted text.

### 2. Multiple consecutive punctuation groups with spacing
```
Input: Wait ... What !? No way !! Are you serious ... ?
Expected: Wait... What!? No way!! Are you serious...?
```
**Why tricky:** Tests complex punctuation grouping rules. `...` and `!?` are groups (no space before), but `!!` and `...?` need to be handled correctly. The challenge is identifying what constitutes a "group" versus separate punctuation.

### 3. Overlapping multi-word transformations
```
Input: Make these THREE WORDS (low, 2) very (up, 3) important
Expected?: Make these THREE WORDS VERY important -OR- Make these three words VERY important
```
**Why tricky:** The `(low, 2)` affects "THREE WORDS" and `(up, 3)` affects "three WORDS VERY". This tests how overlapping transformation ranges are handled, in what order will the transformations be executed?

### 4. Article + quoted word + transformation
```
Input: It was a 'honest (cap)' decision .
Expected: It was an 'Honest' decision.
```
**Why tricky:** The article correction (`a` → `an`) must look at the next word even when it's inside quotes and will be modified by a marker `(cap)`. Quotes remove internal spaces and the capitalization marker changes the first letter — the implementation must apply markers and quote formatting in the right order so the article rule sees the correct next-word initial.

### 5. Markers adjacent to punctuation / no-space cases
```
Input: Make it loud(up)!
Expected: Make it LOUD!
```
**Why tricky:** Ensures markers apply to the correct preceding token even when followed immediately by punctuation (no separating space). Also verifies that punctuation attaches to the transformed token correctly.

## COMPREHENSIVE TEST TEXT

### TEXT WITH TRANSFORMATION RULES
```
When developing a text processing application ,you must understand that ff (hex) different algorithms exist for parsing . A experienced (cap) programmer knows that 1010 (bin) main approaches (up, 2) can be used: streaming parsers and batch processors . The streaming approach processes data ' character by character ' while batch processing loads everything into memory first ... However ,both methods have trade-offs ! Streaming uses less memory but batch processing is often faster ? The choice depends on your specific requirements : memory constraints ,performance needs ,and data size . For instance ,if you need to process a HUGE file containing 1a (hex) million records ,streaming might be BETTER (low, 2) for memory usage . Remember that a efficient algorithm should handle edge cases gracefully ... What happens when you encounter malformed input !? Your parser should be ROBUST ENOUGH (low, 2) to continue processing . As donald knuth (cap, 2) said : ' premature optimization is the root of all evil ' in programming . Focus on correctness first ,then optimize . Whether you choose 101 (bin) different data structures or stick to basic arrays ,make sure your code is readable and maintainable . Additionally ,consider that 1f (hex) different optimization techniques (up, 2) exist for performance tuning ... Some developers prefer a object-oriented (cap, 3) approach while others favor functional programming !? The key is understanding your data flow : input validation ,transformation logic ,and output formatting . Modern compilers can optimize 1100 (bin) percent of simple operations automatically . However ,complex algorithms still require careful design ... Remember that a elegant (up) solution often beats a brute-force approach ! When working with large datasets containing abc (hex) thousand entries ,consider using parallel processing . As Linus Torvalds once said : ' talk is cheap ,show me the code ' . Whether you implement 11 (bin) different parsing strategies or just one robust solution ,always test thoroughly with edge cases .
```

## Test Implementation Guidelines

### Unit Tests
- Each transformation rule should have dedicated unit tests
- Test both valid and invalid inputs
- Verify edge cases and boundary conditions

### Integration Tests
- Test complete pipeline with functional test cases
- Verify complex interactions between rules
- Test all tricky examples

### Performance Tests
- Benchmark processing speed with large inputs
- Monitor memory usage during processing
- Verify no memory leaks in long-running tests