package main

import "strings"

func toLowerWords(words []string, count int) {
	changed := 0

	for i := len(words) - 1; i >= 0; i-- {
		if changed == count {
			break
		}
		if isPunctuation(words[i]) || words[i] == "'" {
			continue
		}

		words[i] = strings.ToLower(words[i])
		changed = changed + 1
	}
}
