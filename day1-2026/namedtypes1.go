package main

import "fmt"

/*
NAMED TYPES IN GO - Complete Examples

A named type is any type you declare using the 'type' keyword.
You can define methods on any named type in the same package.

CATEGORIES OF NAMED TYPES:
===========================
1. Basic types (aliases)
2. Structs
3. Interfaces
4. Arrays
5. Slices
6. Maps
7. Channels
8. Functions
9. Pointers
*/

// ============================================================================
// 1. BASIC TYPES (ALIASES)
// ============================================================================

type MyInt int
type Temperature float64
type UserID string
type Flag bool

// Methods on basic types
func (m MyInt) Double() MyInt {
	return m * 2
}

func (t Temperature) ToFahrenheit() float64 {
	return float64(t)*9/5 + 32
}

func (u UserID) IsValid() bool {
	return len(u) > 0
}

func (f Flag) Toggle() Flag {
	return !f
}

// ============================================================================
// 2. STRUCTS
// ============================================================================

type Person1 struct {
	Name string
	Age  int
}

func (p Person1) Greet() string {
	return fmt.Sprintf("Hello, I'm %s, %d years old", p.Name, p.Age)
}

func (p *Person1) HaveBirthday() {
	p.Age++
}

// ============================================================================
// 3. INTERFACES
// ============================================================================

type Speaker interface {
	Speak() string
}

type Animal interface {
	Sound() string
	Move() string
}

// Dog implements Animal
type Dog struct {
	Name string
}

func (d Dog) Sound() string {
	return "Woof!"
}

func (d Dog) Move() string {
	return "Running"
}

// ============================================================================
// 4. ARRAYS (FIXED SIZE)
// ============================================================================

type Matrix [3][3]int
type RGB [3]uint8
type Coordinates [2]float64

func (m Matrix) Sum() int {
	total := 0
	for _, row := range m {
		for _, val := range row {
			total += val
		}
	}
	return total
}

func (rgb RGB) ToHex() string {
	return fmt.Sprintf("#%02X%02X%02X", rgb[0], rgb[1], rgb[2])
}

func (c Coordinates) Distance() float64 {
	return c[0]*c[0] + c[1]*c[1] // Simplified distance
}

// ============================================================================
// 5. SLICES (DYNAMIC SIZE)
// ============================================================================

type IntSlice []int
type StringList []string
type ScoreList []float64

func (s IntSlice) Sum() int {
	total := 0
	for _, v := range s {
		total += v
	}
	return total
}

func (s IntSlice) Average() float64 {
	if len(s) == 0 {
		return 0
	}
	return float64(s.Sum()) / float64(len(s))
}

func (sl StringList) Contains(str string) bool {
	for _, s := range sl {
		if s == str {
			return true
		}
	}
	return false
}

func (scores ScoreList) Max() float64 {
	if len(scores) == 0 {
		return 0
	}
	max := scores[0]
	for _, score := range scores {
		if score > max {
			max = score
		}
	}
	return max
}

// ============================================================================
// 6. MAPS
// ============================================================================

type StringMap map[string]string
type UserCache map[int]Person1
type Counter map[string]int

func (sm StringMap) Keys() []string {
	keys := make([]string, 0, len(sm))
	for k := range sm {
		keys = append(keys, k)
	}
	return keys
}

func (uc UserCache) GetByName(name string) (Person1, bool) {
	for _, person := range uc {
		if person.Name == name {
			return person, true
		}
	}
	return Person1{}, false
}

func (c Counter) Increment(key string) {
	c[key]++
}

func (c Counter) Total() int {
	total := 0
	for _, count := range c {
		total += count
	}
	return total
}

// ============================================================================
// 7. CHANNELS
// ============================================================================

type IntChannel chan int
type ErrorChannel chan error
type MessageChannel chan string

func (ic IntChannel) SendRange(start, end int) {
	go func() {
		for i := start; i <= end; i++ {
			ic <- i
		}
		close(ic)
	}()
}

// ============================================================================
// 8. FUNCTIONS (FUNCTION TYPES)
// ============================================================================

type Handler func(string) error
type Comparator func(int, int) bool
type Transformer func(string) string
type Validator func(interface{}) bool

// Method on function type
func (h Handler) WithLogging() Handler {
	return func(s string) error {
		fmt.Printf("Handling: %s\n", s)
		err := h(s)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
		return err
	}
}

// ============================================================================
// 9. POINTERS (Note: Cannot define methods on pointer types)
// ============================================================================

// You CANNOT define methods on pointer types like this:
// type IntPtr *int
// func (ip IntPtr) Method() { }  // ❌ Error: invalid receiver type

// Instead, use a struct wrapper if you need methods:
type IntWrapper struct {
	Value *int
}

func (iw IntWrapper) IsNil() bool {
	return iw.Value == nil
}

func (iw IntWrapper) Get() int {
	if iw.Value == nil {
		return 0
	}
	return *iw.Value
}

func (iw *IntWrapper) Set(val int) {
	if iw.Value == nil {
		iw.Value = new(int)
	}
	*iw.Value = val
}

// ============================================================================
// SUMMARY TABLE
// ============================================================================

/*
┌──────────────────┬─────────────────────┬──────────────────────────────┐
│ Category         │ Example             │ Use Case                     │
├──────────────────┼─────────────────────┼──────────────────────────────┤
│ Basic Types      │ type MyInt int      │ Add methods to primitives    │
│ Structs          │ type Person struct  │ Data modeling                │
│ Interfaces       │ type Reader interface│ Polymorphism                │
│ Arrays           │ type Matrix [3][3]int│ Fixed-size collections      │
│ Slices           │ type IntSlice []int │ Dynamic collections          │
│ Maps             │ type Cache map[k]v  │ Key-value storage            │
│ Channels         │ type IntChan chan int│ Concurrency                 │
│ Functions        │ type Handler func() │ Callbacks, strategies        │
│ Pointers         │ type IntPtr *int    │ Optional values (rare)       │
└──────────────────┴─────────────────────┴──────────────────────────────┘

KEY RULES:
==========
1. You can define methods on ANY named type
2. Named type must be in the same package as the method
3. Cannot define methods on built-in types directly (must create named type)
4. Receiver can be value or pointer

Examples:
✅ type MyInt int; func (m MyInt) Double() MyInt { }
❌ func (i int) Double() int { }  // Cannot add method to built-in type
*/

func main() {
	fmt.Println("=== NAMED TYPES IN GO ===\n")

	// 1. Basic Types
	fmt.Println("--- 1. Basic Types ---")
	num := MyInt(5)
	fmt.Printf("MyInt: %d, Doubled: %d\n", num, num.Double())

	temp := Temperature(25)
	fmt.Printf("Temperature: %.1f°C = %.1f°F\n", temp, temp.ToFahrenheit())

	id := UserID("user123")
	fmt.Printf("UserID: %s, Valid: %v\n", id, id.IsValid())

	flag := Flag(true)
	fmt.Printf("Flag: %v, Toggled: %v\n\n", flag, flag.Toggle())

	// 2. Structs
	fmt.Println("--- 2. Structs ---")
	person := Person1{Name: "Alice", Age: 30}
	fmt.Println(person.Greet())
	person.HaveBirthday()
	fmt.Printf("After birthday: %d years old\n\n", person.Age)

	// 3. Interfaces
	fmt.Println("--- 3. Interfaces ---")
	var animal Animal = Dog{Name: "Buddy"}
	fmt.Printf("Dog says: %s\n", animal.Sound())
	fmt.Printf("Dog is: %s\n\n", animal.Move())

	// 4. Arrays
	fmt.Println("--- 4. Arrays ---")
	matrix := Matrix{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Printf("Matrix sum: %d\n", matrix.Sum())

	rgb := RGB{255, 128, 0}
	fmt.Printf("RGB: %v = %s\n", rgb, rgb.ToHex())

	coords := Coordinates{3.0, 4.0}
	fmt.Printf("Coordinates distance: %.2f\n\n", coords.Distance())

	// 5. Slices
	fmt.Println("--- 5. Slices ---")
	nums := IntSlice{1, 2, 3, 4, 5}
	fmt.Printf("IntSlice: %v\n", nums)
	fmt.Printf("Sum: %d, Average: %.2f\n", nums.Sum(), nums.Average())

	names := StringList{"Alice", "Bob", "Charlie"}
	fmt.Printf("Contains 'Bob': %v\n", names.Contains("Bob"))

	scores := ScoreList{85.5, 92.0, 78.5, 95.0}
	fmt.Printf("Max score: %.1f\n\n", scores.Max())

	// 6. Maps
	fmt.Println("--- 6. Maps ---")
	config := StringMap{"host": "localhost", "port": "8080"}
	fmt.Printf("Config keys: %v\n", config.Keys())

	counter := Counter{"apples": 5, "oranges": 3}
	counter.Increment("apples")
	fmt.Printf("Counter total: %d\n\n", counter.Total())

	// 7. Channels
	fmt.Println("--- 7. Channels ---")
	ch := make(IntChannel, 5)
	ch.SendRange(1, 5)
	fmt.Print("Channel values: ")
	for val := range ch {
		fmt.Printf("%d ", val)
	}
	fmt.Println("\n")

	// 8. Functions
	fmt.Println("--- 8. Functions ---")
	var handler Handler = func(s string) error {
		fmt.Printf("Processing: %s\n", s)
		return nil
	}
	loggedHandler := handler.WithLogging()
	loggedHandler("test data")
	fmt.Println()

	// 9. Pointers (using struct wrapper)
	fmt.Println("--- 9. Pointers ---")
	wrapper := IntWrapper{Value: nil}
	fmt.Printf("IntWrapper is nil: %v\n", wrapper.IsNil())

	val := 42
	wrapper.Value = &val
	fmt.Printf("IntWrapper is nil: %v, value: %d\n", wrapper.IsNil(), wrapper.Get())
	
	wrapper.Set(100)
	fmt.Printf("After Set(100): %d\n", wrapper.Get())
}
