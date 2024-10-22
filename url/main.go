package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("Learning url handling in Go lang")
	myURL := "https://example.com:8080/path/to/resource?key1=value1&key2=value2#section1"
	// fmt.Printf("myURL is %T\n", myURL)

	parseURL, err := url.Parse(myURL)
	if err != nil {
		fmt.Println("error in url parsing ",err)
		return
	}
	// fmt.Printf("parseURL is %T\n", parseURL)

	fmt.Println("Scheme is ", parseURL.Scheme)
	fmt.Println("Host is ", parseURL.Host)
	fmt.Println("Path is ", parseURL.Path)
	fmt.Println("RawQuery is ", parseURL.RawQuery)

	// modifying url components
	parseURL.Path = "/newpath"
	parseURL.RawQuery = "key3=value3"
	fmt.Println("Modified URL is ", parseURL.String())
}