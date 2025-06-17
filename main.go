package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	var filePath string

	for {
		fmt.Println("What is the file path?")
		fmt.Scan(&filePath)

		fileByte, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Println("File path not valid")
			return
		}

		text := string(fileByte)
		lineCount := strings.Count(text, "\n") + 1
		wordCount := len(strings.Fields(text))
		charCount := utf8.RuneCountInString(strings.TrimSpace(text))

		fmt.Println("Word Count:", wordCount)
		fmt.Println("Line Count:", lineCount)
		fmt.Println("Character Count:", charCount)
	}

}
