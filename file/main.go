package main

import (
	"fmt"
	// "io/ioutil"
	"os"
)

func main() {
	fmt.Println("Learning file management in golang")

	// create file

	// file,err := os.Create("test.txt")
	// if err != nil {
	// 	fmt.Println("Error in creating file")
	// }
	// defer file.Close()
	// fmt.Println("File created successfully")
	// byte, err := io.WriteString(file, "Hello content is written file main.go \n")
	// fmt.Println("Number of bytes written", byte)	

	// fmt.Println("File written successfully")
	// if err != nil {
	// 	fmt.Println("Error in writing file")
	// }

	// opening file 
	file,err := os.Open("test.txt")
	if err != nil {
		fmt.Println("Error in opening file")
	}
	defer file.Close()

	// one way to read file

	// byte, err := io.ReadAll(file)
	// fmt.Println("Number of bytes read", byte)	
	// fmt.Println("File read successfully")
	// if err != nil {
	// 	fmt.Println("Error in reading file")
	// }

	// another way to read file

	// buffer := make([]byte , 1024);
	// for{
	// 	n,err := file.Read(buffer)
	// 	if err== io.EOF{
	// 		break
	// 	}
	// 	if err != nil{
	// 		fmt.Println("error while reading file",err)
	// 	}
	// 	// process the read content
	// 	fmt.Println(string(buffer[:n]))
	// }


	// another way to read file

	// content , err := ioutil.ReadFile("test.txt")
	content , err := os.ReadFile("test.txt")
	if err != nil {
		fmt.Println("error while reading file", err)
		return 
	}
	fmt.Println(string(content))
}