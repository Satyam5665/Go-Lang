package main

import "fmt"

func divide(a, b float32) float32 {
	return a / b
}
func main() {
	fmt.Println("Error Handling in the Function")
	ans := divide(10, 4)
	fmt.Println("Division of two numbers : ", ans)
}
