package main

import "strings"

func SplitInput(input string) []string {
	word := strings.Split(input, "\\n")
	return word
}
