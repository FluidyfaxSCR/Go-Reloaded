package main

import "strconv"

func binToDecimal(words []string) {
	if len(words) == 0 {
		return
	}

	last := len(words) - 1

	number, err := strconv.ParseInt(words[last], 2, 64)
	if err != nil {
		return
	}

	words[last] = strconv.FormatInt(number, 10)
}
