package main

import "fmt"

func main() {
	a := 3
	b := 5
	N := 50
	var sum int
	fmt.Printf("a = : %v b = %v \n", a, b)
	for i := 0; i < N; i++ {
		if i%a == 0 {
			sum = sum + i
		} else if i%b == 0 {
			sum = sum + i
		}
	}
	fmt.Println("Sum = %v", sum)
}
