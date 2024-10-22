package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello from goroutine")
	time.Sleep(2000 * time.Millisecond) // sleep for 100 milliseconds
	fmt.Println("Hello again from goroutine")
}

func sayHi(){
	fmt.Println("Hi from goroutine")
}

func main(){
	fmt.Println("learning Goroutines in golang")
	go sayHello()
	go sayHi()
}