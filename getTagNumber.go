package main

import (
	"strconv"
	"strings"
)

func getTagNumber(word string, name string) (int, bool) {
	start := "(" + name

	if !strings.HasPrefix(word, start) {
		return 0, false
	}
	if !strings.HasSuffix(word, ")") {
		return 0, false
	}

	middle := strings.TrimPrefix(word, start)
	middle = strings.TrimSuffix(middle, ")")
	middle = strings.TrimPrefix(middle, ",")

	if middle == "" {
		return 0, false
	}

	number, err := strconv.Atoi(middle)
	if err != nil {
		return 0, false
	}

	return number, true
}
