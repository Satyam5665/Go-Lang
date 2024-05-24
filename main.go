package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hi")

	var name string = "Satyam"
	var version = 1.2
	version = 4
	const age = 4
	person := "Satyam Singh"
	fmt.Println((person))
	fmt.Println((version))
	fmt.Println(name)
	fmt.Println((age))

	var Public = "data is important"
	var privatevariable = "data will be used only here"
	fmt.Println((privatevariable))
	fmt.Println((Public))
}

//If a variable is initialized in Capital Letter then it can be imported in other foles

//Pointers:
//Packages are tightly coupled to a Directory.
//In one directory, only one package can be present, if you want to create more than one, than simply create anothe
// Folder.
