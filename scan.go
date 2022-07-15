package main

import "fmt"

func main() {
	var fname string
	fmt.Println("Enter fname: ")
	fmt.Scan(&fname)
	var lname string
	fmt.Println("Enter lname: ")
	fmt.Scan(&lname)
	fmt.Println(fname + lname)
}
