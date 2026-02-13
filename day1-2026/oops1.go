package main

import "fmt"

/*
GO'S APPROACH TO OBJECT-ORIENTED PROGRAMMING

Go doesn't have traditional OOP (classes, inheritance), but provides similar
functionality through structs, methods, interfaces, and composition.

┌─────────────────┬──────────────────────┬─────────────────────────────┐
│ OOP Concept     │ Go Equivalent        │ Description                 │
├─────────────────┼──────────────────────┼─────────────────────────────┤
│ Classes         │ Structs + Methods    │ Data + behavior together    │
│ Inheritance     │ Composition          │ Embed structs, not inherit  │
│ Polymorphism    │ Interfaces           │ Implicit interface impl     │
│ Encapsulation   │ Exported/Unexported  │ Capital = public, lower = private │
└─────────────────┴──────────────────────┴─────────────────────────────┘

DETAILED EXAMPLES:
==================

1. CLASSES → STRUCTS + METHODS
   Traditional OOP uses classes with methods. Go uses structs with methods.
*/

// Struct definition (like a class)
type Rectangle struct {
	Length, Width float64
}

// Methods on struct (like class methods)
func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Width)
}

// Pointer receiver - modifies original
func (r *Rectangle) Scale(factor float64) {
	r.Length *= factor
	r.Width *= factor
}

/*
2. INHERITANCE → COMPOSITION (EMBEDDING)
   Go doesn't support inheritance. Instead, use composition by embedding structs.
*/

// Base "class" equivalent
type Shape struct {
	Color string
	X, Y  float64 // Position
}

func (s Shape) Position() string {
	return fmt.Sprintf("Position: (%.2f, %.2f)", s.X, s.Y)
}

// "Derived class" equivalent - embeds Shape
type Circle struct {
	Shape  // Embedded struct (composition, not inheritance)
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

// Another "derived class"
type Square struct {
	Shape // Embedded struct
	Side  float64
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}

/*
3. POLYMORPHISM → INTERFACES
   Go uses implicit interfaces - no need to declare "implements"
*/

// Interface definition
type Drawable interface {
	Area() float64
	Position() string
}

// Any type with Area() and Position() methods automatically implements Drawable
// Rectangle, Circle, Square all implement Drawable implicitly

func PrintDrawableInfo(d Drawable) {
	fmt.Printf("Area: %.2f\n", d.Area())
	fmt.Println(d.Position())
}

/*
4. ENCAPSULATION → EXPORTED/UNEXPORTED
   Capital letter = exported (public)
   Lowercase letter = unexported (private)
*/

type BankAccount struct {
	Owner   string  // Exported (public)
	balance float64 // Unexported (private)
}

// Public method to access private field
func (b *BankAccount) GetBalance() float64 {
	return b.balance
}

// Public method to modify private field
func (b *BankAccount) Deposit(amount float64) {
	if amount > 0 {
		b.balance += amount
	}
}

func (b *BankAccount) Withdraw(amount float64) bool {
	if amount > 0 && amount <= b.balance {
		b.balance -= amount
		return true
	}
	return false
}

/*
COMPARISON WITH TRADITIONAL OOP:
=================================

Traditional OOP (Java/C++):
----------------------------
class Rectangle {
    private double length;
    private double width;
    
    public Rectangle(double l, double w) {
        length = l;
        width = w;
    }
    
    public double area() {
        return length * width;
    }
}

class ColoredRectangle extends Rectangle {
    private String color;
    
    public ColoredRectangle(double l, double w, String c) {
        super(l, w);
        color = c;
    }
}

Go Equivalent:
--------------
type Rectangle struct {
    length, width float64  // unexported (private)
}

func NewRectangle(l, w float64) Rectangle {
    return Rectangle{length: l, width: w}
}

func (r Rectangle) Area() float64 {
    return r.length * r.width
}

type ColoredRectangle struct {
    Rectangle  // Composition, not inheritance
    Color string
}

KEY DIFFERENCES:
================
1. No inheritance - use composition
2. Interfaces are implicit - no "implements" keyword
3. No constructors - use factory functions (NewXxx)
4. No method overloading - each method must have unique name
5. No "this" or "self" - use receiver variable name
6. Encapsulation via naming convention (capital vs lowercase)
*/

func main() {
	fmt.Println("=== GO OOP CONCEPTS ===\n")

	// 1. Structs + Methods (Classes)
	fmt.Println("--- 1. Structs + Methods ---")
	rect := Rectangle{Length: 5, Width: 3}
	fmt.Printf("Rectangle: %+v\n", rect)
	fmt.Printf("Area: %.2f\n", rect.Area())
	fmt.Printf("Perimeter: %.2f\n", rect.Perimeter())
	rect.Scale(2)
	fmt.Printf("After scaling: %+v\n\n", rect)

	// 2. Composition (Inheritance)
	fmt.Println("--- 2. Composition (Embedding) ---")
	circle := Circle{
		Shape:  Shape{Color: "red", X: 10, Y: 20},
		Radius: 5,
	}
	fmt.Printf("Circle: %+v\n", circle)
	fmt.Printf("Color: %s\n", circle.Color) // Access embedded field
	fmt.Println(circle.Position())          // Call embedded method
	fmt.Printf("Area: %.2f\n\n", circle.Area())

	square := Square{
		Shape: Shape{Color: "blue", X: 5, Y: 15},
		Side:  4,
	}
	fmt.Printf("Square: %+v\n", square)
	fmt.Printf("Area: %.2f\n\n", square.Area())

	// 3. Polymorphism (Interfaces)
	fmt.Println("--- 3. Polymorphism (Interfaces) ---")
	var drawable Drawable

	drawable = circle
	fmt.Println("Circle as Drawable:")
	PrintDrawableInfo(drawable)

	drawable = square
	fmt.Println("\nSquare as Drawable:")
	PrintDrawableInfo(drawable)

	// 4. Encapsulation
	fmt.Println("\n--- 4. Encapsulation ---")
	account := BankAccount{Owner: "Alice"}
	account.Deposit(1000)
	fmt.Printf("Account: %s, Balance: %.2f\n", account.Owner, account.GetBalance())

	account.Withdraw(300)
	fmt.Printf("After withdrawal: %.2f\n", account.GetBalance())

	// Cannot access private field directly:
	// fmt.Println(account.balance)  // Error: unexported field

	// Multiple types implementing same interface
	fmt.Println("\n--- Interface Flexibility ---")
	shapes := []Drawable{circle, square}
	for i, shape := range shapes {
		fmt.Printf("\nShape %d:\n", i+1)
		PrintDrawableInfo(shape)
	}
}
