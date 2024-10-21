package myutil

import "fmt"

func PrintMessage(message string) {
	fmt.Println(message) 
	var name string = "aaryan"
	var version = "version 1" // string is not compulsory to write , at compile time it automatically considered as string
	fmt.Println(name)
	fmt.Println(version)

	var money int = 1000
	var money1 = 1000
	fmt.Println(money);
	fmt.Println(money1);

	var dimension float64 = 3.14
	fmt.Println(dimension)

	var decided bool = true
	fmt.Println(decided)

	const pi = 67.2
	// pi = 3.14 error cannot change value 
	fmt.Println(pi)



	// shortcut to write variables 
	person := "223"
	fmt.Println(person) 

	// the variables whose first letter is capital we can export those variables 
	Public := "capital variable"
	fmt.Println(Public)
	// same this applicable for function name as well 

}