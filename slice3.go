package main

import "fmt"

func main() {

	myslice := []string{"Hi", "everyone", "today", "we",
		"will", "learn", "GoLang"}

	for index, ele := range myslice {
		fmt.Printf("Index = %d and element = %s\n", index+3, ele)
	}
}
