package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age int `json:"age"`
	IsAdult bool `json:"is_adult"`
}

func main() {
	fmt.Println("Learning json in golang")
	person1 := Person{Name: "aaryan", Age: 23, IsAdult: true}
	fmt.Println(person1)

	// convert person into JSON Encoding (Marshalling)
	jsonData, err :=json.Marshal(person1)
	if err != nil {
		fmt.Println("error in marshalling", err)
	}

	fmt.Println("Person1 data is ", string(jsonData))

	// decoding (unmarshalling)

	var personData Person
	err = json.Unmarshal(jsonData, &personData)
	if err != nil {
		fmt.Println("error in unmarshalling", err)
	}
	fmt.Println("Person data is after unmarshalling ", personData)
}