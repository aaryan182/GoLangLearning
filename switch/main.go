package main

import "fmt"

func main() {
	fmt.Println("Learning switch in golang")

	day := 2

	switch day {
	case 1: 
		fmt.Println("Monday")
	case 2: 
		fmt.Println("Tuesday")
	case 3: 
		fmt.Println("Wednesday")
	default:
		fmt.Println("Not a valid day")
	}
}