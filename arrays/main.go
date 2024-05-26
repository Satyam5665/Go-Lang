package main

import "fmt"

//Arrays are intialized to default value
func main() {
	fmt.Println("Arrays in Go Lang")
	var name [5]string
	name[0] = "abc"
	name[4] = "def"
	fmt.Println(name)

	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers)

	var price [5]bool
	fmt.Println(price)

	fmt.Println(len(numbers))

}
