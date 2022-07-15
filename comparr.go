package main

import "fmt"

func main() {

	arr1 := [3]int{9, 7, 6}
	arr2 := [...]int{9, 7, 6}
	arr3 := [3]int{9, 5, 3}

	fmt.Println(arr1 == arr2)
	fmt.Println(arr2 == arr3)
	fmt.Println(arr1 == arr3)

}
