package main

import "fmt"

/*
DEFER, PANIC, AND RECOVER IN GO

1. DEFER:
   - Defers execution of a function until surrounding function returns
   - Executed in LIFO order (Last In, First Out)
   - Commonly used for cleanup (closing files, unlocking mutexes, etc.)

   Syntax: defer functionCall()

   Example:
   func main() {
       defer fmt.Println("World")
       fmt.Println("Hello")
   }
   Output: Hello
           World

   Multiple defers (LIFO):
   defer fmt.Println("1")
   defer fmt.Println("2")
   defer fmt.Println("3")
   Output: 3, 2, 1

2. PANIC:
   - Stops normal execution of current goroutine
   - Begins unwinding the stack, executing deferred functions
   - If reaches top of goroutine without recovery, program crashes
   - Similar to throwing exceptions in other languages

   Syntax: panic(value)

   Common causes:
   - Out of bounds array access
   - Type assertion failure
   - Nil pointer dereference
   - Division by zero (integers)
   - Explicit panic() call

   Example:
   func riskyOperation() {
       panic("something went wrong!")
   }

3. RECOVER:
   - Regains control after a panic
   - Only useful inside deferred functions
   - Returns the value passed to panic
   - If no panic, returns nil

   Syntax: recover()

   Example:
   defer func() {
       if r := recover(); r != nil {
           fmt.Println("Recovered from:", r)
       }
   }()

TYPICAL PATTERN:
func safeFunction() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from panic:", r)
        }
    }()

    // Code that might panic
    panic("oops!")
}

USE CASES:

1. Defer for cleanup:
   file, err := os.Open("file.txt")
   if err != nil {
       return err
   }
   defer file.Close()  // Ensures file is closed when function returns

2. Defer for unlocking:
   mu.Lock()
   defer mu.Unlock()  // Ensures mutex is unlocked

3. Panic for unrecoverable errors:
   if config == nil {
       panic("config cannot be nil")
   }

recover allows graceful degradation instead of catastrophic failure

4. Recover for graceful degradation:
   defer func() {
       if r := recover(); r != nil {
           log.Println("Recovered:", r)
           // Continue with fallback logic
       }
   }()

KEY POINTS:
- Defer executes in LIFO order
- Defer arguments are evaluated immediately, but function executes later
- Panic stops normal flow, executes defers, then crashes (unless recovered)
- Recover only works inside deferred functions
- Don't overuse panic/recover - use error returns for normal error handling
- Panic/recover is for exceptional situations, not normal control flow
*/

func main() {
	fmt.Println("=== DEFER, PANIC, AND RECOVER ===")

	// Example 1: Basic defer
	fmt.Println("\n--- Example 1: Basic Defer ---")
	deferExample()

	// Example 2: Multiple defers (LIFO)
	fmt.Println("\n--- Example 2: Multiple Defers (LIFO) ---")
	multipleDeferExample()

	// Example 3: Defer with cleanup
	fmt.Println("\n--- Example 3: Defer for Cleanup ---")
	cleanupExample()

	// Example 4: Panic without recover (commented out to prevent crash)
	// fmt.Println("\n--- Example 4: Panic (Uncomment to see crash) ---")
	// panicExample()

	// Example 5: Panic with recover
	fmt.Println("\n--- Example 5: Panic with Recover ---")
	recoverExample()

	// Example 6: Defer argument evaluation
	fmt.Println("\n--- Example 6: Defer Argument Evaluation ---")
	deferArgumentExample()

	fmt.Println("\n--- Program completed successfully ---")
}

func deferExample() {
	defer fmt.Println("Deferred: This prints last")
	fmt.Println("Normal: This prints first")
	fmt.Println("Normal: This prints second")
}

func multipleDeferExample() {
	defer fmt.Println("Defer 1")
	defer fmt.Println("Defer 2")
	defer fmt.Println("Defer 3")
	fmt.Println("Normal execution")
	// Output order: Normal execution, Defer 3, Defer 2, Defer 1
}

func cleanupExample() {
	fmt.Println("Opening resource...")
	defer fmt.Println("Closing resource (cleanup)")
	fmt.Println("Using resource...")
	fmt.Println("Done with resource")
	// Cleanup happens automatically when function returns
}

func panicExample() {
	fmt.Println("About to panic...")
	panic("Something went wrong!")
	fmt.Println("This will never print")
}

func recoverExample() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
			fmt.Println("Program continues normally")
		}
	}()

	fmt.Println("Before panic")
	panic("Intentional panic!")
	fmt.Println("This won't print")
}

func deferArgumentExample() {
	x := 10
	defer fmt.Println("Deferred x:", x) // x evaluated now (10)
	x = 20
	fmt.Println("Current x:", x) // 20
	// When defer executes, it prints 10 (not 20)
}

// Real-world example: Safe division
func safeDivide(a, b int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("division error: %v", r)
		}
	}()

	result = a / b // May panic if b == 0
	return result, nil
}
