package main

import "strings"

func joinTagWords(words []string) []string {
	result := []string{}

	for i := 0; i < len(words); i++ {
		word := words[i]

		startsLikeTag := strings.HasPrefix(word, "(")
		isClosed := strings.Contains(word, ")")
		hasNextWord := i+1 < len(words)

		if startsLikeTag && !isClosed && hasNextWord {
			word = word + words[i+1]
			i = i + 1
		}

		result = append(result, word)
	}

	return result
}
