package main

import "fmt"

func main() {

	fmt.Println("Hello")
	day := 1

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	default:
		fmt.Println("Unknow Day")
	}

}
