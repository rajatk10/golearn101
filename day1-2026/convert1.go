package main

import (
	"fmt"
	"strconv"
)

/*
TYPE CONVERSION NOTES:

Go requires EXPLICIT type conversion - no automatic conversions!

NUMERIC CONVERSIONS:
- int to float: float64(i)
- float to int: int(f) - TRUNCATES decimal (3.99 → 3)
- Between sizes: int64(i32), float64(f32)

STRING CONVERSIONS:
int/float → string:
- strconv.Itoa(i) - int to string
- strconv.FormatInt(i64, 10) - int64 to string (base 10)
- strconv.FormatFloat(f, 'f', 2, 64) - float to string (2 decimals)
- fmt.Sprintf("%d", i) - quick way

string → int/float:
- strconv.Atoi(s) - returns (int, error)
- strconv.ParseInt(s, 10, 64) - returns (int64, error)
- strconv.ParseFloat(s, 64) - returns (float64, error)

COMMON MISTAKE:
string(65) = "A" (ASCII character) ❌
strconv.Itoa(65) = "65" (number as string) ✅

STRING ↔ BYTES:
- []byte(s) - string to bytes
- string(b) - bytes to string
*/

func main() {
	// int to float
	i := 10
	f := float64(i)
	fmt.Printf("%d as float: %f\n", i, f)

	// float to int (truncates)
	f2 := 3.99
	i2 := int(f2)
	fmt.Printf("%.2f as int: %d (truncated)\n", f2, i2)

	// int to string
	num := 42
	str := strconv.Itoa(num)
	fmt.Printf("%d as string: '%s'\n", num, str)

	// string to int
	str2 := "123"
	num2, err := strconv.Atoi(str2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("'%s' as int: %d\n", str2, num2)
	}

	// float to string
	pi := 3.14159
	piStr := strconv.FormatFloat(pi, 'f', 2, 64)
	fmt.Printf("%.5f as string: '%s'\n", pi, piStr)

	// string to float
	str3 := "2.718"
	num3, err := strconv.ParseFloat(str3, 64)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("'%s' as float: %f\n", str3, num3)
	}

	// Common mistake: string() vs strconv.Itoa()
	ascii := 65
	wrong := string(ascii)           // "A" (ASCII character)
	right := strconv.Itoa(ascii)     // "65" (number as string)
	fmt.Printf("\nstring(%d) = '%s' ❌ (ASCII character)\n", ascii, wrong)
	fmt.Printf("strconv.Itoa(%d) = '%s' ✅ (number as string)\n", ascii, right)

	// string to bytes and back
	s := "Hello"
	bytes := []byte(s)
	backToStr := string(bytes)
	fmt.Printf("\n'%s' → %v → '%s'\n", s, bytes, backToStr)
}
