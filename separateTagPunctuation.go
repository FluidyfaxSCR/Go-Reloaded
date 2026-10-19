package main

import "strings"

func separateTagPunctuation(words []string) []string {
	result := []string{}

	for _, word := range words {
		startsLikeTag := strings.HasPrefix(word, "(")
		closeIndex := strings.Index(word, ")")

		if !startsLikeTag || closeIndex == -1 {
			result = append(result, word)
			continue
		}

		if closeIndex == len(word)-1 {
			result = append(result, word)
			continue
		}

		tag := word[:closeIndex+1]
		rest := word[closeIndex+1:]
		result = append(result, tag)
		result = append(result, rest)
	}

	return result
}
