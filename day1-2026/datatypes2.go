package main

import "fmt"

/*
GO DATA TYPES CLASSIFICATION

Go has two main categories based on how data is stored and passed:

1. VALUE TYPES (Pass by Value - Copied)
   - Data is copied when assigned or passed to functions
   - Changes to copy don't affect original
   - Examples: int, float64, bool, string, arrays, structs

2. REFERENCE TYPES (Pass by Reference - Shared)
   - Store pointer/reference to underlying data
   - Changes affect all references to same data
   - Examples: slices, maps, channels, pointers, functions, interfaces

CLASSIFICATION TABLE:
┌─────────────────┬──────────────┬─────────────────────────────┐
│ Category        │ Type         │ Examples                    │
├─────────────────┼──────────────┼─────────────────────────────┤
│ VALUE TYPES     │ Basic        │ int, float64, bool, string  │
│ (Copied)        │ Array        │ [5]int, [3]string           │
│                 │ Struct       │ type Person struct {...}    │
├─────────────────┼──────────────┼─────────────────────────────┤
│ REFERENCE TYPES │ Slice        │ []int, []string             │
│ (Shared)        │ Map          │ map[string]int              │
│                 │ Channel      │ chan int, chan string       │
│                 │ Pointer      │ *int, *Person               │
│                 │ Function     │ func(), func(int) string    │
│                 │ Interface    │ interface{}, io.Reader      │
└─────────────────┴──────────────┴─────────────────────────────┘

MEMORY REPRESENTATION
arr1 := [3]int{1, 2, 3}
┌─────┬─────┬─────┐
│  1  │  2  │  3  │  ← Actual data in arr1
└─────┴─────┴─────┘

arr2 := arr1  // Full copy
┌─────┬─────┬─────┐
│  1  │  2  │  3  │  ← Separate copy in arr2
└─────┴─────┴─────┘

slice1 := []int{1, 2, 3}
┌─────────┐      ┌─────┬─────┬─────┐
│ pointer │─────→│  1  │  2  │  3  │  ← Actual data
└─────────┘      └─────┴─────┴─────┘
   slice1

slice2 := slice1  // Copy reference only
┌─────────┐      ┌─────┬─────┬─────┐
│ pointer │─────→│  1  │  2  │  3  │  ← Same data!
└─────────┘      └─────┴─────┴─────┘
   slice2           ↑
                    └── Both point to same data


KEY DIFFERENCES:

VALUE TYPE BEHAVIOR (Array):
   arr1 := [3]int{1, 2, 3}
   arr2 := arr1           // Full copy of data
   arr2[0] = 99
   // arr1[0] = 1 (unchanged)
   // arr2[0] = 99

REFERENCE TYPE BEHAVIOR (Slice):
   slice1 := []int{1, 2, 3}
   slice2 := slice1       // Copy reference only
   slice2[0] = 99
   // slice1[0] = 99 (changed!)
   // slice2[0] = 99 (both point to same data)

MEMORY REPRESENTATION:

Value Type (Array):
   arr1 := [3]int{1, 2, 3}
   [1][2][3]  ← Actual data in arr1

   arr2 := arr1  // Full copy
   [1][2][3]  ← Separate copy in arr2

Reference Type (Slice):
   slice1 := []int{1, 2, 3}
   [ptr] ──→ [1][2][3]  ← Actual data

   slice2 := slice1  // Copy reference
   [ptr] ──→ [1][2][3]  ← Same data!
            ↑
            Both point here

PRACTICAL IMPLICATIONS:

1. Function Parameters:
   // Value type - expensive for large arrays
   func process(arr [1000000]int) {  // Copies 1M integers!
   }

   // Reference type - efficient
   func process(slice []int) {  // Just copies pointer
   }

2. Safety vs Efficiency:
   Value Types:
   ✅ Safe (no shared state)
   ❌ Expensive to copy (large data)

   Reference Types:
   ✅ Efficient (share data)
   ⚠️  Can cause bugs (shared mutations)
   ⚠️  Need initialization (nil often unusable)

3. Nil Values:
   // Value types have zero values
   var i int        // 0
   var arr [3]int   // [0, 0, 0]

   // Reference types can be nil
   var slice []int           // nil (usable with append)
   var m map[string]int      // nil (NOT usable - panic!)
   var ch chan int           // nil (NOT usable - panic!)

RULE OF THUMB:
- Use arrays when size is fixed and small
- Use slices for dynamic/large collections
- Use pointers to share structs efficiently
- Use maps/channels for their specific purposes
- Prefer value types for safety, reference types for efficiency
*/

func main() {
	fmt.Println("=== VALUE TYPE EXAMPLE (Array) ===")
	arr1 := [3]int{1, 2, 3}
	arr2 := arr1 // Full copy
	arr2[0] = 99
	fmt.Printf("arr1: %v (unchanged)\n", arr1) // [1 2 3]
	fmt.Printf("arr2: %v (modified)\n", arr2)  // [99 2 3]

	fmt.Println("\n=== REFERENCE TYPE EXAMPLE (Slice) ===")
	slice1 := []int{1, 2, 3}
	slice2 := slice1 // Copy reference
	slice2[0] = 99
	fmt.Printf("slice1: %v (changed!)\n", slice1) // [99 2 3]
	fmt.Printf("slice2: %v (changed!)\n", slice2) // [99 2 3]

	fmt.Println("\n=== STRUCT (Value Type) ===")
	type Person struct {
		Name string
		Age  int
	}
	p1 := Person{"Alice", 30}
	p2 := p1 // Full copy
	p2.Age = 35
	fmt.Printf("p1: %+v (unchanged)\n", p1) // {Name:Alice Age:30}
	fmt.Printf("p2: %+v (modified)\n", p2)  // {Name:Alice Age:35}

	fmt.Println("\n=== MAP (Reference Type) ===")
	m1 := map[string]int{"a": 1, "b": 2}
	m2 := m1 // Copy reference
	m2["a"] = 99
	fmt.Printf("m1: %v (changed!)\n", m1) // map[a:99 b:2]
	fmt.Printf("m2: %v (changed!)\n", m2) // map[a:99 b:2]
}
