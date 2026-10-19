package main

import "strings"

func transformLine(line string) string {
	words := strings.Fields(line)

	if len(words) == 0 {
		return line
	}

	words = separatePunctuation(words)

	words = joinTagWords(words)

	words = separateTagPunctuation(words)

	if hasNegativeTag(words) {
		return "[" + line + "]=err syntax"
	}

	words = applyTags(words)

	words = fixArticles(words)

	words = fixPunctuation(words)

	words = fixQuotes(words)

	return strings.Join(words, " ")
}
