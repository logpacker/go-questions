package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	// 1
	s := ""
	for i := 0; i < 5; i++ {
		s += "a"
	}
	fmt.Println(s)

	// 2
	var buffer bytes.Buffer
	for i := 0; i < 5; i++ {
		buffer.WriteString("a")
	}
	fmt.Println(buffer.String())

	// 3
	sl := []string{"a", "a", "a", "a", "a"}
	fmt.Printf(strings.Join(sl, ""))
}
