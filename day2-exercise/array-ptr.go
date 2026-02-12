package main

import "fmt"

func updateArray(arr *[3]int) {
	(*arr)[0] = 100
	fmt.Printf("Updated array %v\n", *arr)
}
func main() {
	arr := [3]int{1, 2, 3}
	updateArray(&arr)
}
