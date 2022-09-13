package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"unicode/utf8"
)

func main() {
	flag.Parse()

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
	}

	fmt.Printf("Number of lines: %d\n", lineCounter)
	fmt.Printf("Number of words: %d\n", wordCounter)
	fmt.Printf("Number of characters: %d\n", characterCounter)
	fmt.Printf("Number of bytes: %d\n", byteCounter)
}
