package main

import "fmt"

func main() {
	fmt.Println("Learning map in golang")

	studentGrades := make(map[string]int)
	studentGrades["aaryan"] = 90
	studentGrades["bajaj"] = 95
	studentGrades["rahul"] = 80
	studentGrades["rohan"] = 44

	fmt.Println("student grade of aaryan is", studentGrades["aaryan"]) // 90

	// modify map 
	studentGrades["bajaj"] = 88
	fmt.Println("student grade of bajaj is", studentGrades["bajaj"]) // 88

	// to delete map 
	delete(studentGrades, "rohan")
	fmt.Println("student grade of rohan is", studentGrades["rohan"]) // 0

	// check if key exists 
	_, exists := studentGrades["rohan"]
	fmt.Println(exists)

	// unordered way 
	for key, value := range studentGrades {
		fmt.Printf("The grade of %s is %d\n", key, value)
	}

	person := map[string]int {
		"aaryan": 90,
		"bajaj": 95,
		"rahul": 80,
		"rohan": 44,
	}

	fmt.Println(person)
}
