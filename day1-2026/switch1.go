package main

import "fmt"

func main() {
	fmt.Println("=== SWITCH STATEMENTS IN GO ===")
	month := 1
	switch month {
	case 0:
		fmt.Println("0 not allowed")
	case 1:
		fmt.Println("January")
	}
}
