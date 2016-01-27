package main

import "fmt"
import "unsafe"

func main() {
	s := struct {
		A float32
		B string
	}{0, "go"}

	fmt.Printf("%T, %d\n", s, unsafe.Sizeof(s))
}
