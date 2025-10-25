# Go-Reloaded Agile Development Tasks

## Overview

You are a senior software architect with expertise in Go and Test Driven Development (TDD).

Your task is to analyze the provided analysis/documentation and generate a list of small, incremental Agile tasks for an entry-level developer using AI agents.

### Task Requirements
Each task must:
- Describe the functionality
- Start with test writing (TDD)
- Include the implementation goal
- End with validation

Ensure the tasks are ordered and collectively lead to full project completion with all tests passing.

---

## Task List

### Task 1: Project Setup and Basic Structure
**Functionality:** Initialize Go project with basic file I/O handling

**Key Learning Concepts:**
- Go project structure and `go mod init`
- Command-line argument parsing with `os.Args`
- File I/O operations: `os.Open()`, `os.Create()`, `ioutil.ReadFile()`
- Error handling patterns in Go
- Writing testable main functions

**TDD Steps:**
- Write test for main function accepting two command-line arguments
- Write test for reading input file and writing to output file
- Write test for error handling (missing files, invalid arguments)

**Implementation Goal:** Create main.go with argument parsing and file operations

**References:**
- [Go by Example: Command Line Arguments](https://gobyexample.com/command-line-arguments)
- [Go by Example: Reading Files](https://gobyexample.com/reading-files)
- [Go Testing Package](https://pkg.go.dev/testing)

**Validation:** Tests pass, program can read/write files without transformation

### Task 2: Tokenizer Foundation
**Functionality:** Split input text into tokens (words, punctuation, markers)

**Key Learning Concepts:**
- String manipulation with `strings` package
- Regular expressions with `regexp` package
- Slice operations and append
- Struct design for token representation
- Package organization in Go

**TDD Steps:**
- Write tests for tokenizing simple text into words
- Write tests for identifying punctuation marks
- Write tests for recognizing transformation markers like (hex), (bin)

**Implementation Goal:** Create tokenizer package that returns slice of tokens

**References:**
- [Go by Example: Regular Expressions](https://gobyexample.com/regular-expressions)
- [Go by Example: Slices](https://gobyexample.com/slices)
- [Effective Go: Package Names](https://golang.org/doc/effective_go#package-names)

**Validation:** Tokenizer correctly identifies all token types from test cases

### Task 3: Hexadecimal Conversion
**Functionality:** Convert hexadecimal numbers to decimal when followed by (hex)

**Key Learning Concepts:**
- Number base conversion (hex to decimal)
- `strconv.ParseInt()` with base parameter
- Error handling for invalid input
- String formatting with `fmt.Sprintf()`

**TDD Steps:**
- Write tests for "1E (hex)" → "30"
- Write tests for "ff (hex)" → "255"
- Write tests for invalid hex handling

**Implementation Goal:** Implement hex transformation stage

**References:**
- [Go by Example: Number Parsing](https://gobyexample.com/number-parsing)
- [strconv Package Documentation](https://pkg.go.dev/strconv)
- [Hexadecimal Number System](https://en.wikipedia.org/wiki/Hexadecimal)

**Validation:** All hex conversion tests pass

### Task 4: Binary Conversion
**Functionality:** Convert binary numbers to decimal when followed by (bin)

**Key Learning Concepts:**
- Binary number system (base 2)
- Reusing `strconv.ParseInt()` with different base
- Pattern recognition for similar functionality
- Code reusability and DRY principle

**TDD Steps:**
- Write tests for "10 (bin)" → "2"
- Write tests for "1010 (bin)" → "10"
- Write tests for invalid binary handling

**Implementation Goal:** Implement binary transformation stage

**References:**
- [Binary Number System](https://en.wikipedia.org/wiki/Binary_number)
- [Go strconv.ParseInt with base 2](https://pkg.go.dev/strconv#ParseInt)

**Validation:** All binary conversion tests pass

### Task 5: Basic Case Transformations
**Functionality:** Handle (up), (low), (cap) for single words

**Key Learning Concepts:**
- String case manipulation: `strings.ToUpper()`, `strings.ToLower()`
- Unicode handling with `strings.Title()` or manual capitalization
- Switch statements for multiple conditions
- Function composition and chaining

**TDD Steps:**
- Write tests for "word (up)" → "WORD"
- Write tests for "WORD (low)" → "word"
- Write tests for "word (cap)" → "Word"

**Implementation Goal:** Implement basic case transformation stage

**References:**
- [Go by Example: String Functions](https://gobyexample.com/string-functions)
- [strings Package Documentation](https://pkg.go.dev/strings)
- [Unicode in Go](https://blog.golang.org/strings)

**Validation:** Single-word case transformations work correctly

### Task 6: Multi-word Case Transformations
**Functionality:** Handle (up, N), (low, N), (cap, N) for multiple words

**Key Learning Concepts:**
- Parsing numbers from strings with `strconv.Atoi()`
- Slice manipulation and bounds checking
- Loop control with range and counters
- Edge case handling (bounds checking)
- Extending existing functionality

**TDD Steps:**
- Write tests for "this is exciting (up, 2)" → "this IS EXCITING"
- Write tests for "BREAKFAST IN BED (low, 3)" → "breakfast in bed"
- Write tests for edge cases (N > available words)

**Implementation Goal:** Extend case transformation to handle counts

**References:**
- [Go by Example: Range](https://gobyexample.com/range)
- [Go Slices: usage and internals](https://blog.golang.org/slices-intro)

**Validation:** Multi-word transformations work with various counts

### Task 7: Basic Punctuation Formatting
**Functionality:** Format single punctuation marks with proper spacing

**Key Learning Concepts:**
- String building and concatenation
- Character classification and sets
- State machine concepts (tracking previous/next tokens)
- String trimming with `strings.TrimSpace()`

**TDD Steps:**
- Write tests for "word ,next" → "word, next"
- Write tests for "word !" → "word!"
- Write tests for all punctuation types (.,!?:;)

**Implementation Goal:** Implement punctuation formatting stage

**References:**
- [Go by Example: String Formatting](https://gobyexample.com/string-formatting)
- [strings.Builder for efficient string building](https://pkg.go.dev/strings#Builder)

**Validation:** Single punctuation marks formatted correctly

### Task 8: Punctuation Groups
**Functionality:** Handle punctuation groups like "..." and "!?"

**Key Learning Concepts:**
- Pattern recognition in sequences
- Lookahead/lookbehind logic
- State tracking across multiple tokens
- Complex conditional logic

**TDD Steps:**
- Write tests for "word ... next" → "word... next"
- Write tests for "What !?" → "What!?"
- Write tests for "word !!" → "word!!"

**Implementation Goal:** Extend punctuation formatter for groups

**References:**
- [Finite State Machines](https://en.wikipedia.org/wiki/Finite-state_machine)
- [Go by Example: If/Else](https://gobyexample.com/if-else)

**Validation:** Punctuation groups formatted without internal spaces

### Task 9: Quote Handling
**Functionality:** Format single quotes around words

**Key Learning Concepts:**
- Paired delimiter handling
- Stack-like data structures for tracking state
- String manipulation within boundaries
- Nested processing (transformations inside quotes)

**TDD Steps:**
- Write tests for "' word '" → "'word'"
- Write tests for "' multiple words '" → "'multiple words'"
- Write tests for nested transformations inside quotes

**Implementation Goal:** Implement quote formatting stage

**References:**
- [Go by Example: Maps](https://gobyexample.com/maps) (for state tracking)
- [Parsing Techniques](https://en.wikipedia.org/wiki/Parsing)

**Validation:** Quotes properly positioned around content

### Task 10: Article Correction
**Functionality:** Convert "a" to "an" before vowels and "h"

**Key Learning Concepts:**
- Lookahead processing (checking next token)
- Character classification (vowels vs consonants)
- Case-insensitive string comparison
- Conditional text replacement

**TDD Steps:**
- Write tests for "a amazing" → "an amazing"
- Write tests for "a honest" → "an honest"
- Write tests for "a car" → "a car" (no change)

**Implementation Goal:** Implement article correction stage

**References:**
- [Go by Example: String Functions](https://gobyexample.com/string-functions)
- [English Grammar: Articles](https://en.wikipedia.org/wiki/Article_(grammar))

**Validation:** Articles correctly adjusted based on following word

### Task 11: Pipeline Integration
**Functionality:** Chain all transformation stages together

**Key Learning Concepts:**
- Pipeline design pattern
- Function composition and chaining
- Interface design for pluggable components
- Integration testing vs unit testing
- Order of operations importance

**TDD Steps:**
- Write tests combining multiple transformations
- Write tests for the provided functional test cases
- Write tests for edge cases and error conditions

**Implementation Goal:** Create pipeline that runs all stages in sequence

**References:**
- [Go by Example: Interfaces](https://gobyexample.com/interfaces)
- [Pipeline Pattern](https://en.wikipedia.org/wiki/Pipeline_(software))
- [Go Testing: Table-driven tests](https://github.com/golang/go/wiki/TableDrivenTests)

**Validation:** All provided test cases pass

### Task 12: Complex Integration Testing
**Functionality:** Handle tricky edge cases and interactions

**Key Learning Concepts:**
- Edge case identification and handling
- Order of operations in complex scenarios
- Debugging complex interactions
- Regression testing
- Code robustness and defensive programming

**TDD Steps:**
- Write tests for transformations inside quotes
- Write tests for overlapping multi-word transformations
- Write tests for markers adjacent to punctuation

**Implementation Goal:** Refine pipeline to handle complex interactions

**References:**
- [Go Testing: Subtests](https://blog.golang.org/subtests)
- [Debugging Go Code](https://golang.org/doc/gdb)
- [Edge Case Testing](https://en.wikipedia.org/wiki/Edge_case)

**Validation:** All tricky test cases pass correctly

### Task 13: Performance and Final Validation
**Functionality:** Ensure robust handling of the complete text sample

**Key Learning Concepts:**
- Performance testing and benchmarking
- Memory profiling with `go test -bench`
- Error recovery and graceful degradation
- Code optimization techniques
- Production readiness checklist

**TDD Steps:**
- Write test using the full transformation rules text sample
- Write tests for malformed input handling
- Write performance tests for reasonable file sizes

**Implementation Goal:** Optimize and finalize the complete solution

**References:**
- [Go by Example: Testing and Benchmarking](https://gobyexample.com/testing-and-benchmarking)
- [Go Performance Tuning](https://github.com/golang/go/wiki/Performance)
- [Profiling Go Programs](https://blog.golang.org/pprof)

**Validation:** Full sample text processes correctly, all tests pass

---

## Summary

Each task builds incrementally on the previous ones, following TDD principles where tests are written first, then minimal implementation to pass those tests, followed by validation that the functionality works as expected.