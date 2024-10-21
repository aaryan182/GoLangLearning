package main

import "fmt"

func divide(a, b float64) (float64,error) {
	if(b == 0) {
		return 0, fmt.Errorf("denominator must not be 0") 
	}
	return a/b, nil
}

func divide1(a, b float64) (float64,string) {
	if(b == 0) {
		return 0, "denominator must not be 0" 
	}
	return a/b, "nil"
}

func main() {
	fmt.Println("error handling learing in go ")
	ans, _ := divide(10,4)
	fmt.Println(ans)
	ans1, _ := divide1(10,3)
	fmt.Println(ans1)
}