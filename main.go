package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: Expecting>> go run . [string] Or go run . [string] [banner]")
		return
	}
	if len(os.Args) > 3 {
		fmt.Println("Error: Expecting>> go run . [string] Or go run . [string] [banner]")
		return
	}
	bannerfile := "standard.txt"
	if len(os.Args) == 3 {
		bannerfile = os.Args[2] + ".txt"
	}
	input := os.Args[1]
	if input == "" {
		return

	}
	if input == "\\n" {
		fmt.Println()
		return
	}

	banner, err := LoadBanner(bannerfile)
	if err != nil {
		return
	}
	output := GenerateArt(input, banner)
	if strings.HasSuffix(output, "\n") {
		output = strings.TrimSuffix(output, "\n")
		fmt.Println(output)
	}

}
