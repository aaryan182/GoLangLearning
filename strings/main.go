package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Learning strings in GoLang")

	data := "aaryan, bajaj, rohan"
	parts := strings.Split(data, ",") // separator is comma in this case
	// data := "aaryan.bajaj.rohan"
	// parts := strings.Split(data, ".") // separator is dot in this case
	fmt.Println(parts)


	str := "Time is decreasing now now"
	count := strings.Count(str, "now")
	fmt.Println(count)

	str1 := "    Hello World    "
	str1 = strings.TrimSpace(str1)
	fmt.Println(str1)

	str3 := "aaryan"
	str4 := "bajaj"
	result := strings.Join([]string{str3,"sher", str4}, " ")
	fmt.Println(result)
}