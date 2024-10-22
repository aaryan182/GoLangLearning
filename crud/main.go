package main

import (
	"fmt"
	// "io/ioutil"
	"encoding/json"
	"net/http"
)

type Todo struct {
	UserID int `json:"userId"`
	Id int `json:"id"`
	Title string `json:"title"`
	Completed bool `json:"completed"`
}

func main() {
	fmt.Println("Learning crud in Golang")
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error getting GET response", err)
		return
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK{
		fmt.Println("Error status code",res.StatusCode)
		return
	}

	// data,err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println("error reading response body",err)
	// 	return
	// }
	// fmt.Println("response: ",string(data))

	var todo Todo
	json.NewDecoder(res.Body).Decode(&todo)
	if err != nil {
		fmt.Println("error reading response body", err)
		return
	}
	fmt.Println("response: ", todo)
}