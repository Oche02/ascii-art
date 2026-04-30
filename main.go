package main

import(
	"os"
	"fmt"
)

func main(){
	if len(os.Args) != 2{
		fmt.Println("ATTENTION: expecting >> go run . [string]")
		os.Exit(0)
	}

	// input: os.Args[1]

	// data, err:= os.Reader(input)
	// if err != nil{
	// 	fmt.Println("ATTENTION: failed to load banner file:%v")
	// }

	fmt.Println("Successfull")
}