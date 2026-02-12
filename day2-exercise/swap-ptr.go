package main

import "fmt"

func swapNumbers(a, b int) {
	var pa *int = &a
	var pb *int = &b
	fmt.Printf("Values prior to swap a : %d , b : %d \n", a, b)
	*pa = *pb
	*pb = *pa

	fmt.Printf("Values post swap a : %d , b : %d \n", a, b)
}

func printAddressPointer(a int) {
	var p1 *int = &a
	fmt.Printf("Address of variable a : %d is %v \n", a, p1)
}
func main() {
	swapNumbers(5, 10)
	printAddressPointer(500)
}
