package main

import "fmt"

func main() {
	fmt.Println("Learning array in Golang")

	var name[5] string
	name[0] = "aaryan"
	name[2] = "bajaj"
	
	fmt.Println("Names in the array are:" , name)

	var numbers = [5]int {1,2,3,4,5}
	fmt.Println(numbers)
	fmt.Println(len(numbers))
	fmt.Println(numbers[2])

	var price[5]int 
	fmt.Println(price)  //[0 0 0 0 0]
	fmt.Printf("Price with quoted string: %q\n" , price)
}
