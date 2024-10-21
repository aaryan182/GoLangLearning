package main

import "fmt"

func main() {
	fmt.Println("Learing pointer in Golang")

	// one way

	// var num int 
	// num = 2

	// var ptr *int 
	// ptr = &num

	// another way 
	num := 2
	ptr := &num

	fmt.Println("Value of num is", num)
	fmt.Println("Value of ptr is", ptr)
	fmt.Println("Value of *ptr is", *ptr)
}

