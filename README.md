# go-wc

[![Go Report Card](https://goreportcard.com/badge/github.com/OxygenTrader/go-wc)](https://goreportcard.com/report/github.com/OxygenTrader/go-wc)

A clone of the [Unix wc tool](<https://en.wikipedia.org/wiki/Wc_(Unix)>).

## Installation

```
$ go install github.com/OxygenTrader/go-wc@latest
```

## Usage

```
$ go-wc -h | -help | --help
$ go-wc
```

Stdin: text in the [UTF-8 encoding](https://en.wikipedia.org/wiki/UTF-8).

Options:

- `-h`, `-help`, `--help` &mdash; show the help message and exit.

## Output Example

```
Number of lines: 10
Number of words: 10
Number of characters: 21
Number of bytes: 21
```

## License

The MIT License (MIT)

Copyright &copy; 2022 OxygenTrader
