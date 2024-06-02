package main

import "fmt"

func main() {
	fmt.Println("Hello")
	x := 10
	if x == 0 {
		fmt.Println("x is zero")
	} else if x > 5 {
		fmt.Println("dONE")
	} else {
		fmt.Println("bye!")
	}

	y := 10
	if x > 5 && y > 5 {
		fmt.Println("Hello")
	} else {
		fmt.Println("Bye")
	}
}
