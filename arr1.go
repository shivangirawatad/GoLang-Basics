package main

import "fmt"

func main() {

	var myarr [3]string

	myarr[0] = "Hello"
	myarr[1] = "there"
	myarr[2] = "!!!"

	fmt.Println("Elements of Array:")
	fmt.Println("Element 1: ", myarr[0])
	fmt.Println("Element 2: ", myarr[1])
	fmt.Println("Element 3: ", myarr[2])
}
