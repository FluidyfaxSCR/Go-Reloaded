package main

import "strings"

func startsWithVowelOrH(word string) bool {
	if word == "" {
		return false
	}

	first := strings.ToLower(word[:1])

	if first == "a" || first == "e" || first == "i" || first == "o" || first == "u" || first == "h" {
		return true
	}

	return false
}
