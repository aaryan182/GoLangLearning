package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Learning web request in Golang")
	res,err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error getting GET response",err)
		return 
	}
	defer res.Body.Close()
	fmt.Printf("Type of response is %T\n",res)

	// read the response body 
	data , err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("error reading response body",err)
		return 
	}

	fmt.Println("response: ",string(data));
}