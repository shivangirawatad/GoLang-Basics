package main

import "fmt"

func fun(a [6]int, size int) int {
	var k, val, r int

	for k = 0; k < size; k++ {
		val += a[k]
	}

	r = val / size
	return r
}

func main() {

	var arr = [6]int{17, 60, 29, 35, 4, 34}
	var res int

	res = fun(arr, 6)
	fmt.Printf("Final result is: %d ", res)
}
