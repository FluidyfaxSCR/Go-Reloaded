package main

func separatePunctuation(words []string) []string {
	result := []string{}

	for _, word := range words {
		if isPunctuation(word) {
			result = append(result, word)
			continue
		}

		end := 0
		for end < len(word) && isPunctuation(string(word[end])) {
			end = end + 1
		}

		if end == 0 {
			result = append(result, word)
			continue
		}

		punctuation := word[:end]
		rest := word[end:]
		result = append(result, punctuation)
		result = append(result, rest)
	}

	return result
}
