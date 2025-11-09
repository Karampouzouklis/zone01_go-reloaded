# Go-Reloaded Development Progress

**PURPOSE**: This file tracks our step-by-step development progress and serves as a complete session restoration guide. **MUST BE UPDATED** after every major step, task completion, or significant change to maintain continuity across sessions.

## Project Context
**Goal**: Build the go-reloaded text manipulation project step-by-step while learning Go language concepts and implementation methods along the way.

**Learning Approach**: 
- Step-by-step guidance with Go concept explanations
- TDD (Test-Driven Development) methodology
- Minimal code approach - only write what's absolutely necessary
- Follow task-based development from tasks/ directory
- Multi-stage pipeline architecture as per AGENTS.md
- **PRIORITY: Maximum simplicity, readability, modularity, and separation of concerns**
- **HONEST EVALUATION: Critical feedback on design decisions, not just agreement**

**Student Requirements**:
- Learn Go language fundamentals during implementation
- Understand implementation methods and best practices
- Get explanations for each Go concept as we encounter it
- Follow professional development practices

**Key Documentation (Must Follow)**:
- `README.md` - Main project requirements and transformation rules
- `AGENTS.md` - Architecture, TDD requirements, implementation guidance
- `docs/ARCHITECTURE.md` - Technical design decisions and patterns
- `docs/TEST_CASES.md` - Comprehensive test examples and edge cases
- `docs/REQUIREMENTS.md` - Original exercise specifications
- `tasks/TASK-*.md` - Step-by-step implementation roadmap
- `CONTRIBUTING.md` - Development workflow and guidelines

## Current Status: TASK-013 (Performance & Final Validation) - Ready to Start

### Completed Tasks
- ✅ **TASK-001: Project Setup and Basic Structure**
  - Go module initialized (`go.mod` exists)
  - Basic `main.go` with file I/O and error handling
  - Comprehensive test suite in `main_test.go`
  - All tests passing
  - Program builds and runs successfully
  - Basic file copy functionality working

- ✅ **TASK-002: Tokenizer Foundation**
  - Created `tokenizer/` package with Token struct and TokenType enum
  - Implemented comprehensive tokenization logic using regex
  - Handles words, punctuation, whitespace, and transformation markers
  - Supports counted markers like `(up, 2)`
  - All tokenizer tests passing (6 test cases)
  - Successfully integrated with main.go
  - Text reconstruction from tokens working perfectly

- ✅ **TASK-003: Hex Conversion**
  - Created `transform/` package for transformation pipeline
  - Implemented hex to decimal conversion using `strconv.ParseInt(s, 16, 64)`
  - Handles uppercase and lowercase hex digits
  - Gracefully handles invalid hex input (leaves unchanged)
  - Properly removes whitespace and markers after conversion
  - All hex conversion tests passing (6 test cases)
  - Successfully integrated with main.go pipeline

- ✅ **TASK-004: Binary Conversion**
  - **ARCHITECTURAL IMPROVEMENT**: Unified hex and binary conversion into single `processNumberConversion()` function
  - Applied DRY principle - eliminated duplicate code by using dynamic base selection
  - Implemented binary to decimal conversion using `strconv.ParseInt(s, 2, 64)`
  - Handles both hex (base 16) and binary (base 2) in same function
  - All binary conversion tests passing (11 total test cases)
  - Successfully integrated with existing pipeline
  - **Key Learning**: Code refactoring, function generalization, DRY principle

- ✅ **TASK-005: Basic Case Transformations**
  - Implemented `processCaseTransformations()` function for (up), (low), (cap) markers
  - Used switch statement for clean case handling
  - Applied `strings.ToUpper()`, `strings.ToLower()`, manual capitalization
  - Handles edge cases (empty strings, single characters)
  - All case transformation tests passing (6 dedicated + 5 integration tests)
  - Successfully integrated with existing pipeline
  - **Key Learning**: Switch statements, string manipulation, string slicing

- ✅ **TASK-006: Multi-word Case Transformations**
  - Extended `processCaseTransformations()` to handle counted markers (up, 2), (low, 3), (cap, 1)
  - Implemented backward word traversal with bounds checking
  - Added count parameter handling with default to 1 for single words
  - Maintains backward compatibility with existing single-word functionality
  - All multi-word transformation tests passing (4 new + all existing tests)
  - Successfully integrated with existing pipeline
  - **Key Learning**: Loop control with counters, bounds checking, extending existing functions

- ✅ **TASK-007: Punctuation Formatting**
  - Implemented `processPunctuation()` function for .,!?:; spacing rules
  - Removes spaces before punctuation marks
  - Ensures single space after punctuation (when not at end)
  - Handles multiple punctuation marks correctly
  - All punctuation formatting tests passing (6 new test cases)
  - Successfully integrated with existing pipeline
  - **Key Learning**: Token state management, conditional spacing, slice manipulation

- ✅ **TASK-008: Punctuation Groups**
  - Extended `processPunctuation()` to handle punctuation groups (..., !?, !!)
  - Implemented consecutive punctuation collection with lookahead logic
  - Removes spaces within punctuation groups while maintaining external spacing
  - Handles complex scenarios with mixed groups
  - All punctuation group tests passing (5 new test cases)
  - Maintains backward compatibility with basic punctuation formatting
  - **Key Learning**: Lookahead logic, state tracking, pattern recognition

- ✅ **TASK-009: Quote Handling**
  - **ARCHITECTURAL FIX**: Created separate `Quote` token type to distinguish from punctuation
  - Implemented `processQuotes()` function for single quote positioning ('text')
  - Handles quote pairing and removes internal whitespace
  - Processes content inside quotes through transformation pipeline
  - Fixed tokenizer inconsistency - now uses extracted text from regex like other types
  - All quote handling working correctly: `I am ' awesome '` → `I am 'awesome'`
  - Successfully integrated with existing pipeline without order dependencies
  - **Key Learning**: Token type design, regex consistency, separation of concerns

- ✅ **TASK-010: Article Correction**
  - Implemented `processArticles()` function for `a` → `an` conversion before vowels and `h`
  - Handles case-insensitive vowel detection (a, e, i, o, u, h)
  - Maintains original case: `a` → `an`, `A` → `An`
  - Uses lookahead processing to check next word's first character
  - All article correction working correctly: `A amazing rock!` → `An amazing rock!`
  - Successfully integrated with existing pipeline
  - **Key Learning**: Lookahead processing, character classification, case-insensitive comparison

- ✅ **TASK-011: Pipeline Integration**
  - **MAJOR ARCHITECTURAL REFACTORING**: Simplified `processQuotes()` to only handle quote positioning
  - Removed nested transformation calls that violated separation of concerns
  - Fixed commands after quotes by making case transformations skip over quote tokens
  - Established correct processing order: Numbers → Articles → Quotes → Case → Punctuation
  - All core transformations now work together correctly
  - Pipeline integration successful: `'a upper case' (up, 3)` → `'AN UPPER CASE'`
  - **Key Learning**: Pipeline architecture, separation of concerns, token type handling

- ✅ **TASK-012: Complex Integration Testing**
  - **ARCHITECTURAL IMPROVEMENT**: Refactored all transformations to use target-specific processing
  - Changed from "skip certain types" to "only process target types" for consistency
  - All transformations now only look for Word tokens, ignoring all other token types
  - **TEST SUITE OVERHAUL**: Replaced token-level tests with superior string-based integration tests
  - Comprehensive test coverage: All README functional cases and tricky edge cases
  - All TASK-012 requirements satisfied: overlapping transformations, adjacent commands, complex scenarios
  - Edge case resolution: Commands adjacent to quotes now work correctly
  - **DEFENSIVE PROGRAMMING**: Added comprehensive edge case testing with 13 test scenarios
  - Robust error handling: empty strings, unmatched quotes, invalid markers, malformed input
  - Graceful degradation: invalid markers become regular text, no crashes on edge cases
  - **SPACING REFINEMENTS**: Fixed command removal spacing - adds space only when no whitespace exists after command
  - **INVALID COMMAND HANDLING**: Valid commands with invalid input (e.g., `5 (bin)`) consume command but preserve word
  - **Key Learning**: Target-specific processing, integration testing vs unit testing, test design principles, defensive programming

- ✅ **TASK-013: Performance & Final Validation**
  - **PERFORMANCE OPTIMIZATION**: Implemented string builder optimization in main.go
  - Pre-allocation with `builder.Grow(len(content))` for memory efficiency
  - **BENCHMARK RESULTS**: String builder 54x faster than concatenation (3.1μs vs 183μs)
  - **MEMORY IMPROVEMENT**: 260x less memory usage (2.3KB vs 601KB)
  - **ALLOCATION REDUCTION**: 589x fewer allocations (1 vs 589)
  - **PROCESSING SPEED**: 1.23 MB/sec ✅ (exceeds >1MB/sec target)
  - **LARGE FILE TESTING**: Successfully processed 1.1MB file in 0.895 seconds
  - **ERROR HANDLING**: Graceful degradation with malformed input (no crashes)
  - **STRESS TESTING**: Handles files up to 1MB+ without issues
  - **DOCUMENTATION**: Complete performance analysis in docs/PERFORMANCE.md
  - **PRODUCTION READY**: Complete working binary with robust error handling
  - **Key Learning**: Performance optimization, benchmarking, memory profiling, production readiness

### Current Status: PROJECT COMPLETE ✅
**All 13 tasks completed successfully!**

### Files Created/Modified
- `main.go` - Entry point with file I/O, error handling, tokenizer and transform integration
- `main_test.go` - Comprehensive test suite for main functionality
- `tokenizer/tokenizer.go` - Token struct, TokenType enum (Word, Punctuation, Command, Quote, Whitespace), and tokenization logic
- `tokenizer/tokenizer_test.go` - Comprehensive tokenizer tests (7 test cases including quotes)
- `transform/transform.go` - Complete transformation pipeline with target-specific processing
- `transform/transform_test.go` - Comprehensive string-based integration tests (6 test functions, 50+ test cases)
- `transform/edge_cases_test.go` - Defensive programming and error handling tests (13 edge case scenarios)
- `sample.txt`, `result.txt` - Basic test files
- `test_tricky_cases.txt`, `result_tricky_cases.txt` - TASK-012 edge case test files
- `debug_test.txt`, `debug_out.txt` - Debug test files
- `go-reloaded` - Compiled binary

### Go Concepts Learned by Task

#### TASK-001: Project Setup & Foundations
1. **Project Structure**: Go modules (`go.mod`), package main, imports
2. **Error Handling**: `if err != nil` pattern, error wrapping with `fmt.Errorf()`
3. **File I/O**: `os.ReadFile()`, `os.WriteFile()`, file permissions (0644)
4. **Command-line Args**: `os.Args[1:]` for argument parsing and validation
5. **Testing Fundamentals**: `*testing.T`, basic test structure, `t.Errorf()`
6. **Build System**: `go build`, `go run`, `go test -v`
7. **Function Design**: Separating `main()` from `run()` for testability

#### TASK-002: Tokenizer & Advanced Testing
8. **Structs & Enums**: Custom types, struct fields, `iota` for constants
9. **Regular Expressions**: `regexp` package, `FindStringIndex()`, pattern matching
10. **String Manipulation**: `strings` package, `Split()`, `Trim()`, `TrimSpace()`
11. **Advanced Testing**: Table-driven tests, `t.TempDir()`, subtests, `reflect.DeepEqual()`
12. **Type Conversion**: `strconv.Atoi()` for string to integer conversion
13. **Slices**: Dynamic arrays, `append()`, slice operations and indexing
14. **Package Design**: Exported vs unexported functions, clean APIs
15. **Package Organization**: Creating separate packages, import paths

#### TASK-003: Transformations & Pipeline Architecture
16. **Number Base Conversion**: `strconv.ParseInt(s, 16, 64)` for hex to decimal
17. **String Formatting**: `fmt.Sprintf("%d", number)` for formatted output
18. **Loop Control**: `continue` statement, custom loop indexing with `i++`
19. **Slice Manipulation**: Removing elements, `result[:len(result)-1]`
20. **Function Composition**: Building pipelines with function calls
21. **Error Handling in Conversions**: Graceful handling of `ParseInt` errors
22. **Architecture Patterns**: Pipeline pattern, separation of concerns
23. **Modular Testing**: Testing individual functions vs full pipeline

#### TASK-004: Code Refactoring & DRY Principle
24. **DRY Principle**: "Don't Repeat Yourself" - combining similar functions
25. **Function Generalization**: Making functions handle multiple related cases
26. **Conditional Logic**: Dynamic base selection using `if` statements
27. **Code Refactoring**: Improving code structure without changing functionality
28. **Architecture Evolution**: Recognizing patterns and eliminating duplication
29. **Multi-base Conversion**: `strconv.ParseInt(value, base, 64)` with dynamic base

#### TASK-005: String Manipulation & Control Flow
30. **Switch Statements**: Clean alternative to multiple `if-else` statements
31. **String Case Functions**: `strings.ToUpper()`, `strings.ToLower()`
32. **String Slicing**: `word[:1]` and `word[1:]` for accessing parts of strings
33. **String Concatenation**: Combining string parts for capitalization
34. **Edge Case Handling**: Checking string length before slicing operations
35. **Pattern Consistency**: Reusing same function structure across transformations

#### TASK-006: Loop Control & Function Extension
36. **Loop Control with Counters**: Using `wordsTransformed` counter with loop conditions
37. **Bounds Checking**: Preventing array access errors with `k >= 0 && wordsTransformed < count`
38. **Backward Traversal**: Processing tokens in reverse order for multi-word operations
39. **Function Extension**: Adding functionality to existing functions without breaking changes
40. **Default Values**: Using `if count == 0 { count = 1 }` for backward compatibility
41. **Nested Loops**: Combining outer token loop with inner word transformation loop

#### TASK-007: Token State Management & Formatting
42. **Token State Management**: Tracking previous and next tokens for formatting decisions
43. **Conditional Spacing**: Adding spaces based on token types and positions
44. **Slice Manipulation**: Removing elements from end of slice with `result[:len(result)-1]`
45. **Range Loops**: Using `for i, token := range tokens` for index and value access
46. **Boundary Checking**: Using `i < len(tokens)-1` to avoid index out of bounds
47. **Token Creation**: Creating new tokens with `tokenizer.Token{Type: ..., Value: ...}`

#### TASK-008: Pattern Recognition & Lookahead Logic
48. **Lookahead Logic**: Checking future tokens with `j+1 < len(tokens) && tokens[j+1].Type`
49. **Pattern Recognition**: Identifying consecutive punctuation sequences
50. **State Tracking**: Using variables to track processing position across multiple tokens
51. **String Concatenation**: Building punctuation groups with `punctGroup += tokens[j].Value`
52. **Complex Conditionals**: Nested if-else logic for different token type combinations
53. **Index Management**: Manual loop control with `i = j` to skip processed tokens

#### TASK-009: Token Type Design & Architecture Debugging
54. **Token Type Design**: Creating separate token types for different behaviors (`Quote` vs `Punctuation`)
55. **Architecture Debugging**: Identifying and resolving processing conflicts between pipeline stages
56. **Separation of Concerns**: Understanding why different token types need different processing logic
57. **Regex Consistency**: Ensuring all tokenization follows same pattern (extract text, use extracted value, increment by match length)
58. **Pipeline Order Independence**: Designing transformations that work regardless of processing order
59. **Nested Processing**: Processing content inside quotes through the same transformation pipeline
60. **Content Trimming**: Removing leading/trailing whitespace with start/end index management
61. **Token Matching**: Finding matching pairs of tokens (opening/closing quotes) with forward scanning

#### TASK-010: Lookahead Processing & Character Classification
62. **Lookahead Processing**: Scanning forward through tokens to check next word with `for j := i + 1; j < len(tokens)`
63. **Character Classification**: Using `strings.ToLower(string(nextWord[0]))` for vowel detection
64. **Case-insensitive Comparison**: Converting to lowercase for consistent checking regardless of input case
65. **Conditional Text Replacement**: Modifying token values based on conditions (`a` → `an`, `A` → `An`)
66. **String Indexing**: Accessing first character with `nextWord[0]` and handling empty strings
67. **Multiple Condition Checking**: Using OR operators for vowel detection (`firstChar == "a" || firstChar == "e"`)
68. **Case Preservation**: Maintaining original case pattern when making replacements

#### TASK-011: Pipeline Architecture & Cross-Token Processing
69. **Pipeline Architecture**: Designing multi-stage processing with clear separation of concerns
70. **Function Composition**: Chaining transformation functions in correct order for desired results
71. **Cross-Token Processing**: Making transformations work across different token types (words, quotes, whitespace)
72. **Token Type Filtering**: Using `tokens[j].Type != tokenizer.Whitespace && tokens[j].Type != tokenizer.Quote` for selective skipping
73. **Architectural Refactoring**: Removing nested function calls that violate separation of concerns
74. **Processing Order Design**: Understanding how transformation sequence affects final results
75. **Integration Debugging**: Identifying and fixing conflicts between pipeline stages
76. **Token Boundary Handling**: Making transformations work regardless of intervening non-word tokens

#### TASK-012: Target-Specific Processing & Integration Testing
77. **Target-Specific Processing**: Changing from "skip certain types" to "only process target types" approach
78. **Consistent Processing Patterns**: Making all transformation functions use same token filtering logic
79. **Integration Testing vs Unit Testing**: Understanding when to test complete pipeline vs individual functions
80. **String-Based Testing**: Testing user-facing behavior with input/output text pairs instead of internal tokens
81. **Test Design Principles**: Black-box testing for user requirements vs white-box testing for implementation
82. **Edge Case Identification**: Systematically finding and testing boundary conditions and complex interactions
83. **Regression Testing**: Ensuring new changes don't break existing functionality
84. **Test Suite Architecture**: Organizing tests by function and complexity level for maintainability
85. **Defensive Programming**: Writing code that handles unexpected input gracefully without crashing
86. **Error Recovery**: Designing systems that continue functioning when encountering malformed input
87. **Graceful Degradation**: Converting invalid input to safe fallback behavior instead of failing
88. **Boundary Testing**: Testing extreme cases like empty strings, very large counts, invalid formats
89. **Robustness Validation**: Ensuring code doesn't crash on any reasonable input variation
90. **Command Spacing Logic**: Adding spaces only when needed to prevent double spaces while ensuring proper word separation
91. **Invalid Input Classification**: Distinguishing between malformed commands (syntax errors) vs valid commands with invalid input
92. **Conditional Space Insertion**: Using lookahead logic to determine when space insertion is necessary after command removal

#### TASK-013: Performance Optimization & Production Readiness
93. **Performance Benchmarking**: Using `go test -bench` for performance measurement and comparison
94. **Memory Profiling**: Using `-benchmem` and `-memprofile` flags for memory usage analysis
95. **String Builder Optimization**: `strings.Builder` with `Grow()` pre-allocation for efficient string construction
96. **Performance Testing**: Creating realistic performance test scenarios with large inputs
97. **Stress Testing**: Testing with large files (1MB+) to validate scalability and robustness
98. **Production Validation**: Comprehensive testing for production readiness including error scenarios
99. **Benchmark Comparison**: Measuring performance improvements (54x faster string building)
100. **Memory Efficiency**: Understanding Go runtime overhead vs actual memory usage patterns
101. **Large File Handling**: Processing files efficiently without memory leaks or performance degradation
102. **Final Integration Testing**: End-to-end validation of complete system functionality

### Architecture Decisions Made
- **Pipeline Architecture**: Multi-stage transformation pipeline for maximum clarity
- **Separation of Concerns**: Each transformation is a separate function with single responsibility
- **Modularity**: Individual transformation functions can be tested and maintained independently
- **Readability Over Performance**: Prioritizing code clarity and learning over optimization
- **TDD Approach**: Tests first, then minimal implementation
- **Package Organization**: Clean separation - tokenizer, transform, main
- **Honest Feedback**: Critical evaluation of design decisions, not blind agreement

**Performance Trade-offs Acknowledged**:
- Multiple passes through tokens (O(k*n)) instead of single pass (O(n))
- Multiple memory allocations instead of single allocation
- **Justification**: Learning project prioritizes understanding over optimization

### Project Completion Summary
✅ **ALL TASKS COMPLETED** - Go-reloaded is production ready!

**Final Performance Metrics**:
- Processing Speed: 1.23 MB/sec (exceeds target)
- String Builder Optimization: 54x performance improvement
- Large File Support: Successfully handles 1MB+ files
- Error Handling: Graceful degradation, no crashes
- Memory Usage: Reasonable with Go runtime overhead
- Test Coverage: Comprehensive test suite (90+ test cases)
- Code Quality: Clean, modular, well-documented

### Test Results & Coverage
```
=== RUN   TestRun
=== RUN   TestRun/valid_arguments
=== RUN   TestRun/missing_arguments
=== RUN   TestRun/too_many_arguments
=== RUN   TestRun/input_file_not_found
--- PASS: TestRun (0.00s)
PASS
ok  	go-reloaded	0.012s

=== RUN   TestTokenize
=== RUN   TestTokenize/simple_words
=== RUN   TestTokenize/with_punctuation
=== RUN   TestTokenize/transformation_commands
=== RUN   TestTokenize/counted_commands
=== RUN   TestTokenize/hex_command
=== RUN   TestTokenize/multiple_punctuation
=== RUN   TestTokenize/quotes
--- PASS: TestTokenize (0.00s)
PASS
ok  	go-reloaded/tokenizer	0.009s

=== RUN   TestProcessArticles
=== RUN   TestProcessArticles/a_before_vowel_a
=== RUN   TestProcessArticles/a_before_vowel_e
=== RUN   TestProcessArticles/a_before_h
=== RUN   TestProcessArticles/uppercase_A
=== RUN   TestProcessArticles/no_change_for_consonant
=== RUN   TestProcessArticles/all_vowels
--- PASS: TestProcessArticles (0.00s)
PASS
ok  	go-reloaded/transform	0.003s
```

**Coverage Analysis (using `go test -cover ./...`):**
- **Overall**: Excellent coverage across all packages
- **Main package**: Core functionality fully tested
- **Tokenizer**: All token types including Quote type tested
- **Transform**: All transformation functions including article correction tested
- **Quality**: Comprehensive test coverage for learning project

### Current Working Directory
`go-reloaded/` (project root)

### Architectural Philosophy (Established)
**LEARNING-FIRST APPROACH**: This project prioritizes educational value over production optimization.

**Design Principles**:
1. **Maximum Simplicity**: Code should be immediately understandable
2. **Clear Separation**: Each function has exactly one responsibility
3. **Modularity**: Components can be tested and modified independently
4. **Readability**: Self-documenting code over clever optimizations
5. **Pipeline Clarity**: Transformation flow should be obvious

**Acknowledged Trade-offs**:
- Performance: Multiple passes instead of single-pass optimization
- Memory: Multiple allocations instead of pre-allocated buffers
- Complexity: Simple functions over complex state machines

**Justification**: Learning Go fundamentals, design patterns, and best practices is more valuable than micro-optimizations at this stage.

### Last Session Notes
- ✅ Successfully completed TASK-012: Complex Integration Testing
- **MAJOR ARCHITECTURAL IMPROVEMENT**: Refactored all transformations to use target-specific processing
- Changed approach from inconsistent "skip certain token types" to consistent "only process Word tokens"
- **TEST SUITE OVERHAUL**: Replaced inferior token-based tests with superior string-based integration tests
- Restored comprehensive test suite with 6 test functions covering all transformation types
- All TASK-012 tricky edge cases now working: adjacent commands, overlapping transformations, complex scenarios
- **DEFENSIVE PROGRAMMING ADDED**: Created comprehensive edge case test suite with 13 scenarios
- **SPACING REFINEMENTS**: Fixed command removal to add space only when no whitespace exists after command
- **INVALID COMMAND HANDLING**: Valid commands with invalid input consume command but preserve original word
- Edge case insights: Commands work across punctuation (`word! (up)` → `WORD!`), invalid markers become text
- Robust error handling: empty strings, unmatched quotes, invalid counts, malformed input all handled gracefully
- Command spacing: `word(up)next` → `WORD next`, `word (up) next` → `WORD next` (no double spaces)
- Invalid number handling: `5 (bin)` → `5`, `xyz (hex)` → `xyz` (command consumed, word preserved)
- **Key insight**: Target-specific processing eliminates inconsistencies and makes code more robust
- All README functional cases, tricky scenarios, and edge cases passing
- Ready for TASK-013: Performance optimization and final validation

### Next Session Instructions
1. Continue with TASK-013: Performance benchmarking and optimization
2. Test large file handling and memory usage
3. Add performance tests and stress testing
4. Final code cleanup and documentation
5. Maintain step-by-step Go learning approach with explanations
6. **Note**: All core functionality complete - focus on production readiness

### Session Restoration Command
**To continue in a new session, simply say:**
"Please read PROGRESS.md and continue from where we left off"

### Critical Implementation Rules
- **TDD**: Always write tests FIRST before implementation
- **Minimal Code**: Write only what's absolutely necessary (implicit instruction)
- **Step-by-step**: Explain each Go concept as we encounter it
- **Architecture**: Follow multi-stage pipeline with maximum separation of concerns
- **Compliance**: Must satisfy all transformation rules from README.md
- **Quality**: All tests must pass before moving to next task
- **Coverage**: Use `go test -cover ./...` to monitor test coverage (target: >90%)
- **Learning Priority**: Simplicity, readability, modularity over performance optimization
- **Honest Evaluation**: Provide critical feedback, acknowledge trade-offs, explain best practices