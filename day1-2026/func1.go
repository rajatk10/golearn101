package main

import (
	"errors"
	"fmt"
)

/*
FUNCTION NOTES:
1. Function syntax: func name(param type) returnType { body }
2. Parameter types are MANDATORY - func greet(name) won't compile
3. Return type is MANDATORY if function returns something
4. No return type needed for void functions
5. Multiple parameters: func add(a, b int) int - shorthand for same type
6. Multiple returns: func divide(a, b float64) (float64, error)
7. Named returns: func getUser() (name string, age int) { return }
8. Variadic: func sum(nums ...int) int - accepts variable number of args
*/

func greet(name string) string {
	return "Hello , " + name
}
func sayHello() {
	fmt.Println("Hello World function without any return type")
}
func sayHello1(name string) {
	fmt.Println("Hello World Function without return type", name)
}

// Variadic functions - functions with multiple or unknown no. of inputs.
func sum1(num ...int) int {
	total := 0
	for _, num := range num {
		total += num
	}
	return total
}

// function with multiple return values.
func multipleReuturnValues(a, b int) (int, error) {
	if b != 0 {
		//Here error is return type while if no error set it as nil
		return a / b, nil
	} else {
		return 0, fmt.Errorf("Zero Division Error")
	}
	return 0, errors.New("Cannot divide by zero")
}
func main() {
	fmt.Println(greet("Rajat"))
	sayHello()
	sayHello1("Rajat")
	fmt.Println("Sum from variadic function is ", sum1(1, 2, 3, 4, 5))
	res, err := multipleReuturnValues(10, 2)
	if err != nil {
		fmt.Println("Error is Multiple Return Value function", err)
	} else {
		fmt.Println("Result is Multiple Return Value function", res)
	}
}
