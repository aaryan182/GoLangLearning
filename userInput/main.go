package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("What is your name ?")
	// one way 

	// var name string
	// fmt.Scan(&name);
	// fmt.Println("Hello ", name)

	// another way 
	reader := bufio.NewReader(os.Stdin)
	name , _ := reader.ReadString('\n')
	fmt.Println("Hello", name)
}