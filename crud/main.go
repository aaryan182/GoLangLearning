package main

import (
	"fmt"
	"io/ioutil"
	"strings"

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

func performingGetMethod(){
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

func performingPostMethod(){
	todo := Todo{
		UserID: 22,
		Title: "Aaryan Bajaj",
		Completed: false,
	}

	// convert the todo struct into JSON
	jsonData,err := json.Marshal(todo)
	if err != nil {
		fmt.Println("error in marshalling", err)
		return
	}

	// convert json data into string
	jsonString := string(jsonData)

	// convert string data into reader
	jsonReader := strings.NewReader(jsonString)

	myURL := "https://jsonplaceholder.typicode.com/todos"

	// send post request
	res, err := http.Post(myURL, "application/json", jsonReader)

	if err != nil {
		fmt.Println("Error getting POST response", err)
		return
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusCreated{
		fmt.Println("Error status code",res.StatusCode)
		return
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("error reading response body",err)
		return
	}
	fmt.Println("response: ",string(data))

}

func main() {
	fmt.Println("Learning crud in Golang")
	// performingGetMethod()
	performingPostMethod()
}