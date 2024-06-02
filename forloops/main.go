package main

import "fmt"

func main() {
	fmt.Println("Hi")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	count := 0
	for {
		fmt.Println(count)
		count++
		if count == 5 {
			break
		}
	}

	numbers := []int{1, 23, 4, 5, 6}
	for index, value := range numbers {
		fmt.Println(index, value)
	}

	data := "Hello, world!"
	for index, char := range data {
		fmt.Println("Index is: ", index, "Character is : ", char)
	}

}
