package main

import "fmt"

func main() {
	fmt.Println("pointers in go ")

	myNumber := 23

	var ptr = &myNumber

	fmt.Println("value of actual pointer is ", ptr)
	fmt.Println("value of actual pointer is ", *ptr)

	*ptr = *ptr + 5
	fmt.Println("new value is ", *ptr)
	fmt.Println("new value is ", &ptr)
}
