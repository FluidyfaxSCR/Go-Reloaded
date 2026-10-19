package main

import "strings"

func isPunctuation(word string) bool {
	if word == "" {
		return false
	}

	for _, char := range word {
		if !strings.ContainsRune(".,!?:;", char) {
			return false
		}
	}

	return true
}
