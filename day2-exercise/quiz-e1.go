package main

/*
GO DATA TYPES QUIZ - Questions & Answers

WHAT Questions:
================

Q1: What is the difference between int and int32 in Go?
A: int is architecture-dependent (32-bit or 64-bit based on system).
   int32 is always 32-bit (4 bytes) on all systems.
   Use int for general purpose, int32 when size must be guaranteed.

Q2: What is a rune, and how does it differ from a byte?
A: rune is alias for int32, represents a Unicode code point (any character).
   byte is alias for uint8, represents a single byte (0-255, ASCII only).
   Use rune for Unicode text, byte for binary data.

Q3: What is the default size of the float type in Go?
A: There is NO "float" type in Go. Must use float32 or float64.
   Type inference (x := 3.14) defaults to float64.

Q4: What is the significance of immutability in Go strings?
A: Strings cannot be modified after creation.
   Benefits: thread-safe, memory efficient (can share), safe as map keys.
   To modify: convert to []byte, modify, convert back to string.


WHY Questions:
==============

Q1: Why is float64 preferred over float32 in most applications?
A: Better precision (~15 vs ~7 decimal digits).
   Standard library uses float64.
   No performance penalty on modern 64-bit systems.
   Use float32 only for memory-critical scenarios (graphics, large arrays).

Q2: Why does Go enforce strict type safety for bool values?
A: Prevents bugs from implicit conversions (no truthy/falsy values).
   Only true and false are valid - no numbers, strings, or pointers.
   Makes code explicit and prevents accidental assignment in conditions.

Q3: Why are strings stored as UTF-8 in Go?
A: Universal character support (all languages, emoji).
   Backward compatible with ASCII (1 byte per ASCII char).
   Space efficient (variable length encoding).
   Web standard - works directly with HTTP/JSON/APIs.

Q4: Why is it important to consider architecture dependency when using int?
A: int size varies (4 bytes on 32-bit, 8 bytes on 64-bit).
   Can cause overflow on 32-bit systems.
   Breaks serialization/network protocols across platforms.
   Use int for local operations, int32/int64 for portable code.


HOW Questions:
==============

Q1: How does Go handle type conversions between int, float, and string?
A: Requires EXPLICIT conversion - no automatic conversions.
   int ↔ float: float64(i), int(f)
   int → string: strconv.Itoa(i), fmt.Sprintf("%d", i)
   string → int: strconv.Atoi(s) - returns (int, error)
   float → string: strconv.FormatFloat(f, 'f', 2, 64)
   string → float: strconv.ParseFloat(s, 64) - returns (float64, error)

Q2: How can you iterate over runes in a string to process Unicode characters?
A: Use range loop - automatically decodes UTF-8 to runes:
   for i, r := range str {
       // i = byte index, r = rune (Unicode code point)
   }
   Or convert to []rune: runes := []rune(str)

Q3: How does Go ensure memory safety while handling strings?
A: Strings are immutable - cannot be modified after creation.
   Prevents data races in concurrent code.
   Bounds checking on string indexing (panics on out-of-bounds).
   Slicing shares memory safely (original can't change).

Q4: How can you determine the size of a variable of any type at runtime?
A: Use unsafe.Sizeof(variable) - returns size in bytes.
   import "unsafe"
   size := unsafe.Sizeof(myVar)
   Note: Returns compile-time size, not heap-allocated size.
*/

import (
	"fmt"
	"strconv"
	"unsafe"
)

func main() {
	fmt.Println("=== GO DATA TYPES QUIZ ===")
	fmt.Println("See comments above for all questions and answers!")
	fmt.Println()

	// Example demonstrations
	fmt.Println("--- Type Size Examples ---")
	var i int
	var i32 int32
	var f64 float64
	var r rune
	var b byte
	fmt.Printf("int size: %d bytes\n", unsafe.Sizeof(i))
	fmt.Printf("int32 size: %d bytes\n", unsafe.Sizeof(i32))
	fmt.Printf("float64 size: %d bytes\n", unsafe.Sizeof(f64))
	fmt.Printf("rune size: %d bytes\n", unsafe.Sizeof(r))
	fmt.Printf("byte size: %d bytes\n", unsafe.Sizeof(b))

	fmt.Println("\n--- String Immutability ---")
	s := "Hello"
	fmt.Printf("Original: %s\n", s)
	// s[0] = 'h'  // Error: cannot assign
	bytes := []byte(s)
	bytes[0] = 'h'
	s = string(bytes)
	fmt.Printf("Modified: %s\n", s)

	fmt.Println("\n--- Type Conversions ---")
	num := 42
	str := strconv.Itoa(num)
	fmt.Printf("int %d → string '%s'\n", num, str)

	str2 := "123"
	num2, _ := strconv.Atoi(str2)
	fmt.Printf("string '%s' → int %d\n", str2, num2)

	fmt.Println("\n--- Rune vs Byte ---")
	text := "Go🚀"
	fmt.Printf("String: %s\n", text)
	fmt.Printf("Bytes: %d\n", len(text))
	fmt.Printf("Runes: %d\n", len([]rune(text)))
}
