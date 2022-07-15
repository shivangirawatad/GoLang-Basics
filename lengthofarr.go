package main

import "fmt"

func main() {

	var myarr [3]string

	myarr[0] = "Hello"
	myarr[1] = "there"
	myarr[2] = "!!!"

	fmt.Println("Length of Array:", len(myarr))
}
