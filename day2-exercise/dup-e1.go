package main

import "fmt"

/*
VARIADIC FUNCTIONS AND ... OPERATOR NOTES:

The ... (spread/variadic) operator has two main uses in Go:

1. VARIADIC FUNCTIONS - Accept variable number of arguments:
   func sum(nums ...int) int {
       total := 0
       for _, num := range nums {
           total += num
       }
       return total
   }
   
   // Call with any number of arguments
   sum(1, 2, 3)           // 6
   sum(1, 2, 3, 4, 5)     // 15
   
   // Or pass a slice with ...
   numbers := []int{1, 2, 3}
   sum(numbers...)        // 6

2. UNPACKING SLICES - When appending or passing to variadic functions:
   a := []int{1, 2, 3}
   b := []int{4, 5, 6}
   
   // ❌ Wrong - tries to append slice as single element
   result := append(a, b)  // Error!
   
   // ✅ Correct - unpacks b into individual elements
   result := append(a, b...)  // [1 2 3 4 5 6]

COMMON USES:
- append(slice, otherSlice...)  // Merge slices
- fmt.Println(args...)          // Pass multiple args
- myFunc(slice...)              // Pass slice to variadic function

KEY POINT:
... unpacks a slice into individual elements, allowing you to pass/append
multiple values at once.
*/

// Remove duplicate elements from slice
// a = [1,2,3,2,3,4]
// output = [1,2,3,4]
func removeDuplicate(a []int) []int {
	seen := make(map[int]bool)
	result := []int{}

	for _, num := range a {
		if !seen[num] {
			seen[num] = true
			result = append(result, num)
		}
	}

	return result
}

func mergeTwoSlices(a []int, b []int) []int {
	return append(a, b...)
}

func main() {
	a := []int{1, 2, 3, 2, 3, 4}
	fmt.Printf("Original: %v\n", a)
	fmt.Printf("After removing duplicates: %v\n", removeDuplicate(a))
	s1 := make([]int, 0, 10)
	s2 := make([]int, 0, 10)
	s1 = append(s1, 1, 2, 3)
	s2 = append(s2, 3, 4, 5)
	fmt.Printf("Merged: %v\n", mergeTwoSlices(s1, s2))
}
