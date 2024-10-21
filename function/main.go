package main

import "fmt"

func simpleFunction (){
	fmt.Println("this is simple function")
}

func add(a int , b int) int{
	return a + b
}

func add1(a , b , c , d int) int{
	return a + b + c + d
}

func multiply(a, b int) (result int){
	result = a * b
	return
}

func main() {
	fmt.Println("learning golang")
	simpleFunction()
	sum := add(10, 20)
	fmt.Println(sum)
	sum1 := add1(10, 20, 30, 40)
	fmt.Println(sum1)
	pro := multiply(10,20)
	fmt.Println(pro)
}