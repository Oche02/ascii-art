package main

import(
	"os"
	"fmt"
)

func main(){
	if len(os.Args) != 2{
		fmt.Println("ATTENTION: expecting >> go run . [string]")
		os.Exit(1)
	}
	input := os.Args[1]
	if input == ""{
		os.Exit(0)
	}
	if input == "\n"{
		fmt.Print("\n")
		os.Exit(0)
	}
	// data, err:= os.ReadFile("banner.txt")
	// if err != nil{
	// 	fmt.Printf("ATTENTION: failed to load banner file:%v%v", input, "banner.txt", err)
	// 	os.Exit(1)
	// }

	// file := string(data)

	// err := os.WriteFile(input)
	// if err != nil{
	// 	fmt.Printf("ATTENTION: failed to write banner file:%v%v", input, "banner.txt", err)
	// 	os.Exit(1)
	// }


	fmt.Println("\nSuccessfull")
}