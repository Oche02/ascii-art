package main

func RenderLine(text string, banner map[rune][]string) []string {
	final := make([]string, 8)

	if text == "" {
		return final
	}
	for _, ch := range text {
		for row := 0; row < 8; row++ {
			final[row] += banner[ch][row]
		}
	}
	return final
}
