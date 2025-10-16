# Go-Reloaded Project

## Analysis Document

### Problem Description

It is a tool that is given an input text file and outputs a modified text file according to the **modification rules.**

### Modification Rules

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

### Architectures
#### Comparing `Pipeline` and `Streaming FSM (Finite State Machine)`
#### Choice for this project

## Golden Test Set (Success Test Cases)

### FUNCTIONAL TEST CASES

#### Input and expected output

```
If I make you BREAKFAST IN BED (low, 3) just say thank you instead of: how (cap) did you get in my house (up, 2) ?
-> If I make you breakfast in bed just say thank you instead of: How did you get in MY HOUSE?
```

```
I have to pack 101 (bin) outfits. Packed 1a (hex) just to be sure

-> I have to pack 5 outfits. Packed 26 just to be sure
```

```
Don not be sad ,because sad backwards is das . And das not good
->
Don not be sad, because sad backwards is das. And das not good
```

```
harold wilson (cap, 2) : ' I am a optimist ,but a optimist who carries a raincoat . '
->
Harold Wilson: 'I am an optimist, but an optimist who carries a raincoat.'
```


### 5 original tricky examples
### Μεγάλη παράγραφος με πολλούς κανόνες.
