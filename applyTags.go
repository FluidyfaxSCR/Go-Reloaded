package main

func applyTags(words []string) []string {
	result := []string{}

	for _, word := range words {
		if word == "(hex)" {
			hexToDecimal(result)
			continue
		}

		if word == "(bin)" {
			binToDecimal(result)
			continue
		}

		if word == "(up)" {
			toUpperWords(result, 1)
			continue
		}

		if word == "(low)" {
			toLowerWords(result, 1)
			continue
		}

		if word == "(cap)" {
			capitalizeWords(result, 1)
			continue
		}

		number, isUpTag := getTagNumber(word, "up")
		if isUpTag {
			toUpperWords(result, number)
			continue
		}

		number, isLowTag := getTagNumber(word, "low")
		if isLowTag {
			toLowerWords(result, number)
			continue
		}

		number, isCapTag := getTagNumber(word, "cap")
		if isCapTag {
			capitalizeWords(result, number)
			continue
		}

		result = append(result, word)
	}

	return result
}
