package main

import "fmt"

//Slices is a flexible and dynamic ds that provides more power than arrays
//Dynamic Size

//In slice --> cap(capacity) and size

func main() {
	fmt.Println("Hello")
	number := []int{1, 23, 4, 4}
	stringsData := []string{"hello, hi, bye"}
	fmt.Println(number)
	fmt.Println(stringsData)
	number = append(number, 2, 20, 30)
	fmt.Printf("number has a data type : %T\n", number)
}
