package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		fmt.Println("1")
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		fmt.Println("2")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("3")
}
