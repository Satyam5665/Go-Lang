package main

import "fmt"

func divide(a, b float32) (float32, error) {
	if b == 0 {
		return 0, fmt.Errorf("denominator can't be zero")
	}
	return a / b, nil
}

//We can also use _ , it is a read only variable type
func main() {
	fmt.Println("Error Handling in the Function")
	ans, _ := divide(10, 0)
	fmt.Println("Division of two numbers : ", ans)
}
