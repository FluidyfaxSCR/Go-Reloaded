package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("usage: go run . <input-file> <output-file>")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	content, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	result := transformText(string(content))

	err = os.WriteFile(outputFile, []byte(result), 0644)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}
