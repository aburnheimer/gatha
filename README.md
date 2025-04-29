# Gāthā

Gāthā is a Sanskrit term for 'song' or 'verse', especially referring to any poetic meter which is used in legends or folklores. The word is originally derived from the Sanskrit/Prakrit root `gai`, which means 'to speak, sing, recite or extol'. This project provides a simple library for converting plain text poetry verse into JSON objects and vice versa.

## Project Structure

```
gatha
├── test
|   └── expected_....json  # Various test artifacts used in unit testing
├── ....go                 # Various structs representing aspects of poetic meter
├── ..._test.go            # Unit tests for those structs
├── go.mod                 # Module definition
└── README.md              # Project documentation
```

## Installation

To install the project, clone the repository and navigate to the project directory:

```
git clone https://github.com/aburnheimer/gatha.git
cd gatha
```

Then, run:

```
go mod tidy
```

## Usage

### Verse

A poem's structure (Verse) is made up of Stanzas, which are groupings of Lines, and these Lines are themselves made up of Words. The library can import a Verse as simple text, where Lines are separated by single linefeeds (`\n`), and Stanzas are separated by double linefeeds (`\n\n`). This example has two Stanzas, each having two Lines, and several Words in all of those Lines:

```
The quick brown fox jumps over the lazy dog
and in to the box, right next to the log.

Her coat of reddish hues,
the most lovely of views.
```

Each of the objects can be annotated accordingly, e.g. Words can have Rhyme info attached, Lines having a count of syllables, Stanzas having a rhyme pattern, etc. Please see embedded documentation for more info.

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue for any suggestions or improvements.

## License

This project is licensed under the MIT License. See the LICENSE file for more details.