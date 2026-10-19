package main

func fixPunctuation(words []string) []string {
	result := []string{}

	for _, word := range words {
		hasPreviousWord := len(result) > 0

		if isPunctuation(word) && hasPreviousWord {
			last := len(result) - 1
			result[last] = result[last] + word
			continue
		}

		result = append(result, word)
	}

	return result
}
