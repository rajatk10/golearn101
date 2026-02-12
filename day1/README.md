# Day 1: Advanced Data Structures & Methods

## 🎯 Today's Goals

By the end of today, you will:
- Master structs with methods
- Understand pointer vs value receivers
- Implement and use interfaces
- Build a complete Task Management System

---

## 📚 Part 1: Structs & Methods (45 min)

### What are Structs?

Structs are Go's way of creating custom types that group related data together.

```go
type Person struct {
    Name string
    Age  int
    Email string
}
```

### Methods

Methods are functions attached to types. They give behavior to your data.

```go
func (p Person) Greet() string {
    return "Hello, I'm " + p.Name
}
```

### Pointer Receivers vs Value Receivers

**Value Receiver**: Gets a copy of the struct
```go
func (p Person) GetAge() int {
    return p.Age  // Can read but changes won't persist
}
```

**Pointer Receiver**: Gets a reference to the original struct
```go
func (p *Person) SetAge(age int) {
    p.Age = age  // Changes WILL persist
}
```

**When to use which?**
- Use pointer receiver when you need to modify the struct
- Use pointer receiver for large structs (avoid copying)
- Use value receiver for small, immutable data
- Be consistent: if one method uses pointer, all should

---

## 📚 Part 2: Interfaces (45 min)

### What are Interfaces?

Interfaces define behavior (what something can do), not data.

```go
type Speaker interface {
    Speak() string
}
```

Any type that implements `Speak()` automatically satisfies the `Speaker` interface.

### Interface Implementation

```go
type Dog struct {
    Name string
}

func (d Dog) Speak() string {
    return "Woof!"
}

type Cat struct {
    Name string
}

func (c Cat) Speak() string {
    return "Meow!"
}

// Both Dog and Cat implement Speaker interface
```

### Empty Interface

`interface{}` (or `any` in Go 1.18+) can hold any value:

```go
func PrintAnything(v interface{}) {
    fmt.Println(v)
}
```

### Type Assertions

Extract the concrete type from an interface:

```go
var i interface{} = "hello"

s := i.(string)  // Type assertion
fmt.Println(s)

// Safe type assertion
s, ok := i.(string)
if ok {
    fmt.Println(s)
}
```

### Type Switch

```go
func describe(i interface{}) {
    switch v := i.(type) {
    case int:
        fmt.Printf("Integer: %d\n", v)
    case string:
        fmt.Printf("String: %s\n", v)
    default:
        fmt.Printf("Unknown type\n")
    }
}
```

---

## 📚 Part 3: Type Embedding & Composition (30 min)

### Struct Embedding

Go doesn't have inheritance, but it has composition through embedding:

```go
type Address struct {
    Street string
    City   string
}

type Employee struct {
    Name    string
    Address // Embedded struct
}

// You can access embedded fields directly
emp := Employee{
    Name: "John",
    Address: Address{
        Street: "123 Main St",
        City:   "NYC",
    },
}

fmt.Println(emp.City)  // Direct access!
```

### Method Promotion

Methods from embedded types are promoted:

```go
type Base struct {
    ID int
}

func (b Base) GetID() int {
    return b.ID
}

type Derived struct {
    Base
    Name string
}

d := Derived{Base: Base{ID: 1}, Name: "Test"}
fmt.Println(d.GetID())  // Method promoted from Base!
```

---

## 💻 Practice Exercises

### Exercise 1: Person Struct
Create in `exercises/ex1_person.go`

### Exercise 2: Shape Calculator
Create in `exercises/ex2_shapes.go`

### Exercise 3: Stringer Interface
Create in `exercises/ex3_stringer.go`

---

## 🚀 Day 1 Project: Task Management System

Build a complete task management system with:
- Different task types (Work, Personal, Shopping)
- CRUD operations
- Interface-based design
- Priority levels

See `project/` folder for starter code.

---

## ✅ Daily Checklist

- [ ] Read all theory sections
- [ ] Complete Exercise 1: Person Struct
- [ ] Complete Exercise 2: Shape Calculator
- [ ] Complete Exercise 3: Stringer Interface
- [ ] Build the Task Management System
- [ ] Run all tests and ensure they pass
- [ ] Format code with `go fmt`
- [ ] Review and understand every line

---

## 🔑 Key Takeaways

1. **Structs** group related data
2. **Methods** give behavior to types
3. **Pointer receivers** modify, **value receivers** read
4. **Interfaces** define contracts
5. **Embedding** enables composition
6. Go favors **composition over inheritance**

---

## 📖 Additional Resources

- [Effective Go - Methods](https://go.dev/doc/effective_go#methods)
- [Go by Example - Interfaces](https://gobyexample.com/interfaces)
- [Go Tour - Methods and Interfaces](https://go.dev/tour/methods/1)

---

**Ready? Start with the examples folder, then move to exercises!**
