package main

import "fmt"

func main() {
	var i interface{} = "hello"

	fmt.Println("=== Basic Type Assertion ===")
	s := i.(string)
	fmt.Printf("String value: %s\n", s)

	fmt.Println("\n=== Safe Type Assertion ===")
	s, ok := i.(string)
	if ok {
		fmt.Printf("Successfully got string: %s\n", s)
	}

	n, ok := i.(int)
	if !ok {
		fmt.Printf("Failed to convert to int, got: %v\n", n)
	}

	fmt.Println("\n=== Type Switch ===")
	describe(42)
	describe("hello")
	describe(3.14)
	describe(true)
	describe([]int{1, 2, 3})
}

func describe(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Integer: %d\n", v)
	case string:
		fmt.Printf("String: %s (length: %d)\n", v, len(v))
	case float64:
		fmt.Printf("Float: %.2f\n", v)
	case bool:
		fmt.Printf("Boolean: %t\n", v)
	default:
		fmt.Printf("Unknown type: %T\n", v)
	}
}
