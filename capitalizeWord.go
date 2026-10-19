package main

import (
	"strings"
	"unicode"
)

func capitalizeWord(word string) string {
	if word == "" {
		return word
	}

	letters := []rune(strings.ToLower(word))
	letters[0] = unicode.ToUpper(letters[0])

	return string(letters)
}
