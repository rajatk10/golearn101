package main

import "fmt"

/*
OPERATORS IN GO

1. ARITHMETIC OPERATORS:
   Used for mathematical calculations on numeric types
   
   Operator | Description      | Example
   ---------|-----------------|----------
   +        | Addition        | 5 + 3 = 8
   -        | Subtraction     | 5 - 3 = 2
   *        | Multiplication  | 5 * 3 = 15
   /        | Division        | 10 / 3 = 3 (integer), 10.0 / 3.0 = 3.333... (float)
   %        | Modulus         | 10 % 3 = 1 (remainder)
   
   Note: Integer division truncates, float division preserves precision

2. COMPARISON OPERATORS:
   Used to compare values, returns boolean (true/false)
   
   Operator | Description              | Example
   ---------|-------------------------|----------
   ==       | Equal to                | 5 == 5 → true
   !=       | Not equal to            | 5 != 3 → true
   <        | Less than               | 3 < 5 → true
   >        | Greater than            | 5 > 3 → true
   <=       | Less than or equal      | 3 <= 5 → true
   >=       | Greater than or equal   | 5 >= 5 → true
   
   Note: Both operands must be same type (type compatible)

3. LOGICAL OPERATORS:
   Used to combine boolean expressions
   
   Operator | Description | Example              | Short-circuit?
   ---------|-------------|---------------------|---------------
   &&       | AND         | true && false → false | Yes (stops at first false)
   ||       | OR          | true || false → true  | Yes (stops at first true)
   !        | NOT         | !true → false        | No
   
   Short-circuit evaluation:
   - && stops if left side is false (result will be false)
   - || stops if left side is true (result will be true)

4. OPERATOR PRECEDENCE (Highest to Lowest):
   
   Level | Operators                    | Description
   ------|------------------------------|---------------------------
   1     | ()                           | Parentheses (highest)
   2     | !, +, -, *, & (unary)       | Unary operators
   3     | *, /, %, <<, >>, &, &^      | Multiplication, Division, Modulus
   4     | +, -, |, ^                   | Addition, Subtraction
   5     | ==, !=, <, <=, >, >=        | Comparison operators
   6     | &&                           | Logical AND
   7     | ||                           | Logical OR (lowest)
   
   ASSOCIATIVITY: All binary operators are LEFT-TO-RIGHT
   
   Examples:
   
   a) 2 + 3 * 4
      Step 1: 3 * 4 = 12  (* higher precedence than +)
      Step 2: 2 + 12 = 14
   
   b) 10 - 5 - 2
      Left-to-right: (10 - 5) - 2 = 5 - 2 = 3
   
   c) a == b || a < b && b > 20
      Step 1: Comparisons: (a == b), (a < b), (b > 20)
      Step 2: && before ||: (a < b) && (b > 20)
      Step 3: ||: (a == b) || (result from step 2)
      Final: (a == b) || ((a < b) && (b > 20))
   
   d) 5 + 3 * 2 - 8 / 4
      Step 1: 3 * 2 = 6
      Step 2: 8 / 4 = 2
      Step 3: 5 + 6 = 11
      Step 4: 11 - 2 = 9

KEY POINTS:
- Multiplication/Division/Modulus have SAME precedence (evaluated left-to-right)
- Addition/Subtraction have SAME precedence (evaluated left-to-right)
- && has HIGHER precedence than ||
- Use parentheses for clarity when in doubt!
*/

func main() {
	fmt.Println("=== OPERATORS IN GO ===")

	// Arithmetic Operators
	fmt.Println("\n--- Arithmetic Operators ---")
	a, b := 10, 3
	fmt.Printf("%d + %d = %d\n", a, b, a+b)
	fmt.Printf("%d - %d = %d\n", a, b, a-b)
	fmt.Printf("%d * %d = %d\n", a, b, a*b)
	fmt.Printf("%d / %d = %d (integer division)\n", a, b, a/b)
	fmt.Printf("%.1f / %.1f = %.2f (float division)\n", 10.0, 3.0, 10.0/3.0)
	fmt.Printf("%d %% %d = %d (modulus)\n", a, b, a%b)

	// Comparison Operators
	fmt.Println("\n--- Comparison Operators ---")
	x, y := 5, 5
	fmt.Printf("%d == %d → %v\n", x, y, x == y)
	fmt.Printf("%d != %d → %v\n", x, 3, x != 3)
	fmt.Printf("%d < %d → %v\n", 3, x, 3 < x)
	fmt.Printf("%d > %d → %v\n", x, 3, x > 3)
	fmt.Printf("%d <= %d → %v\n", x, y, x <= y)
	fmt.Printf("%d >= %d → %v\n", x, y, x >= y)

	// Logical Operators
	fmt.Println("\n--- Logical Operators ---")
	fmt.Printf("true && true = %v\n", true && true)
	fmt.Printf("true && false = %v\n", true && false)
	fmt.Printf("true || false = %v\n", true || false)
	fmt.Printf("false || false = %v\n", false || false)
	fmt.Printf("!true = %v\n", !true)

	// Short-circuit evaluation
	fmt.Println("\n--- Short-Circuit Evaluation ---")
	num := 0
	if num != 0 && 10/num > 2 {
		fmt.Println("Won't execute")
	} else {
		fmt.Println("Short-circuit prevented division by zero!")
	}

	// Operator Precedence
	fmt.Println("\n--- Operator Precedence ---")
	result1 := 2 + 3*4
	fmt.Printf("2 + 3 * 4 = %d (multiplication first)\n", result1)

	result2 := 10 - 5 - 2
	fmt.Printf("10 - 5 - 2 = %d (left-to-right)\n", result2)

	result3 := 5 + 3*2 - 8/4
	fmt.Printf("5 + 3 * 2 - 8 / 4 = %d\n", result3)

	// Complex logical expression
	fmt.Println("\n--- Complex Logical Expression ---")
	p, q, r := 5, 10, 25
	result4 := p == q || p < q && q < r
	fmt.Printf("%d == %d || %d < %d && %d < %d = %v\n", p, q, p, q, q, r, result4)
	fmt.Printf("Evaluated as: (%d == %d) || ((%d < %d) && (%d < %d))\n", p, q, p, q, q, r)
}
