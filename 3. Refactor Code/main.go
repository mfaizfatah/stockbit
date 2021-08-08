package main

import (
	"fmt"
	"strings"
)

func findFirstStringInBracket(str string) string {
	var (
		result                   string
		indexFirstBracketFound   int
		indexClosingBracketFound int
		wordsAfterFirstBracket   string
	)

	if len(str) > 0 {
		indexFirstBracketFound = strings.Index(str, "(")
	}

	if indexFirstBracketFound >= 0 {
		runes := []rune(str)
		wordsAfterFirstBracket = string(runes[indexFirstBracketFound:len(str)])
		indexClosingBracketFound = strings.Index(wordsAfterFirstBracket, ")")
	}

	if indexClosingBracketFound >= 0 {
		runes := []rune(wordsAfterFirstBracket)
		result = string(runes[1 : indexClosingBracketFound-1])
	}

	return result
}

func main() {
	result := findFirstStringInBracket("Hello (World)")
	fmt.Println(result)
}
