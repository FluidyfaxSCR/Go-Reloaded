package main

func hasNegativeTag(words []string) bool {
	for _, word := range words {
		number, isUpTag := getTagNumber(word, "up")
		if isUpTag && number < 0 {
			return true
		}

		number, isLowTag := getTagNumber(word, "low")
		if isLowTag && number < 0 {
			return true
		}

		number, isCapTag := getTagNumber(word, "cap")
		if isCapTag && number < 0 {
			return true
		}
	}

	return false
}
