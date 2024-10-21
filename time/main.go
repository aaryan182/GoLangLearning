package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println(currentTime)

	formattedTime := currentTime.Format("2006-01-02 15:04:05 Monday")
	fmt.Println(formattedTime)

	layoutStr := "2006-01-02 15:04:05"
	DateStr := "2022-12-31 23:59:59"
	parsedTime, _ := time.Parse(layoutStr, DateStr)
	fmt.Println(parsedTime)
}