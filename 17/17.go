package main

import "fmt"

func main() {
	tagsViews := map[string]int{
		"unix":   0,
		"python": 1,
		"go":     2,
		"golang": 3,
		"devops": 4,
		"gc":     5,
	}
	for key, views := range tagsViews {
		fmt.Println("There are", views, "views for", key)
	}
}
