package main

import "fmt"

func main() {
	fmt.Println("Learning slices in golang")
	numbers := []int{1,2,3,4,5}
	fmt.Println("numbers are:",numbers)
	fmt.Printf("Numbers type is:%T\n", numbers)
	fmt.Println("length of numbers is:", len(numbers))
	numbers = append(numbers, 6,7,8,9)
	fmt.Println("new number list is:", numbers)
	fmt.Println("length of numbers is:", len(numbers))

	name := []string{}
	fmt.Println("name slice is:",name)

	// slice making 
	numbers1 := make([]int, 3 , 5) // make(type, length, capacity)
	fmt.Println("Slice is:",numbers1)
	fmt.Println("Length is:",len(numbers1))
	fmt.Println("capacity is:",cap(numbers1))
	numbers1 = append(numbers1, 6)
	numbers1 = append(numbers1,7)
	fmt.Println("Slice is:",numbers1)
	fmt.Println("Length is:",len(numbers1))
	fmt.Println("capacity is:",cap(numbers1))

	// when capacity is now same as length , now we append capacity increases or doubles as it is slice ( same like vector in array)
	numbers1 = append(numbers1,8)
	fmt.Println("Slice is:",numbers1)
	fmt.Println("Length is:",len(numbers1))
	fmt.Println("capacity is:",cap(numbers1))


	// using make function only we can initialise slice 



}

