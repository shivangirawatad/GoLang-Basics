package main

import "fmt"

func main() {
	var numbers = make([]int, 3, 5)
	printSlice(numbers)
	numbers = append(numbers, 2, 3, 4)
	printSlice(numbers)
	/* print the sub slice starting from index 1(included) to index 4(excluded)*/
	printSlice(numbers[1:4])
	/* missing lower bound implies 0*/
	printSlice(numbers[:3])
	/* missing upper bound implies len(s)*/
	printSlice(numbers[4:])
	/* create a slice numbers1 with double the capacity of earlier slice*/
	numbers1 := make([]int, len(numbers), (cap(numbers))*2)
	/* copy content of numbers to numbers1 */
	copy(numbers1, numbers)
	printSlice(numbers1)
}
func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
