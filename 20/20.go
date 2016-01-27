package main

import "os"

func main() {
	src, err := os.Open("filename")
	// 1
	//defer src.Close()
	if err != nil {
		return
	}

	// 2
	//defer src.Close()

	src.WriteString("Hello")

	// 3
	//defer src.Close()
}
