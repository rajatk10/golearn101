package main

import "fmt"

type Counter struct {
	count int
}

func (c Counter) IncrementByValue() {
	c.count++
	fmt.Printf("Inside IncrementByValue: %d\n", c.count)
}

func (c *Counter) IncrementByPointer() {
	c.count++
	fmt.Printf("Inside IncrementByPointer: %d\n", c.count)
}

func main() {
	fmt.Println("=== Value Receiver Demo ===")
	counter1 := Counter{count: 0}
	fmt.Printf("Before: %d\n", counter1.count)
	counter1.IncrementByValue()
	fmt.Printf("After: %d (unchanged!)\n\n", counter1.count)

	fmt.Println("=== Pointer Receiver Demo ===")
	counter2 := Counter{count: 0}
	fmt.Printf("Before: %d\n", counter2.count)
	counter2.IncrementByPointer()
	fmt.Printf("After: %d (changed!)\n", counter2.count)
}
