package main

import "fmt"

/*
POINTERS NOTES:
1. Pointer stores memory address of a variable
2. Declare pointer: var ptr *int (pointer to int)
3. Get address: &variable (& is address-of operator)
4. Dereference: *ptr (get value at address, * is dereference operator)
5. Modify via pointer: *ptr = newValue (changes original variable)
6. Zero value of pointer is nil
7. Pointers are used for:
   - Passing large structs efficiently
   - Modifying function parameters
   - Sharing data between functions

IMPORTANT TYPE MATCHING:
8. &a returns *int (pointer type) - NOT int
9. Must store address in pointer variable: var ptr *int = &a ✅
10. CANNOT store address in regular variable: var ptr int = &a ❌
11. Error: "cannot use &a (value of type *int) as int value"
12. Go is strictly typed - no automatic conversion between int and *int
13. Type inference works: ptr := &a (Go infers ptr is *int)

REFERENCE AND DEREFERENCE OPERATIONS:
┌─────────────────────┬──────────┬─────────────────────────────┐
│ Operation           │ Syntax   │ Meaning                     │
├─────────────────────┼──────────┼─────────────────────────────┤
│ Reference           │ &x       │ Get address of x            │
│ Dereference (read)  │ *p       │ Get value at address p      │
│ Dereference (write) │ *p = 100 │ Set value at address p      │
│ Declare pointer     │ var p *int│ p is a pointer to int      │
│ Allocate memory     │ p = new(int)│ Create new int, return address │
└─────────────────────┴──────────┴─────────────────────────────┘

MEMORY ALLOCATION:
14. var p *int       // Declared but not initialized (nil)
15. p = new(int)     // Allocates memory and returns pointer
16. *p = 100         // Dereference and assign value

Visual representation:
Step 1: var p *int
  p → nil (points nowhere)

Step 2: p = new(int)
  p → [0] (allocated memory, initialized to zero value)

Step 3: *p = 100
  p → [100] (value changed via dereference)

POINTER TYPES:
┌─────────────────────┬──────────────────────────────────────────┐
│ Pointer Type        │ Description                              │
├─────────────────────┼──────────────────────────────────────────┤
│ *int                │ Pointer to int                           │
│ *string             │ Pointer to string                        │
│ *struct             │ Pointer to a custom struct               │
│ nil pointer         │ Zero-value pointer (not initialized)     │
│ *[n]T               │ Pointer to array (entire array mods)    │
│ *interface{}        │ Pointer to interface (rare, needs type assertions) │
└─────────────────────┴──────────────────────────────────────────┘

Examples:
var p1 *int              // Pointer to int
var p2 *string           // Pointer to string
var p3 *MyStruct         // Pointer to struct
var p4 *[3]int           // Pointer to array of 3 ints
var p5 *interface{}      // Pointer to interface (rare)

All uninitialized pointers have value: nil
*/

func main() {
	var a int = 100
	var ptr *int = &a
	fmt.Println("Value of a : ", a)
	fmt.Println("Address of a : ", &a)
	fmt.Println("Value of ptr : ", ptr)
	fmt.Println("Address of ptr : ", &ptr)
	fmt.Println("Value at address stored in ptr : ", *ptr)
	*ptr = 200
	fmt.Println("Value of a after changing ptr : ", a)
}
