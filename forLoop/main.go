package main

import "fmt"

func main() {
	fmt.Println("Learning for loop in Golang")

	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// counter :=0
	// for counter < 10 {
	// 	fmt.Println(counter)
	// 	counter++
	// }

	// range keyword
	numbers := []int{11, 21, 31, 41, 51}
	for index,value := range numbers {
		fmt.Printf("Index %d, Value %d\n" , index, value)
	}

	data := []string{"aaryan", "bajaj"}
	for index,char := range data {
		fmt.Printf("Index %d, Value %s\n" , index, char)
	}

	data1 := "hello aaryan"
	for index,char := range data1 {
		fmt.Printf("Index %d, Value %c\n" , index, char)
	}
}
