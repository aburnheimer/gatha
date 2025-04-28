# Gāthā

Gāthā is a Sanskrit term for 'song' or 'verse', especially referring to any poetic meter which is used in legends or folklores. The word is originally derived from the Sanskrit/Prakrit root `gai`, which means 'to speak, sing, recite or extol'. This project provides a simple library for converting plain text poetry verse into JSON objects and vice versa.

## Project Structure

```
text-json-converter
├── pkg
│   └── gatha
│       ├── word.go      # A struct to represent words in poetry
│       └── word_test.go # Unit tests for the word struct
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

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue for any suggestions or improvements.

## License

This project is licensed under the MIT License. See the LICENSE file for more details.