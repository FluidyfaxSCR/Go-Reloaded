package main

func fixArticles(words []string) []string {
	for i := 0; i < len(words)-1; i++ {
		next := words[i+1]

		if !startsWithVowelOrH(next) {
			continue
		}

		if words[i] == "a" {
			words[i] = "an"
		}

		if words[i] == "A" {
			if isUpperCase(next) {
				words[i] = "AN"
			} else {
				words[i] = "An"
			}
		}
	}

	return words
}
