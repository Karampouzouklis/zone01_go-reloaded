# Go-Reloaded

A text manipulation tool that processes input files and applies various transformation rules to produce formatted output.

## Installation

```bash
go build -o go-reloaded
```

## Usage

```bash
# Using built binary
./go-reloaded input.txt output.txt

# Or run directly
go run . input.txt output.txt
```

The program takes an input text file (first argument), applies transformation rules, and writes the result to an output file (second argument).

### Example

```bash
$ echo "Simply add 42 (hex) and 10 (bin) and you will see the result is 68." > sample.txt
$ go run . sample.txt result.txt
$ cat result.txt
Simply add 66 and 2 and you will see the result is 68.
```

## Requirements

- Go 1.24.7 or higher

## Testing

```bash
go test ./...
```

## Features

### Transformation Rules

- Every instance of `(hex)` should replace the word before with the decimal version of the word (in this case the word will always be a hexadecimal number). (Ex: "1E (hex) files were added" -> "30 files were added")
- Every instance of `(bin)` should replace the word before with the decimal version of the word (in this case the word will always be a binary number). (Ex: "It has been 10 (bin) years" -> "It has been 2 years")
- Every instance of `(up)` converts the word before with the Uppercase version of it. (Ex: "Ready, set, go (up) !" -> "Ready, set, GO!")
- Every instance of `(low)` converts the word before with the Lowercase version of it. (Ex: "I should stop SHOUTING (low)" -> "I should stop shouting")
- Every instance of `(cap)` converts the word before with the capitalized version of it. (Ex: "Welcome to the Brooklyn bridge (cap)" -> "Welcome to the Brooklyn Bridge")

  - For `(low)`, `(up)`, `(cap)` if a number appears next to it, like so: `(low, <number>)` it turns the previously specified number of words in lowercase, uppercase or capitalized accordingly. (Ex: "This is so exciting (up, 2)" -> "This is SO EXCITING")

- Every instance of the punctuations `.`, `,`, `!`, `?`, `:` and `;` should be close to the previous word and with space apart from the next one. (Ex: "I was sitting over there ,and then BAMM !!" -> "I was sitting over there, and then BAMM!!").
  - Except if there are groups of punctuation like: `...` or `!?`. In this case the program should format the text as in the following example: "I was thinking ... You were right" -> "I was thinking... You were right".
- The punctuation mark `'` will always be found with another instance of it and they should be placed to the right and left of the word in the middle of them, without any spaces. (Ex: "I am exactly how they describe me: ' awesome '" -> "I am exactly how they describe me: 'awesome'")
  - If there are more than one word between the two `' '` marks, the program should place the marks next to the corresponding words (Ex: "As Elton John said: ' I am the most well-known homosexual in the world '" -> "As Elton John said: 'I am the most well-known homosexual in the world'")
- Every instance of `a` should be turned into `an` if the next word begins with a vowel (`a`, `e`, `i`, `o`, `u`) or a `h`. (Ex: "There it was. A amazing rock!" -> "There it was. An amazing rock!").

## Architecture Analysis

### Comparing `Pipeline` and `Streaming FSM (Finite State Machine)`

Lorem Ipsuim

### Choice for this project
Lorem Ipsum

## Test Cases

### FUNCTIONAL TEST CASES

#### Input and expected output

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

### ORIGINAL TRICKY EXAMPLES

#### 1. Transformation markers inside quotes
```
Input: He said ' I want 1a (hex) apples (up, 2) ' today .
Expected: He said 'I want 26 APPLES' today.
```
**Why tricky:** Tests whether transformation markers `(hex)` and `(up, 2)` work when they appear inside quotes. The implementation must handle quote pairing while still processing transformations within the quoted text.

#### 2. Multiple consecutive punctuation groups with spacing
```
Input: Wait ... What !? No way !! Are you serious ... ?
Expected: Wait... What!? No way!! Are you serious...?
```
**Why tricky:** Tests complex punctuation grouping rules. `...` and `!?` are groups (no space before), but `!!` and `...?` need to be handled correctly. The challenge is identifying what constitutes a "group" versus separate punctuation.

#### 3. Overlapping multi-word transformations
```
Input: Make these THREE WORDS (low, 2) very (up, 3) important
Expected?: Make these THREE WORDS VERY important -OR- Make these three words VERY important
```
**Why tricky:** The `(low, 2)` affects "THREE WORDS" and `(up, 3)` affects "three WORDS VERY". This tests how overlapping transformation ranges are handled, in what order will the transformations be executed?

#### 4. Nested quotes ambiguity
```
Input: She said ' a elephant , a ' unusual ' animal , is here ' .
Expected?: She said 'an elephant, an' unusual 'animal, is here'. -OR- She said 'an elephant, an 'unusual' animal, is here'.
```
**Why tricky:** The instructions don't specify how to handle nested quotes. Should they be treated as one outer pair with inner text, or as separate quote pairs?

### TEXT WITH TRANSFORMATION RULES
```
When developing a text processing application ,you must understand that ff (hex) different algorithms exist for parsing . A experienced (cap) programmer knows that 1010 (bin) main approaches (up, 2) can be used: streaming parsers and batch processors . The streaming approach processes data ' character by character ' while batch processing loads everything into memory first ... However ,both methods have trade-offs ! Streaming uses less memory but batch processing is often faster ? The choice depends on your specific requirements : memory constraints ,performance needs ,and data size . For instance ,if you need to process a HUGE file containing 1a (hex) million records ,streaming might be BETTER (low, 2) for memory usage . Remember that a efficient algorithm should handle edge cases gracefully ... What happens when you encounter malformed input !? Your parser should be ROBUST ENOUGH (low, 2) to continue processing . As donald knuth (cap, 2) said : ' premature optimization is the root of all evil ' in programming . Focus on correctness first ,then optimize . Whether you choose 101 (bin) different data structures or stick to basic arrays ,make sure your code is readable and maintainable . Additionally ,consider that 1f (hex) different optimization techniques (up, 2) exist for performance tuning ... Some developers prefer a object-oriented (cap, 3) approach while others favor functional programming !? The key is understanding your data flow : input validation ,transformation logic ,and output formatting . Modern compilers can optimize 1100 (bin) percent of simple operations automatically . However ,complex algorithms still require careful design ... Remember that a elegant (up) solution often beats a brute-force approach ! When working with large datasets containing abc (hex) thousand entries ,consider using parallel processing . As Linus Torvalds once said : ' talk is cheap ,show me the code ' . Whether you implement 11 (bin) different parsing strategies or just one robust solution ,always test thoroughly with edge cases .
```

## Building from Source

```bash
git clone <repository-url>
cd go-reloaded
go build -o go-reloaded
```

