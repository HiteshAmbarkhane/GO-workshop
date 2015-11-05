package main

import "fmt"
import "strconv"
import "os"

func main() {
	i, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("couldn't convert number: %v\n", err)
		return
	}
	fmt.Printf("Converted integer: %v \n", i)
	if i%2 == 1 {
		fmt.Println("Odd number")
	} else {
		fmt.Println("Even number")
	}
}
