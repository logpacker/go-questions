package main

import "fmt"

func main() {
	for i := 0; i < 4; i++ {
		defer fmt.Print(i)
	}
}
