package main

/*
GO ARRAYS, SLICES, MAPS, STRUCTS QUIZ - Questions & Answers

WHAT Questions:
================

Q1: What is the difference between arrays and slices in Go?
A: Arrays: Fixed size, value type (copied), size is part of type ([3]int ≠ [5]int)
   Slices: Dynamic size, reference type (shares underlying array), has 3 components
   Example:
   arr := [3]int{1, 2, 3}  // Fixed size
   slice := []int{1, 2, 3}  // Can grow with append

Q2: What are the key components of a slice descriptor?
A: A slice header has 3 components (24 bytes total):
   1. Pointer (8 bytes) - points to underlying array
   2. Length (8 bytes) - number of elements currently in slice
   3. Capacity (8 bytes) - total space in underlying array
   
   Visual:
   ┌─────────────┬────────┬──────────┐
   │ Pointer     │ Length │ Capacity │
   └─────────────┴────────┴──────────┘
         ↓
         [1, 2, 3, 4, 5]

Q3: What is a map in Go, and how is it implemented internally?
A: Map is an unordered collection of key-value pairs (reference type).
   Internal implementation:
   - Hash table with buckets
   - Each bucket holds 8 key-value pairs
   - Hash function converts keys to bucket indices
   - Handles collisions with chaining
   - Auto-grows when load factor > 6.5
   - O(1) average lookup time

Q4: What is a struct, and how does it differ from a map?
A: Struct: Composite type with fixed, named fields of different types
   Map: Dynamic key-value pairs with same key type and same value type
   
   Comparison:
   ┌──────────────┬─────────────────┬──────────────────┐
   │ Feature      │ Struct          │ Map              │
   ├──────────────┼─────────────────┼──────────────────┤
   │ Type         │ Value (copied)  │ Reference        │
   │ Fields       │ Fixed, named    │ Dynamic keys     │
   │ Field types  │ Can mix types   │ All same type    │
   │ Access       │ struct.field    │ map[key]         │
   │ Memory       │ Contiguous      │ Hash table       │
   └──────────────┴─────────────────┴──────────────────┘


WHY Questions:
==============

Q1: Why would you choose a slice over an array in Go?
A: 1. Dynamic size - don't need to know size at compile time
   2. Flexibility - can grow/shrink as needed
   3. Function parameters - works with any size (arrays need exact size)
   4. Memory efficiency - reference type, no copying entire data
   5. Built-in operations - append(), copy(), slicing
   
   Real-world example: E-commerce orders per day (variable count)
   - Array: Must over-provision (wasteful)
   - Slice: Grows as needed (efficient)

Q2: Why is strict bounds checking important for arrays in Go?
A: 1. Memory safety - prevents access to unavailable memory
   2. Security - prevents buffer overflow attacks
   3. Predictable behavior - immediate panic vs undefined behavior
   4. Easier debugging - catches errors at runtime
   
   Comparison:
   C:  arr[10] = 5;  // No check - corrupts random memory! 💥
   Go: arr[10] = 5   // Panic immediately - safe! ✅

Q3: Why is a map preferred over a slice for key-value lookups?
A: MAIN REASON: Performance!
   - Map lookup: O(1) constant time (hash table)
   - Slice lookup: O(n) linear search (must check each element)
   
   Additional reasons:
   - Flexible keys (any comparable type vs only integers)
   - Semantic clarity (self-documenting keys)
   - Automatic uniqueness (keys are unique)


HOW Questions:
==============

Q1: How does the Go compiler handle memory allocation for slices?
A: 1. Initial allocation:
      s := make([]int, 5, 10)  // Allocates array of cap 10 on heap
   
   2. Growth strategy (when append exceeds capacity):
      - Small slices (< 1024): Double capacity
      - Large slices (≥ 1024): Grow by 25%
   
   3. Growth process:
      - Allocate new larger array
      - Copy old elements
      - Add new element
      - Update slice header
   
   Example:
   s := []int{1, 2, 3}     // cap = 3
   s = append(s, 4)        // cap = 6 (doubled!)
   s = append(s, 5, 6, 7)  // cap = 12 (doubled again!)

Q2: How can you dynamically increase the size of an array in Go?
A: TRICK QUESTION: You can't! Arrays are fixed size.
   
   Workarounds:
   1. Use a slice instead (recommended):
      slice := []int{1, 2, 3}
      slice = append(slice, 4, 5)  // Grows automatically
   
   2. Create new array and copy (manual, not dynamic):
      arr1 := [3]int{1, 2, 3}
      arr2 := [6]int{}  // Different type!
      copy(arr2[:], arr1[:])
   
   3. Convert array to slice:
      arr := [3]int{1, 2, 3}
      slice := arr[:]
      slice = append(slice, 4, 5)

Q3: How does Go ensure efficient hash table operations for maps?
A: 1. Fast hash functions - optimized for different key types
   2. Bucket structure - each bucket holds 8 key-value pairs
   3. Load factor management - auto-grows when load > 6.5
   4. Overflow handling - chains overflow buckets
   5. Incremental growth - spreads rehashing across operations
   
   Result: Maintains O(1) average performance for:
   - Lookup
   - Insert
   - Delete


KEY TAKEAWAYS:
==============
- Arrays: Fixed, value type, bounds-checked
- Slices: Dynamic, reference type, auto-growing
- Maps: Hash table, O(1) lookups, unordered
- Structs: Fixed fields, value type, contiguous memory
- Use slices for dynamic collections
- Use maps for key-value lookups
- Use structs for modeling entities
*/

import "fmt"

func main() {
	fmt.Println("=== GO QUIZ 2: Arrays, Slices, Maps, Structs ===")
	fmt.Println("See comments above for all questions and answers!")
	fmt.Println()

	// Demonstrate key concepts
	fmt.Println("--- Array vs Slice ---")
	arr := [3]int{1, 2, 3}
	slice := []int{1, 2, 3}
	fmt.Printf("Array: %v (fixed size)\n", arr)
	fmt.Printf("Slice: %v (can grow)\n", slice)
	slice = append(slice, 4, 5)
	fmt.Printf("After append: %v\n", slice)

	fmt.Println("\n--- Map O(1) Lookup ---")
	ages := map[string]int{"Alice": 30, "Bob": 25}
	fmt.Printf("Bob's age: %d (instant lookup!)\n", ages["Bob"])

	fmt.Println("\n--- Struct vs Map ---")
	type Person struct {
		Name string
		Age  int
	}
	p := Person{"Charlie", 35}
	fmt.Printf("Struct: %+v (fixed fields)\n", p)
	fmt.Printf("Map: %v (dynamic keys)\n", ages)
}
