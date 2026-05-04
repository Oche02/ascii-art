package main

import(
	"os"
	"fmt"
	"time"
)

func main(){
	for i := 0; i <= 100; i++ {
		fmt.Printf("\rLoading...%d%%", i)
		time.Sleep(10 * time.Millisecond)
	}
	if len(os.Args) != 2{
		fmt.Println("ATTENTION: expecting >> go run . [string]")
		os.Exit(1)
	}

	// input := os.Args[1]

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