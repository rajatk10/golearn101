package main

import "fmt"

func isOdd(a int) bool {
	if a%2 == 0 {
		return false
	}
	return true
}

func main() {
	fmt.Printf("Is %d odd ? : %t \n", 10, isOdd(10))
}
