package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"unicode/utf8"
)

func main() {
	var lineCounter, wordCounter, characterCounter, byteCounter int
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}

			// err == io.EOF
			if line == "" {
				break
			}
		}

		lineCounter = lineCounter + 1
		wordCounter = wordCounter + countWords(line)
		characterCounter = characterCounter + utf8.RuneCountInString(line)
		byteCounter = byteCounter + len(line)
		fmt.Println(line)
	}

	fmt.Println(lineCounter)
	fmt.Println(wordCounter)
	fmt.Println(characterCounter)
	fmt.Println(byteCounter)
}
