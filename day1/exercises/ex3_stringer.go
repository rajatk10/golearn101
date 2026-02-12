package main

import "fmt"

// The Stringer interface is defined in the fmt package:
// type Stringer interface {
//     String() string
// }
// When you implement String() method, fmt.Println will automatically use it

// TODO: Create a Book struct with:
// - Title (string)
// - Author (string)
// - Year (int)
// - Pages (int)

// TODO: Implement the String() method for Book
// Format: "Title by Author (Year) - Pages pages"
// Example: "The Go Programming Language by Donovan & Kernighan (2015) - 380 pages"

// TODO: Create a Product struct with:
// - Name (string)
// - Price (float64)
// - InStock (bool)

// TODO: Implement the String() method for Product
// Format: "Name - $Price (In Stock: Yes/No)"
// Example: "Laptop - $999.99 (In Stock: Yes)"

// TODO: Create a Temperature struct with:
// - Value (float64)
// - Unit (string) - either "C" or "F"

// TODO: Implement the String() method for Temperature
// Format: "Value°Unit"
// Example: "25.5°C" or "77.9°F"

// TODO: Implement a method ToFahrenheit() for Temperature (if Celsius)
// TODO: Implement a method ToCelsius() for Temperature (if Fahrenheit)
// Formula: F = C * 9/5 + 32
// Formula: C = (F - 32) * 5/9

func main() {
	// TODO: Create instances of Book, Product, and Temperature
	// TODO: Print them using fmt.Println (which will use String() method)
	// TODO: Test temperature conversion methods

	fmt.Println("Complete the exercise!")
}
