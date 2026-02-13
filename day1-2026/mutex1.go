package main

import (
	"fmt"
	"sync"
)

/*
MUTEX - Prevents race conditions when multiple goroutines access shared data

Problem: Without mutex, goroutines can corrupt data
	counter = 0
	Goroutine 1: reads 0 → adds 1 → writes 1
	Goroutine 2: reads 0 → adds 1 → writes 1
	Result: 1 (WRONG! Should be 2)

Solution: Use mutex to lock access
	mu.Lock()      // Only one goroutine enters
	counter++      // Safe modification
	mu.Unlock()    // Release for others

KEY CONCEPTS

- Mutex ensures only one goroutine accesses a block of code at a time.
- WaitGroup waits for multiple goroutines to finish.
- Always call Add() before launching the goroutine.
- Always call Done() in the goroutine (use defer).

Best Practice:
	mu.Lock()
	defer mu.Unlock()  // Always use defer!
	// protected code here
*/

// Example 1: UNSAFE Counter (Race Condition)
func unsafeCounter() {
	counter := 0
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter++ // ❌ UNSAFE! Multiple goroutines access simultaneously
		}()
	}

	wg.Wait()
	fmt.Println("Unsafe Counter:", counter, "(may be wrong!)")
}

// Example 2: SAFE Counter with Mutex (using defer)
func safeCounter() {
	var mu sync.Mutex
	counter := 0
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock() // ✅ Always use defer for safety
			counter++
		}()
	}

	wg.Wait()
	fmt.Println("Safe Counter:", counter, "(always correct!)")
}

/*
KEY POINTS:
- Mutex ensures only ONE goroutine accesses shared data at a time
- Always use defer mu.Unlock() after mu.Lock()
- WaitGroup waits for all goroutines to finish
- Call wg.Add(1) before launching goroutine
- Call wg.Done() inside goroutine (use defer)

Common Mistakes:
- Forgetting mu.Unlock() → deadlock
- Always use defer mu.Unlock()
*/

func main() {
	fmt.Println("=== MUTEX EXAMPLES ===\n")

	unsafeCounter()
	fmt.Println()
	safeCounter()

	fmt.Println("\n✅ Mutex protects shared data from race conditions")
	fmt.Println("Run with: go run -race mutex1.go (to detect races)")
}
