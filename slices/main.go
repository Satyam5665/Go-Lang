package main

import "fmt"

//Slices is a flexible and dynamic ds that provides more power than arrays
//Dynamic Size

//In slice --> cap(capacity) and size
//Capacity doubles by 2 if there is requirement

//use make function to create empty slice

func main() {
	fmt.Println("Hello")
	number := []int{1, 23, 4, 4}
	stringsData := []string{"hello, hi, bye"}
	fmt.Println(number)
	fmt.Println(stringsData)

	stock := make([]int, 0)
	fmt.Println(stock)
	number = append(number, 2, 20, 30)
	fmt.Printf("number has a data type : %T\n", number)

}
