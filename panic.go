package main

import "fmt"

func main() {
	var myarr [3]string
	myarr[0] = "GFG"
	myarr[1] = "GeeksforGeeks"
	myarr[2] = "Geek"
	fmt.Println("Elements of Array:")
	fmt.Println("Element 1: ", myarr[0])
	// Program panics because the size of the array is
	// 3 and we try to access index 5 which is not
	// available in the current array, So, it gives an runtime error
	fmt.Println("Element 2: ", myarr[5])
}
