package main

import "fmt"

func sum(s []int) int {
	total := 0
	for i := range len(s) {
		total += s[i]
	}
	return total
}

func factorial(fact int) int {
	if fact == 0 {
		return 1
	}
	for i := fact; i > 1; i-- {
		fact *= i
	}
	return fact
}
func main() {
	var a []int
	fmt.Println(sum([]int{1, 2, 3, 4, 5}))
	for i := 0; i < 10; i++ {
		a = append(a, i*i)
	}
	fmt.Printf("Sum of array a int is %d \n", sum(a))
	fmt.Printf("Factorial of given num %d is : %d \n", 5, factorial(5))
}
