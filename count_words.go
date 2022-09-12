package main

func countWords(text string) int {
	var wordCounter int
	previousCharacter := ' '
	for _, character := range text {
		if previousCharacter != ' ' && character == ' ' {
			wordCounter++
		}

		previousCharacter = character
	}
	if previousCharacter != ' ' {
		wordCounter++
	}

	return wordCounter
}
