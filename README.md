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

- Go 1.21 or higher

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

See [docs/ARCHITECTURE.md](docs/ARCHITECTURE.md) for detailed architecture analysis and design decisions.

See [docs/TEST_CASES.md](docs/TEST_CASES.md) for comprehensive test cases and examples.



## Project Structure

This project follows professional development practices:

- `docs/` - Architecture and operational documentation
- `tasks/` - Implementation tasks with learning objectives
- `.github/` - GitHub templates and workflows
- `AGENTS.md` - AI agent instructions
- `CONTRIBUTING.md` - Contribution guidelines
- `RELEASE.md` - Release process documentation

See [CONTRIBUTING.md](CONTRIBUTING.md) for detailed development guidelines.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Author

- **Nikos Doulfis** - [Github Profile](https://github.com/Karampouzouklis)

## Project Status

✅ **Complete** - All transformation rules implemented and tested

See `tasks/` directory for implementation progress and learning objectives.

## Acknowledgments

- Built as part of Zone01 curriculum
- Inspired by text processing challenges
