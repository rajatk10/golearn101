package main

import "fmt"

/*
STRUCTS NOTES:

WHAT IS A STRUCT?
- Struct = Composite data type that groups related fields together
- Similar to classes in other languages (but without inheritance)
- Used to model real-world entities with multiple attributes
- Syntax: type Name struct { field type }

BASIC USAGE:
- Define: type Person struct { Name string; Age int }
- Create: p := Person{Name: "Alice", Age: 30}
- Access fields: p.Name, p.Age
- Cannot iterate over struct fields (not iterable)
- Print with %+v to see field names: fmt.Printf("%+v", struct)

STRUCTS ARE VALUE TYPES:
When you assign or pass a struct, it creates a COPY:
   p1 := Person{Name: "Bob", Age: 25}
   p2 := p1           // Full copy, not reference
   p2.Age = 30
   // p1.Age = 25 (unchanged)
   // p2.Age = 30

To share/modify original, use pointers:
   p1 := Person{Name: "Bob", Age: 25}
   p2 := &p1          // Pointer to p1
   p2.Age = 30
   // p1.Age = 30 (changed!)

MEMORY LAYOUT AND ALIGNMENT:

Go stores struct fields contiguously in memory. The order of fields affects:
1. Memory alignment (CPU efficiency)
2. Padding (internal gaps for alignment)
3. Total size of the struct

Example - Poor field ordering:
   type BadStruct struct {
       a bool    // 1 byte
       // 7 bytes padding
       b int64   // 8 bytes
       c bool    // 1 byte
       // 7 bytes padding
   }
   // Total: 24 bytes (with 14 bytes wasted on padding!)

Example - Good field ordering (largest to smallest):
   type GoodStruct struct {
       b int64   // 8 bytes
       a bool    // 1 byte
       c bool    // 1 byte
       // 6 bytes padding
   }
   // Total: 16 bytes (only 6 bytes padding)

ALIGNMENT RULES:
- bool, int8, uint8: 1-byte aligned
- int16, uint16: 2-byte aligned
- int32, uint32, float32: 4-byte aligned
- int64, uint64, float64, pointers: 8-byte aligned
- Structs are aligned to their largest field

Use unsafe.Sizeof() and unsafe.Alignof() to check:
   import "unsafe"
   fmt.Println("Size:", unsafe.Sizeof(myStruct))
   fmt.Println("Alignment:", unsafe.Alignof(myStruct))

MEMORY REPRESENTATION:
   type Person struct {
       Name string  // 16 bytes (string header: ptr + len)
       Age  int     // 8 bytes (on 64-bit)
   }
   
   Memory layout:
   ┌──────────────────┬─────────┐
   │ Name (16 bytes)  │ Age (8) │
   └──────────────────┴─────────┘
   Total: 24 bytes

BEST PRACTICES:
1. Order fields from largest to smallest to minimize padding
2. Group related fields together for readability
3. Use pointers for large structs to avoid copying
4. Use struct tags for JSON/XML serialization

FUNCTION vs METHOD:
Function:
- Standalone: func FunctionName(params) returnType { }
- Called directly: FunctionName(args)
- Example: RectangleArea(10, 20)

Method:
- Attached to a type: func (receiver Type) MethodName(params) returnType { }
- Called on instance: instance.MethodName(args)
- Has receiver (e Employee) before function name
- Example: emp1.Display()

KEY DIFFERENCE:
- Function: Independent, no receiver
- Method: Belongs to a type, has receiver

INTERFACES NOTES:
1. Interface = Contract that defines behavior (method signatures)
2. Syntax: type Name interface { MethodName() returnType }
3. Implicit implementation - no "implements" keyword needed
4. Any type that has the interface methods automatically satisfies the interface
5. Interface defines WHAT, not HOW
6. Used for polymorphism and abstraction

EXAMPLE:
type Shape interface {
    IArea() float64  // Any type with IArea() method is a Shape
}

type Circle struct { radius float64 }
func (c Circle) IArea() float64 { return 3.14 * c.radius * c.radius }
// Circle automatically implements Shape interface

INTERFACE vs STRUCT:
- Struct: Defines data (fields)
- Interface: Defines behavior (methods)
- Struct: Concrete type with implementation
- Interface: Abstract type, just signatures
*/

type Employee struct {
	firstName string
	lastName  string
	age       int
	gender    string
	empCode   int
}

func (e Employee) Display() { //Go Method
	fmt.Printf("First Name is : %s \n", e.firstName)
	fmt.Printf("Last Name is : %s \n", e.lastName)
	fmt.Printf("Age is : %d \n", e.age)
	fmt.Printf("Gender is : %s \n", e.gender)
	fmt.Printf("Employee Code is : %d \n", e.empCode)

}

// Another example for method and functions
func RectangleArea(a, b int) int {
	return a * b
}

type Rectangle struct {
	length int
	width  int
}

func (r Rectangle) Area() int {
	return r.length * r.width
}

type Shape interface {
	IArea() float64
}

type Circle struct {
	radius float64
}

func (c Circle) IArea() float64 {
	return 3.14 * c.radius * c.radius
}

func main() {
	emp1 := Employee{firstName: "Rajat", lastName: "Kumar", age: 31, gender: "Male"}
	fmt.Println(emp1)
	//print value
	fmt.Printf("%+v \n", emp1)
	// In go methods are those which has a defined type and used over it example Display()
	emp1.Display()
	//Rectangle area using function
	fmt.Printf("Area of Rectangle using function is : %d \n", RectangleArea(10, 20))

	//Rectangle area using method
	rect := Rectangle{15, 25}
	fmt.Printf("Area of Rectangle using method is : %d \n", rect.Area())

	//Now coming to interfaces.
	circle := Circle{5}
	fmt.Printf("Area of Circle using interface is : %f \n", circle.IArea())

}
