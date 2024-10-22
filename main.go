package main

import (
	"fmt"
	myutil "myLearning/myUtil"
	"math/rand"
)

func main() {
	fmt.Println("Learning Go lang ")
	myutil.PrintMessage("Calling from other package")
	fmt.Println(rand.Intn(10))
}