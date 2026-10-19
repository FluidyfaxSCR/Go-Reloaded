package main

func fixQuotes(words []string) []string {
	result := []string{}
	isOpening := true

	for i := 0; i < len(words); i++ {
		word := words[i]

		if word != "'" {
			result = append(result, word)
			continue
		}

		hasNextWord := i+1 < len(words)
		hasPreviousWord := len(result) > 0

		if isOpening && hasNextWord {
			words[i+1] = "'" + words[i+1]
			isOpening = false
			continue
		}

		if !isOpening && hasPreviousWord {
			last := len(result) - 1
			result[last] = result[last] + "'"
			isOpening = true
			continue
		}

		result = append(result, word)
	}

	return result
}
