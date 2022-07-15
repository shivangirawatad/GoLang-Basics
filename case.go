package main

import "fmt"

func main() {
	value := 88
	switch {
	case value > 90:
		fmt.Println("A")
	case value > 80:
		fmt.Println("B")
	case value > 70:
		fmt.Println("C")
	case value > 60:
		fmt.Println("D")
	case value > 50:
		fmt.Println("E")
	default:
		fmt.Println("F")
	}

}
