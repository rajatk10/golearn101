package main

import "fmt"

/*
ARRAYS & SLICES NOTES:
ARRAYS:
1. Fixed size: var arr [10]int - size is part of the type
2. Cannot be resized or appended
3. Passed by value (copied) to functions
4. Access: arr[index], Length: len(arr)

SLICES:
1. Dynamic size: var slice []int - no size specified
2. Can grow with append(): slice = append(slice, value)
3. Passed by reference (shares underlying array)
4. Three ways to create:
   - Literal: slice := []int{1, 2, 3}
   - make(): slice := make([]int, length, capacity)
   - From array: slice := arr[start:end]
5. len(slice) - number of elements
6. cap(slice) - capacity (reserved memory)
7. Capacity growth: doubles for small slices (< 256), ~25% for larger
8. Pre-allocate capacity for performance: make([]int, 0, 100)
*/

func main() {
	var arr [10]int
	for i := 0; i < len(arr); i++ {
		arr[i] = i * 10
		fmt.Println("Element arr [", i, "] is: ", arr[i])
	}
	//brr = append(arr, 100) // array cannot be appended

	//slices
	var srr []int = []int{1, 2, 3, 4, 5}
	fmt.Println("srr is: ", srr)
	fmt.Println("Length of srr is: ", len(srr))
	fmt.Println("Capacity of srr is: ", cap(srr))
	srr = append(srr, 6)
	fmt.Println("srr is: ", srr)
	fmt.Println("Length of srr is: ", len(srr))
	fmt.Println("Capacity of srr is: ", cap(srr))
}
