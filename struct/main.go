package main

import "fmt"

type person struct {
	name string
	age int
	height int
}

type Contact struct {
	email string
	phone int
}

type Address struct {
	city string
}

type Employee struct{
	Person_Details person
	Contact_Details Contact
	Address_Details Address
}

func main() {
	fmt.Println("Learning structure in goLang")

	var aaryan person
	fmt.Println(aaryan)

	// one way to assign values
	aaryan.name = "aaryan"
	aaryan.age = 23
	aaryan.height = 6
	fmt.Println(aaryan)

	// another way to assign values 
	aaryan1 := person{"aaryan", 23, 6}
	fmt.Println(aaryan1)


	employee1 := Employee{
		Person_Details: person{"aaryan", 23, 6},
		Contact_Details: Contact{"aaryandotcom", 123456789},
		Address_Details: Address{"aaryan"},
	}
	fmt.Println(employee1)
}