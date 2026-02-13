package main

import "fmt"

/*
RECURSION NOTES:
1. Recursion: Function calling itself
2. Must have base case (stopping condition) to avoid infinite loop
3. Each recursive call works on smaller problem
4. Example: factorial(5) = 5 * factorial(4) = 5 * 4 * 3 * 2 * 1

FUNCTION ORDERING IN GO:
5. Function declaration order does NOT matter in Go
6. Can call functions before or after they are declared
7. Go compiles in two passes: collects declarations first, then compiles bodies
8. Convention: main() first, then helper functions (but not required)
*/

func main() {
	fmt.Println("Factorial of 5 is ", factorial(5))
}
func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}
