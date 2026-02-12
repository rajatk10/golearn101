package main

import (
	"fmt"
	"math"
)

/*
GO OPERATORS QUIZ - Questions & Answers

WHAT Questions:
================

Q1: What are arithmetic operators, and how do they differ from comparison operators?
A: Arithmetic Operators:

  - Perform mathematical calculations

  - Return numeric result

  - Operators: +, -, *, /, %

  - Modify/create values
    Example: 10 + 3 = 13

    Comparison Operators:

  - Compare two values

  - Return boolean result (true/false)

  - Operators: ==, !=, <, >, <=, >=

  - Don't modify values
    Example: 10 > 3 = true

    Key Difference:
    ┌──────────────┬─────────────┬────────────┐
    │ Feature      │ Arithmetic  │ Comparison │
    ├──────────────┼─────────────┼────────────┤
    │ Purpose      │ Calculate   │ Compare    │
    │ Return type  │ Number      │ Boolean    │
    │ Modifies     │ Yes         │ No         │
    └──────────────┴─────────────┴────────────┘

Q2: What is the purpose of the modulus operator, and where can it be practically applied?
A: Modulus (%) returns the remainder after division.

	Examples:
	10 % 3 = 1  (10 ÷ 3 = 3 remainder 1)
	15 % 4 = 3
	20 % 5 = 0

	Practical Applications:
	1. Check even/odd:
	   if num % 2 == 0 {  }

	2. Cycle through array:
	   arr[i % len(arr)]  // Wraps around

	3. Time calculations:
	   hours := 25 % 24  // 1 hour next day

	4. Distribute tasks:
	   workerID := taskID % numWorkers

	5. Check divisibility:
	   if year % 4 == 0 { }

	Complete leap year check:
	   isLeap := (year % 400 == 0) || (year % 4 == 0 && year % 100 != 0)

Q3: What is short-circuit evaluation in logical operators?
A: Go stops evaluating logical expressions as soon as the result is determined.

	For && (AND):
	false && anything  // Stops at false

	Example:
	if x != 0 && 10/x > 2 {
	    // If x == 0, stops before division (no panic!)
	}

	For || (OR):
	true || anything  // Stops at true

	Example:
	if user != nil || user.Name == "Admin" {
	    // If user != nil is true, user.Name never accessed
	}

	Benefits:
	- Performance: Skips unnecessary evaluations
	- Safety: Prevents nil pointer errors
	- Efficiency: Avoids expensive operations

Q4: What does it mean when we say operators are "type compatible" in Go?
A: Operators require operands to be of the same type or compatible types.

	❌ Not compatible:
	var a int = 10
	var b float64 = 3.5
	result := a + b  // Error: mismatched types

	✅ Compatible (after conversion):
	result := float64(a) + b  // Both float64

	Rules:
	1. Arithmetic: Same type required
	   int32 + int64  // ❌ Error
	   int64(x) + y   // ✅ Works

	2. Comparison: Same type required
	   int == int64   // ❌ Error
	   int64(x) == y  // ✅ Works

	3. Untyped constants are flexible:
	   var x int = 10
	   result := x + 5  // ✅ 5 becomes int

WHY Questions:
==============

Q1: Why does Go enforce type compatibility for operands in arithmetic and comparison operations?
A: 1. Type safety - prevents bugs at compile time

 2. Explicit intent - forces clear conversions

 3. Predictable behavior - no implicit conversions

 4. Performance - compiler knows exact types

 5. Early error detection - compile-time vs runtime

    Example:
    var a int = 5
    var b float64 = 2.5
    c := a / b  // ❌ Compile error (good!)

    vs other languages with implicit conversion:
    c := a / b  // Silently converts, might lose precision (bad!)

    Go forces you to be explicit:
    c := float64(a) / b  // ✅ Clear intent

Q2: Why is it beneficial for the Go compiler to perform optimizations for constant expressions?
A: 1. Runtime performance:

	   const area = 3.14159 * 10 * 10  // Calculated at compile time
	   // Runtime just uses 314.159, no multiplication

	2. Smaller binary size:
	   const secondsPerDay = 24 * 60 * 60  // Stored as 86400
	   // No calculation code in executable

	3. Compile-time error detection:
	   const x = 10 / 0  // ❌ Caught at compile time!

	4. No repeated calculations:
	   // Without constant:
	   for i := 0; i < 1000000; i++ {
	       x := 24 * 60 * 60  // Calculates 1M times!
	   }

	   // With constant:
	   const secondsPerDay = 24 * 60 * 60
	   for i := 0; i < 1000000; i++ {
	       x := secondsPerDay  // Just uses value
	   }

Q3: Why does short-circuit evaluation improve performance in logical operations?
A: 1. Skips unnecessary evaluations:

	   if false && expensiveFunction() {
	       // Never calls expensiveFunction()
	   }

	2. Avoids expensive operations:
	   if cachedValue || databaseQuery() {
	       // If cached, skips database call
	   }

	3. Real-world impact:
	   if user != nil && user.IsActive() && user.HasPermission() {
	       // Stops at first false condition
	       // Often just 1 check instead of 3
	   }

	Performance gains:
	- Fewer CPU cycles
	- Fewer memory accesses
	- Fewer function calls

Q4: Why is the division operator different for integers and floating-point numbers in Go?
A: 1. Mathematical correctness:

	   Integers can't store fractional parts
	   var x int = 10 / 3  // Must be whole: 3

	2. Predictable behavior:
	   7 / 2 = 3      (integer division)
	   7.0 / 2.0 = 3.5  (float division)

	3. Different use cases:
	   // Integer - counting items
	   itemsPerBox := 10 / 3  // 3 (can't split items)

	   // Float - measurements
	   speed := 10.0 / 3.0  // 3.333... (precise)

	4. Performance:
	   Integer division: Faster (simple CPU instruction)
	   Float division: Slower (more complex)

	Common pitfall:
	average := 5 / 2      // 2 (integer!)
	average := 5.0 / 2.0  // 2.5 (float)

HOW Questions:
==============

Q1: How does the Go compiler handle precedence and associativity among multiple operators in an expression?
A: Precedence (highest to lowest):

 1. Parentheses: ()

 2. Unary: +, -, !, ^, *, &

 3. Multiply/Divide: *, /, %, <<, >>, &, &^

 4. Add/Subtract: +, -, |, ^

 5. Comparison: ==, !=, <, <=, >, >=

 6. Logical AND: &&

 7. Logical OR: ||

    Associativity: Left-to-right for same precedence

    Example 1:
    result := 2 + 3 * 4
    // Step 1: 3 * 4 = 12 (* higher than +)
    // Step 2: 2 + 12 = 14

    Example 2:
    result := 10 - 5 - 2
    // Left-to-right: (10 - 5) - 2 = 3

    Example 3:
    a == b || a < b && b > 20
    // Step 1: Comparisons: (a == b), (a < b), (b > 20)
    // Step 2: && before ||: (a < b) && (b > 20)
    // Step 3: ||: (a == b) || (result)

    Use parentheses for clarity:
    result := a + (b * c)

Q2: How can logical operators && and || be used to simplify nested if conditions?
A: Nested ifs → Single condition with logical operators

	Example 1 - Multiple AND conditions:
	// Nested:
	if user != nil {
	    if user.Age >= 18 {
	        if user.IsActive {
	            fmt.Println("Valid")
	        }
	    }
	}

	// Simplified:
	if user != nil && user.Age >= 18 && user.IsActive {
	    fmt.Println("Valid")
	}

	Example 2 - Multiple OR conditions:
	// Nested:
	if role == "admin" {
	    grantAccess()
	} else if role == "moderator" {
	    grantAccess()
	} else if role == "editor" {
	    grantAccess()
	}

	// Simplified:
	if role == "admin" || role == "moderator" || role == "editor" {
	    grantAccess()
	}

	Example 3 - Early returns:
	// Nested:
	if user == nil {
	    return errors.New("user is nil")
	}
	if !user.IsActive {
	    return errors.New("not active")
	}

	// Simplified:
	if user == nil || !user.IsActive {
	    return errors.New("invalid user")
	}

Q3: How can you avoid common pitfalls when using comparison operators on floating-point numbers?
A: Problem: Floats have precision errors

	a := 0.1 + 0.2  // 0.30000000000000004
	b := 0.3
	a == b  // false! (unexpected)

	Solutions:

	1. Use epsilon comparison (tolerance):
	   const epsilon = 1e-9

	   func almostEqual(a, b float64) bool {
	       return math.Abs(a - b) < epsilon
	   }

	   if almostEqual(0.1 + 0.2, 0.3) {
	       // Works!
	   }

	2. Never use == for floats:
	   // ❌ Bad:
	   if price == 19.99 { }

	   // ✅ Good:
	   if math.Abs(price - 19.99) < 0.01 { }

	3. Use integer arithmetic:
	   // ❌ Float:
	   price := 19.99

	   // ✅ Integer (cents):
	   priceCents := 1999

	4. Round before comparing:
	   a := math.Round((0.1 + 0.2) * 100) / 100
	   b := 0.3
	   if a == b { }

	Key: Always use tolerance, never exact equality!

Q4: How does Go handle division by zero for integers and floating-point numbers?
A: Integer Division by Zero:

	x := 10 / 0  // ❌ Runtime PANIC!
	// panic: runtime error: integer divide by zero

	Float Division by Zero:
	x := 10.0 / 0.0   // +Inf (no panic)
	y := -10.0 / 0.0  // -Inf (no panic)
	z := 0.0 / 0.0    // NaN (no panic)

	Comparison:
	┌──────────┬──────────────┬──────────┐
	│ Type     │ Operation    │ Result   │
	├──────────┼──────────────┼──────────┤
	│ Integer  │ 10 / 0       │ PANIC 💥 │
	│ Float    │ 10.0 / 0.0   │ +Inf ✅  │
	│ Float    │ -10.0 / 0.0  │ -Inf ✅  │
	│ Float    │ 0.0 / 0.0    │ NaN ✅   │
	└──────────┴──────────────┴──────────┘

	How to handle:

	1. Check before dividing (integers):
	   if divisor != 0 {
	       result := dividend / divisor
	   }

	2. Check for special float values:
	   result := 10.0 / 0.0
	   if math.IsInf(result, 0) {
	       fmt.Println("Infinity")
	   }
	   if math.IsNaN(result) {
	       fmt.Println("NaN")
	   }

	3. Recover from panic (integers):
	   defer func() {
	       if r := recover(); r != nil {
	           fmt.Println("Recovered:", r)
	       }
	   }()
	   x := 10 / 0

KEY TAKEAWAYS:
==============
- Arithmetic operators calculate, comparison operators test
- Modulus (%) is essential for cyclic operations
- Short-circuit evaluation improves performance and safety
- Go enforces strict type compatibility (no implicit conversions)
- Constants are optimized at compile time
- Integer division truncates, float division preserves precision
- Operator precedence: *, / before +, -; && before ||
- Use logical operators to flatten nested ifs
- Never use == for floats - use epsilon comparison
- Integer division by zero panics, float returns Inf/NaN
*/

func main() {
	fmt.Println("=== GO QUIZ 3: Operators ===")
	fmt.Println("See comments above for all questions and answers!")
	fmt.Println()

	// Demonstrate key concepts
	fmt.Println("--- Arithmetic vs Comparison ---")
	a := 10
	b := 3
	fmt.Printf("Arithmetic: %d + %d = %d\n", a, b, a+b)
	fmt.Printf("Comparison: %d > %d = %v\n", a, b, a > b)

	fmt.Println("\n--- Modulus Operator ---")
	fmt.Printf("10 %% 3 = %d (remainder)\n", 10%3)
	fmt.Printf("Is 10 even? %v\n", 10%2 == 0)

	fmt.Println("\n--- Short-Circuit Evaluation ---")
	x := 0
	if x != 0 && 10/x > 2 {
		fmt.Println("Won't execute")
	} else {
		fmt.Println("Short-circuit prevented division by zero!")
	}

	fmt.Println("\n--- Integer vs Float Division ---")
	fmt.Printf("Integer: 10 / 3 = %d\n", 10/3)
	fmt.Printf("Float: 10.0 / 3.0 = %f\n", 10.0/3.0)

	fmt.Println("\n--- Float Comparison with Epsilon ---")
	const epsilon = 1e-9
	result := 0.1 + 0.2
	expected := 0.3
	fmt.Printf("0.1 + 0.2 == 0.3? %v (direct)\n", result == expected)
	fmt.Printf("With epsilon? %v (correct)\n", math.Abs(result-expected) < epsilon)

	fmt.Println("\n--- Division by Zero ---")
	fmt.Printf("Float: 10.0 / 0.0 = %v\n", 10.0/0.0) // +Inf
	fmt.Printf("Float: 0.0 / 0.0 = %v\n", 0.0/0.0)   // NaN
	// Integer division by zero would panic:
	// fmt.Println(10 / 0)  // Uncomment to see panic
}
