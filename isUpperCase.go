package main

import "strings"

func isUpperCase(word string) bool {
	if len(word) < 2 {
		return false
	}

	return word == strings.ToUpper(word)
}
