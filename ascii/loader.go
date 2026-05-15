package ascii

import (
	"fmt"
	"os"
	"strings"
)

func LoadBanner(filePath string)(map[rune][]string, error){
	data, err := os.ReadFile(filePath)
	if err != nil{
		fmt.Fprintf(os.Stderr, "failed to read banner file %v\n", err)
		os.Exit(1)
	}
}