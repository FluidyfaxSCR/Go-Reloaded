package main

import "strings"

func transformText(text string) string {
	text = strings.ReplaceAll(text, "\r\n", "\n")

	lines := strings.Split(text, "\n")

	for i := 0; i < len(lines); i++ {
		lines[i] = transformLine(lines[i])
	}

	return strings.Join(lines, "\n")
}
