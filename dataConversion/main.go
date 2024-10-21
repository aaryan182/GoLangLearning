package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int = 10
	fmt.Println("Value of num is", num)
	fmt.Printf("Type of num is %T\n", num)

	var data float64 = float64(num)
	fmt.Println("Value of data is", data)
	fmt.Printf("Type of data is %T\n", data)

	num = 123
	str := strconv.Itoa(num)
	fmt.Println("Value of str is", str)
	fmt.Printf("Type of str is %T\n", str)

	number_string := "123"
	number_int, _ := strconv.Atoi(number_string)
	fmt.Println("Value of number_int is", number_int)
	fmt.Printf("Type of number_int is %T\n", number_int)

	number_string1 := "123.45"
	number_float, _ := strconv.ParseFloat(number_string1, 64)
	fmt.Println("Value of number_float is", number_float)
	fmt.Printf("Type of number_float is %T\n", number_float)
}