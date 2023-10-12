package main

import (
	"fmt"
	"regexp"
)

func MatchNameRegex() {
	// Input text
	text := "Mỹ Đình and mĩ Đình are places Đình mĩ."

	// Regular expression pattern to match "Mỹ Đình" and "mĩ Đình" case-insensitively
	pattern := "(?i)M[\\p{Ll}\\p{Lu}]+ Đình"

	// Compile the regex pattern
	regex, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return
	}

	// Find all matches in the text
	matches := regex.FindAllString(text, -1)

	// Print the matches
	for _, match := range matches {
		fmt.Println("Match:", match)
	}
}
