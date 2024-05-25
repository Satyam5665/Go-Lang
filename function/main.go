package main

import "fmt"

func simpleFunction() {
	fmt.Println("Hello World")
}

func add(a int, b int) int {
	return a + b
}

//Main function is the executable Function
func main() {
	fmt.Println("we are here!")
	simpleFunction()
	res := add(1, 2)
	fmt.Println(res)
}
