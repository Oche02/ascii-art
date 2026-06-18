package main

import (
	"strings"
)

func GenerateArt(text string, banner map[rune][]string) string {
	var result strings.Builder
	word := SplitInput(text)
	for i, words := range word {
		if words == "" {
			if i != len(word)-1 {
				result.WriteString("\n")
			}
			continue
		}

		for row := 0; row < 8; row++ {
			for _, ch := range words {
				result.WriteString(banner[ch][row])
			}
			result.WriteString("\n")
		}
	}
	return result.String()
}
