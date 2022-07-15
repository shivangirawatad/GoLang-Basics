package main

import "fmt"

func main() {

	arr := [4]string{"Hello", "there", "!!!"}

	fmt.Println("Elements of the array:")

	for i := 0; i < 3; i++ {
		fmt.Println(arr[i])
	}

}
