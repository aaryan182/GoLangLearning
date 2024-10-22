package main

import (
	"fmt"
	"sync"
)

func worker(i  int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("worker %d started\n",i)
	// some task is happening
	fmt.Printf("worker %d end\n",i)
}	


func main() {
	fmt.Println("Learning sync waitGroup in golang")

	// create waitGroup
	var wg sync.WaitGroup

	// start 3 worker go routines
	for i:= 1; i<=5; i++ {
		wg.Add(1) // increment the waitGroup counter
		go worker(i , &wg)
	}

	// wait for the workers to finish
	wg.Wait()
	fmt.Println("main ended")
}