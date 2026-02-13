package main

import (
	"fmt"
	"unsafe"
)

/*
UNSAFE PACKAGE NOTES:

The "unsafe" package allows operations that bypass Go's type safety and memory safety.
Use with EXTREME caution - can cause crashes, data corruption, and security issues.

WHY IT EXISTS:
- Low-level system programming
- Interfacing with C code
- Performance-critical operations
- Implementing runtime/compiler features

MAIN FUNCTIONS:

1. unsafe.Sizeof(x)
   - Returns size of variable x in bytes
   - Compile-time size, not heap-allocated size
   - For slices/strings/maps: returns header size only

2. unsafe.Alignof(x)
   - Returns alignment requirement of type
   - How memory address must be aligned

3. unsafe.Offsetof(x.f)
   - Returns offset of field f in struct x
   - Distance from start of struct to field

4. unsafe.Pointer
   - Generic pointer type (like void* in C)
   - Can convert between any pointer types
   - Bypasses Go's type system

WHEN TO USE:
✅ Measuring variable sizes (Sizeof)
✅ Interfacing with C libraries (cgo)
✅ Implementing low-level data structures
✅ Performance-critical code (rarely needed)

WHEN NOT TO USE:
❌ Regular application code
❌ When there's a safe alternative
❌ Without understanding consequences
❌ In production without thorough testing

DANGERS:
- Can corrupt memory
- Can cause crashes
- Breaks garbage collector assumptions
- Not portable across Go versions
- Violates type safety

SAFER ALTERNATIVES:
- Use reflect package for type inspection
- Use encoding/binary for byte manipulation
- Use standard library functions when available
*/

type Person struct {
	Name string // 16 bytes (string header)
	Age  int    // 8 bytes
	City string // 16 bytes (string header)
}

func main() {
	fmt.Println("=== UNSAFE PACKAGE EXAMPLES ===\n")

	// 1. Sizeof - Get size of variables
	fmt.Println("--- Sizeof Examples ---")
	var i int
	var i32 int32
	var i64 int64
	var f32 float32
	var f64 float64
	var b bool
	var s string
	var slice []int
	var arr [5]int

	fmt.Printf("int:       %d bytes\n", unsafe.Sizeof(i))
	fmt.Printf("int32:     %d bytes\n", unsafe.Sizeof(i32))
	fmt.Printf("int64:     %d bytes\n", unsafe.Sizeof(i64))
	fmt.Printf("float32:   %d bytes\n", unsafe.Sizeof(f32))
	fmt.Printf("float64:   %d bytes\n", unsafe.Sizeof(f64))
	fmt.Printf("bool:      %d bytes\n", unsafe.Sizeof(b))
	fmt.Printf("string:    %d bytes (header only)\n", unsafe.Sizeof(s))
	fmt.Printf("[]int:     %d bytes (slice header)\n", unsafe.Sizeof(slice))
	fmt.Printf("[5]int:    %d bytes (entire array)\n", unsafe.Sizeof(arr))

	// 2. Struct size and field offsets
	fmt.Println("\n--- Struct Layout ---")
	p := Person{Name: "Alice", Age: 30, City: "NYC"}
	fmt.Printf("Person struct size: %d bytes\n", unsafe.Sizeof(p))
	fmt.Printf("Name offset: %d bytes\n", unsafe.Offsetof(p.Name))
	fmt.Printf("Age offset:  %d bytes\n", unsafe.Offsetof(p.Age))
	fmt.Printf("City offset: %d bytes\n", unsafe.Offsetof(p.City))

	// 3. Alignment
	fmt.Println("\n--- Alignment Requirements ---")
	fmt.Printf("int alignment:     %d bytes\n", unsafe.Alignof(i))
	fmt.Printf("int64 alignment:   %d bytes\n", unsafe.Alignof(i64))
	fmt.Printf("float64 alignment: %d bytes\n", unsafe.Alignof(f64))
	fmt.Printf("Person alignment:  %d bytes\n", unsafe.Alignof(p))

	// 4. String vs actual content size
	fmt.Println("\n--- String Size vs Content ---")
	str := "Hello World"
	fmt.Printf("String header size: %d bytes\n", unsafe.Sizeof(str))
	fmt.Printf("Actual content:     %d bytes\n", len(str))

	// 5. Slice vs actual data size
	fmt.Println("\n--- Slice Size vs Data ---")
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Printf("Slice header size: %d bytes\n", unsafe.Sizeof(numbers))
	fmt.Printf("Actual data:       %d bytes (%d elements × %d bytes)\n",
		len(numbers)*int(unsafe.Sizeof(numbers[0])),
		len(numbers),
		unsafe.Sizeof(numbers[0]))

	fmt.Println("\n⚠️  WARNING: This package is called 'unsafe' for a reason!")
	fmt.Println("Use only when absolutely necessary and you understand the risks.")
}
