package main

import "fmt"

func main() {
	age := 12
	name := "aaryan"
	height := 6.2

	fmt.Println("age: ", age, "name: ", name, "height: ", height)
	// fmt.Printf("%v %v %v", age, name, height)
	fmt.Printf("age is %d\n", age);
	fmt.Printf("height is %.2f\n", height); // 6.20
	fmt.Printf("height is %.3f\n", height); // 6.200
	fmt.Printf("Type of name is %T\n", name); 

	fmt.Printf("Name: %s, Age: %d, Height: %.2f\n", name, age, height)
}