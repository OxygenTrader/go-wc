package main

import (
	"testing"
)

func Test_countWords(test *testing.T) {
	text := "one two three"
	actualWordCounter := countWords(text)

	expectedWordCounter := 3
	if actualWordCounter != expectedWordCounter {
		test.Fail()
	}
}

func Test_countWords_withDuplicatedSpaces(test *testing.T) {
	text := "one two  three"
	actualWordCounter := countWords(text)

	expectedWordCounter := 3
	if actualWordCounter != expectedWordCounter {
		test.Fail()
	}
}

func Test_countWords_withPrependedSpaces(test *testing.T) {
	text := " one two three"
	actualWordCounter := countWords(text)

	expectedWordCounter := 3
	if actualWordCounter != expectedWordCounter {
		test.Fail()
	}
}

func Test_countWords_withTrailingSpaces(test *testing.T) {
	text := "one two three "
	actualWordCounter := countWords(text)

	expectedWordCounter := 3
	if actualWordCounter != expectedWordCounter {
		test.Fail()
	}
}
