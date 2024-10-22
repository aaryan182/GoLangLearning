package main

import "fmt"

func main() {
	fmt.Println("Learning defer in golang")
	fmt.Println("Starting")
	defer fmt.Println("Processing") // defer moves to the end of the execution of program 
	defer fmt.Println("Ending")

	// when there are two or more defer , then it works like stack (LIFO)
}