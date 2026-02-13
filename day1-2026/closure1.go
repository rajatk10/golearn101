package main

import "fmt"

/*
CLOSURES NOTES:
1. Closure = Anonymous function + Captured variables from outer scope
2. Anonymous function: func() with NO NAME after func keyword
3. Assigned to variable: increment := func() { }
4. Can access and modify variables from surrounding scope
5. The function "remembers" outer variables even after outer function ends
6. Not all anonymous functions are closures (only if they capture variables)

NAMED vs ANONYMOUS FUNCTION:
Named function:
    func greet(name string) {
        fmt.Println("Hello", name)
    }

Anonymous function (no name after func):
    func(name string) {
        fmt.Println("Hello", name)
    }

Key: Anonymous = NO NAME after 'func' keyword
     Parameters and body work exactly the same

ANONYMOUS FUNCTION vs CLOSURE:
- Anonymous: add := func(a, b int) int { return a + b }  (no outer vars)
- Closure: increment := func() { s++ }  (captures 's' from outer scope)

CLOSURE PATTERNS:
1. Direct assignment:
   increment := func() { s++ }

2. Returned from function:
   func makeCounter() func() int {
       count := 0
       return func() int { count++; return count }  // Closure
   }
   counter := makeCounter()  // Assignment happens here

COMMON USE CASES:
- HTTP handlers with middleware
- Goroutines (capturing loop variables)
- Defer with cleanup
- Callbacks with context
*/

func main() {
	s := 10
	//Closures in go is used to access variables from outer function, i.e. surrounding function
	increment := func() int {
		s++
		return s
	}
	//The anonymous function "remembers" s even though s is defined outside the function.
	fmt.Println(increment())
	fmt.Println(increment())
}
