package main

import "fmt"

func main() {
	// Declaration of variables
	helloMessage := "Hello"
	worldMessage := "world"

	// Println
	fmt.Println(helloMessage, worldMessage)
	fmt.Println(helloMessage, worldMessage)

	// Printf
	name := "Platzi"
	courses := 500
	fmt.Printf("%s has more than %d courses\n", name, courses)
	fmt.Printf("%v has more than %v courses\n", name, courses)

	// Sprintf
	message := fmt.Sprintf("%v has more than %v courses", name, courses)
	fmt.Println(message)

	// Data type
	fmt.Printf("helloMessage : %T", helloMessage)
	fmt.Printf("courses : %T", courses)
}
