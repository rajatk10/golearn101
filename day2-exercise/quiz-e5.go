package main

/*
GO INTERFACES QUIZ - Questions & Answers

WHAT Questions:
================

Q1: What is an interface in Go, and how is it different from a struct?
A: Interface:
   - Defines a CONTRACT (set of method signatures)
   - Specifies BEHAVIOR (what methods must exist)
   - No data fields, only method declarations
   - Implemented IMPLICITLY (no "implements" keyword)
   - Used for POLYMORPHISM and ABSTRACTION
   - Separates "what" from "how"

   Struct:
   - Defines DATA STRUCTURE (fields)
   - Contains STATE (actual data)
   - Can have methods attached to it
   - Concrete type with memory layout
   - Used for DATA MODELING

   Example:
   // Interface - defines WHAT (behavior)
   type Speaker interface {
       Speak() string  // Contract: must have Speak method
   }

   // Struct - defines DATA
   type Dog struct {
       Name string  // Has data fields
   }

   // Implementation - defines HOW
   func (d Dog) Speak() string {
       return "Woof!"  // Dog implements Speaker implicitly
   }

   var s Speaker = Dog{Name: "Buddy"}  // Works!

   Key Difference:
   ┌──────────────┬─────────────┬────────────┐
   │ Feature      │ Interface   │ Struct     │
   ├──────────────┼─────────────┼────────────┤
   │ Purpose      │ Behavior    │ Data       │
   │ Contains     │ Methods     │ Fields     │
   │ Implementation│ Implicit   │ Explicit   │
   │ Use Case     │ Abstraction │ Modeling   │
   └──────────────┴─────────────┴────────────┘

Q2: What is the interface{} type, and why is it called the empty interface?
A: Empty Interface (interface{} or 'any' in Go 1.18+):

   Definition:
   interface{}  // Has ZERO methods

   Why "empty"?
   - Contains NO method requirements
   - Every type implements zero methods
   - Therefore, EVERY TYPE satisfies it
   - Can hold ANY VALUE (int, string, struct, etc.)

   Modern Go (1.18+):
   any  // Alias for interface{} - cleaner syntax

   Examples:
   var x interface{}
   x = 42           // int
   x = "hello"      // string
   x = []int{1,2,3} // slice
   x = Dog{Name: "Buddy"}  // struct
   // All valid! interface{} accepts anything

   Common Use Cases:
   1. Generic containers (before Go 1.18 generics):
      func Print(v interface{}) {
          fmt.Println(v)  // Works with any type
      }

   2. JSON unmarshaling (unknown structure):
      var data interface{}
      json.Unmarshal(jsonBytes, &data)

   3. Heterogeneous collections:
      mixed := []interface{}{42, "hello", true}

   Why Use Sparingly?
   - Loses TYPE SAFETY (compile-time checks gone)
   - Requires TYPE ASSERTIONS (runtime overhead)
   - Makes code HARDER TO UNDERSTAND
   - Defeats Go's strong typing

   Better Alternatives:
   - Use SPECIFIC INTERFACES when possible
   - Use GENERICS (Go 1.18+) for type-safe generic code

Q3: What are some real-world scenarios where interfaces are used effectively in Go?
A: Real-World Interface Use Cases:

   1. I/O Operations (io.Reader, io.Writer):
      type Reader interface {
          Read(p []byte) (n int, err error)
      }

      // Works with files, network, memory, etc.
      func ProcessData(r io.Reader) {
          // Same code works for any Reader!
      }

      ProcessData(file)      // File
      ProcessData(httpResp)  // HTTP response
      ProcessData(buffer)    // Memory buffer

   2. Database Abstraction:
      type Database interface {
          Query(sql string) ([]Row, error)
          Insert(table string, data map[string]interface{}) error
      }

      // Switch between MySQL, PostgreSQL, MongoDB
      func SaveUser(db Database, user User) {
          db.Insert("users", user.ToMap())
      }

   3. Testing & Mocking:
      type EmailSender interface {
          Send(to, subject, body string) error
      }

      // Production: real email
      type SMTPSender struct{}
      func (s SMTPSender) Send(to, subject, body string) error {
          // Actually send email
      }

      // Testing: mock email
      type MockSender struct{}
      func (m MockSender) Send(to, subject, body string) error {
          // Just log, don't send
          return nil
      }

   4. Payment Processing:
      type PaymentGateway interface {
          ProcessPayment(amount float64) error
          Refund(transactionID string) error
      }

      // Stripe, PayPal, Square - all implement same interface
      func Checkout(gateway PaymentGateway, amount float64) {
          gateway.ProcessPayment(amount)
      }

   5. Logging:
      type Logger interface {
          Log(message string)
      }

      // Console, file, cloud - same interface
      func ProcessRequest(logger Logger) {
          logger.Log("Request started")
          // ... processing
          logger.Log("Request completed")
      }

   6. Vehicle Example (Polymorphism):
      type Vehicle interface {
          Start() error
          Move() error
          Stop() error
      }

      type ElectricCar struct{}
      func (e ElectricCar) Start() error { }

      type GasCar struct{}
      func (g GasCar) Start() error { }

      // Same function works for any vehicle
      func Drive(v Vehicle) {
          v.Start()
          v.Move()
          v.Stop()
      }

   Benefits:
   - FLEXIBILITY: Swap implementations easily
   - TESTABILITY: Mock dependencies
   - MAINTAINABILITY: Change internals without affecting callers
   - DECOUPLING: Reduce dependencies between components


WHY Questions:
==============

Q1: Why is polymorphism important in Go, and how do interfaces enable it?
A: Polymorphism Importance:

   1. Write Once, Use Everywhere:
      type Shape interface {
          Area() float64
      }

      // One function works for ALL shapes
      func PrintArea(s Shape) {
          fmt.Printf("Area: %.2f\n", s.Area())
      }

      // Different implementations
      type Circle struct { Radius float64 }
      func (c Circle) Area() float64 { return 3.14 * c.Radius * c.Radius }

      type Rectangle struct { Width, Height float64 }
      func (r Rectangle) Area() float64 { return r.Width * r.Height }

      // Same function, different behaviors
      PrintArea(Circle{Radius: 5})      // Uses Circle's Area()
      PrintArea(Rectangle{Width: 4, Height: 3})  // Uses Rectangle's Area()

   2. Extensibility Without Modification:
      // Add new shape WITHOUT changing PrintArea function
      type Triangle struct { Base, Height float64 }
      func (t Triangle) Area() float64 { return 0.5 * t.Base * t.Height }

      PrintArea(Triangle{Base: 6, Height: 4})  // Works immediately!

   3. Decoupling & Flexibility:
      type Storage interface {
          Save(data string) error
      }

      func BackupData(storage Storage, data string) {
          storage.Save(data)  // Don't care HOW it saves
      }

      // Can swap implementations easily
      BackupData(FileStorage{}, data)   // Save to file
      BackupData(CloudStorage{}, data)  // Save to cloud
      BackupData(DatabaseStorage{}, data)  // Save to DB

   4. Testing & Mocking:
      type EmailService interface {
          SendEmail(to, msg string) error
      }

      // Production
      type RealEmailService struct{}
      func (r RealEmailService) SendEmail(to, msg string) error {
          // Actually send email
      }

      // Testing
      type MockEmailService struct{}
      func (m MockEmailService) SendEmail(to, msg string) error {
          // Just log, don't send
          return nil
      }

   How Interfaces Enable It:
   - IMPLICIT IMPLEMENTATION: No "implements" keyword needed
   - DUCK TYPING: "If it walks like a duck, it's a duck"
   - BEHAVIOR-BASED: Focus on what it does, not what it is
   - LOOSE COUPLING: Depend on abstractions, not concrete types

   Benefits:
   ✅ Code reusability
   ✅ Easy to extend
   ✅ Easy to test
   ✅ Reduced dependencies
   ✅ Flexible architecture

Q2: Why does Go not support explicit implementation of interfaces (like in Java or C#)?
A: Implicit Implementation Benefits:

   Java/C# (Explicit):
   class Dog implements Animal {  // Must declare "implements"
       public void speak() { }
   }

   Go (Implicit):
   type Dog struct{}
   func (d Dog) Speak() string { return "Woof" }
   // Automatically satisfies any interface with Speak()

   Reasons for Implicit Implementation:

   1. DECOUPLING:
      // Package A defines interface
      package logger
      type Logger interface {
          Log(msg string)
      }

      // Package B creates type (doesn't know about Logger)
      package myapp
      type FileWriter struct{}
      func (f FileWriter) Log(msg string) { }

      // Later: FileWriter satisfies Logger automatically!
      // No need to modify Package B

   2. RETROACTIVE INTERFACE SATISFACTION:
      // You can create interfaces for existing types
      type Stringer interface {
          String() string
      }

      // Even standard library types satisfy it if they have String()
      var s Stringer = time.Now()  // time.Time has String() method

   3. FLEXIBILITY:
      // Same type can satisfy multiple interfaces without declaring
      type File struct{}
      func (f File) Read(p []byte) (int, error) { }
      func (f File) Write(p []byte) (int, error) { }
      func (f File) Close() error { }

      // Automatically satisfies:
      var r io.Reader = File{}
      var w io.Writer = File{}
      var c io.Closer = File{}
      var rw io.ReadWriter = File{}

   4. NO DEPENDENCY ON INTERFACE DEFINITION:
      // Type doesn't need to import interface package
      // Reduces circular dependencies
      // Promotes loose coupling

   5. INTERFACE SEGREGATION:
      // Can define small, focused interfaces
      // Types automatically satisfy them if methods match
      type Reader interface { Read([]byte) (int, error) }
      type Writer interface { Write([]byte) (int, error) }
      type ReadWriter interface { Reader; Writer }

   Benefits:
   ✅ Loose coupling between packages
   ✅ No circular dependencies
   ✅ Retroactive interface satisfaction
   ✅ Flexible and composable
   ✅ Simpler code (no "implements" boilerplate)

Q3: Why should you avoid overusing the empty interface in Go?
A: Problems with Overusing interface{}:

   1. LOSES TYPE SAFETY:
      // ❌ Bad: No compile-time checks
      func Add(a, b interface{}) interface{} {
          return a.(int) + b.(int)  // Runtime panic if not int!
      }

      // ✅ Good: Type-safe
      func Add(a, b int) int {
          return a + b  // Compiler catches errors
      }

   2. REQUIRES TYPE ASSERTIONS (Runtime Overhead):
      // ❌ Bad: Manual type checking everywhere
      func Process(data interface{}) {
          if str, ok := data.(string); ok {
              // handle string
          } else if num, ok := data.(int); ok {
              // handle int
          }
          // Tedious and error-prone!
      }

      // ✅ Good: Clear types
      func ProcessString(data string) { }
      func ProcessInt(data int) { }

   3. POOR DOCUMENTATION:
      // ❌ Bad: What does this accept?
      func Save(data interface{}) error {
          // ???
      }

      // ✅ Good: Clear expectations
      func SaveUser(user User) error {
          // Obviously expects User
      }

   4. HARDER TO MAINTAIN:
      // ❌ Bad: What can items contain?
      items := []interface{}{1, "hello", true, 3.14}
      // Need to check each element's type

      // ✅ Good: Homogeneous and clear
      users := []User{{Name: "Alice"}, {Name: "Bob"}}

   5. DEFEATS GO'S PHILOSOPHY:
      - Go emphasizes EXPLICIT over implicit
      - Strong typing catches bugs EARLY
      - Empty interface bypasses all type checks

   When to Use interface{}/any:
   ✅ JSON unmarshaling (unknown structure)
   ✅ Generic containers (before Go 1.18)
   ✅ Reflection-based libraries
   ✅ Printf-style functions

   Better Alternatives:
   - Use SPECIFIC INTERFACES when possible
   - Use GENERICS (Go 1.18+) for type-safe generic code
   - Use UNION TYPES with type switches when needed

   Example with Generics (Go 1.18+):
   // ✅ Type-safe generic function
   func Max[T comparable](a, b T) T {
       if a > b {
           return a
       }
       return b
   }

   Max(5, 10)      // Works with int
   Max(3.14, 2.71) // Works with float64


HOW Questions:
==============

Q1: How can you use interfaces to write flexible and testable code in Go?
A: Dependency Injection with Interfaces:

   1. Define Interfaces for Dependencies:
      type Database interface {
          GetUser(id int) (User, error)
          SaveUser(user User) error
      }

      type EmailService interface {
          Send(to, subject, body string) error
      }

   2. Service Depends on Interfaces:
      type UserService struct {
          db    Database
          email EmailService
      }

      func NewUserService(db Database, email EmailService) *UserService {
          return &UserService{db: db, email: email}
      }

      func (s *UserService) RegisterUser(user User) error {
          if err := s.db.SaveUser(user); err != nil {
              return err
          }
          return s.email.Send(user.Email, "Welcome", "Thanks for joining!")
      }

   3. Production Implementation:
      type PostgresDB struct {  real DB connection }
      func (p *PostgresDB) GetUser(id int) (User, error) { actual query }
      func (p *PostgresDB) SaveUser(user User) error { actual insert}

      type SMTPEmail struct {  real SMTP config }
      func (s *SMTPEmail) Send(to, subject, body string) error { actually send }

      // Production use
      service := NewUserService(&PostgresDB{}, &SMTPEmail{})

   4. Test Implementation (Mocks):
      type MockDB struct {
          users map[int]User
      }
      func (m *MockDB) GetUser(id int) (User, error) {
          return m.users[id], nil
      }
      func (m *MockDB) SaveUser(user User) error {
          m.users[user.ID] = user
          return nil
      }

      type MockEmail struct {
          sentEmails []string
      }
      func (m *MockEmail) Send(to, subject, body string) error {
          m.sentEmails = append(m.sentEmails, to)
          return nil  // Don't actually send
      }

      // Test use
      func TestRegisterUser(t *testing.T) {
          mockDB := &MockDB{users: make(map[int]User)}
          mockEmail := &MockEmail{}
          service := NewUserService(mockDB, mockEmail)

          user := User{ID: 1, Email: "test@example.com"}
          err := service.RegisterUser(user)

          if err != nil {
              t.Errorf("Expected no error, got %v", err)
          }
          if len(mockEmail.sentEmails) != 1 {
              t.Errorf("Expected 1 email sent")
          }
      }

   5. Easy to Swap Implementations:
      // Switch to different database without changing UserService
      service := NewUserService(&MongoDBDB{}, &SMTPEmail{})

      // Or use different email provider
      service := NewUserService(&PostgresDB{}, &SendGridEmail{})

   Benefits:
   ✅ TESTABLE: Mock dependencies easily
   ✅ FLEXIBLE: Swap implementations without code changes
   ✅ MAINTAINABLE: Changes isolated to implementations
   ✅ DECOUPLED: Components don't depend on concrete types

   Best Practices:
   - Accept interfaces, return structs
   - Keep interfaces small (1-3 methods)
   - Define interfaces where they're used, not where they're implemented

Q2: How does the Go runtime internally represent an interface?
A: Interface Internal Representation:

   Interface Structure (2 pointers):
   ┌─────────────────────┐
   │ Interface Value     │
   ├─────────────────────┤
   │ 1. Type Info (tab)  │ → Points to type metadata
   │ 2. Data Pointer     │ → Points to actual value
   └─────────────────────┘

   1. Type Information (itab):
      - Pointer to TYPE METADATA
      - Contains METHOD TABLE (vtable)
      - Stores CONCRETE TYPE information
      - Used for METHOD DISPATCH

   2. Data Pointer:
      - Points to the ACTUAL VALUE
      - Stores the CONCRETE DATA

   Example:
   type Speaker interface {
       Speak() string
   }

   type Dog struct {
       Name string
   }

   func (d Dog) Speak() string {
       return "Woof"
   }

   var s Speaker = Dog{Name: "Buddy"}

   Internal Representation:
   s (interface value)
   ├── Type Info → Points to Dog's type info + Speak() method
   └── Data      → Points to Dog{Name: "Buddy"}

   Empty Interface (interface{}):
   ┌─────────────────────┐
   │ Empty Interface     │
   ├─────────────────────┤
   │ 1. Type Info (_type)│ → Just type, no methods
   │ 2. Data Pointer     │ → Points to value
   └─────────────────────┘

   Key Points:
   - Interface = (type, value) PAIR
   - NIL INTERFACE: Both pointers are nil
   - INTERFACE WITH NIL VALUE: Type pointer set, data pointer nil
   - SIZE: 16 bytes on 64-bit systems (2 pointers × 8 bytes)

   Checking nil:
   var s Speaker          // nil interface (type=nil, data=nil)
   var s Speaker = (*Dog)(nil)  // non-nil interface (type=Dog, data=nil)

   s == nil  // true for first, false for second!

   Performance:
   - Method calls through interfaces are slightly slower (indirect call)
   - Interface assignment involves copying type info and data pointer
   - Empty interface requires type assertion for any operation

Q3: How can you safely extract the underlying value of an empty interface using type assertions or type switches?
A: Type Assertions and Type Switches:

   1. TYPE ASSERTION (Single Type):

   Unsafe (can panic):
   var data interface{} = "hello"
   str := data.(string)  // Works, str = "hello"

   var num interface{} = 42
   str := num.(string)   // ❌ PANIC! num is not a string

   Safe (with ok check):
   var data interface{} = "hello"

   // Two-value form: value, ok
   str, ok := data.(string)
   if ok {
       fmt.Println("It's a string:", str)
   } else {
       fmt.Println("Not a string")
   }

   2. TYPE SWITCH (Multiple Types):

   func Describe(data interface{}) {
       switch v := data.(type) {
       case int:
           fmt.Printf("Integer: %d\n", v)
       case string:
           fmt.Printf("String: %s\n", v)
       case bool:
           fmt.Printf("Boolean: %v\n", v)
       case []int:
           fmt.Printf("Int slice: %v\n", v)
       case nil:
           fmt.Println("Nil value")
       default:
           fmt.Printf("Unknown type: %T\n", v)
       }
   }

   Describe(42)        // Integer: 42
   Describe("hello")   // String: hello
   Describe(true)      // Boolean: true
   Describe(3.14)      // Unknown type: float64

   3. COMPLETE EXAMPLE:

   func ProcessValue(data interface{}) {
       // Method 1: Type assertion with ok check
       if str, ok := data.(string); ok {
           fmt.Println("String:", str)
           return
       }

       if num, ok := data.(int); ok {
           fmt.Println("Int:", num)
           return
       }

       // Method 2: Type switch (cleaner for multiple types)
       switch v := data.(type) {
       case string:
           fmt.Println("String:", v)
       case int:
           fmt.Println("Int:", v)
       case float64:
           fmt.Println("Float:", v)
       case []int:
           fmt.Println("Int slice:", v)
       case nil:
           fmt.Println("Nil value")
       default:
           fmt.Printf("Unknown type: %T\n", v)
       }
   }

   Best Practices:
   ✅ Always use COMMA-OK idiom for safety
   ✅ Use TYPE SWITCH for multiple types
   ❌ Avoid single-value assertion (can panic)

   Key Syntax:
   // Type assertion
   value, ok := interfaceVar.(ConcreteType)

   // Type switch
   switch v := interfaceVar.(type) {
   case Type1:
       // v is Type1
   case Type2:
       // v is Type2
   default:
       // Unknown type
   }

Q4: How would you implement a custom error type using the error interface?
A: Custom Error Implementation:

   The error interface is simple:
   type error interface {
       Error() string  // Just one method!
   }

   Any type with Error() string method implements error.

   1. SIMPLE CUSTOM ERROR:

   type ValidationError struct {
       Field   string
       Message string
   }

   func (e ValidationError) Error() string {
       return fmt.Sprintf("validation error on field '%s': %s", e.Field, e.Message)
   }

   // Usage
   func ValidateAge(age int) error {
       if age < 0 {
           return ValidationError{
               Field:   "age",
               Message: "must be positive",
           }
       }
       return nil
   }

   2. ERROR WITH ADDITIONAL DATA:

   type DatabaseError struct {
       Operation string
       Table     string
       Err       error  // Wrap underlying error
   }

   func (e DatabaseError) Error() string {
       return fmt.Sprintf("database error during %s on table %s: %v",
           e.Operation, e.Table, e.Err)
   }

   // Usage
   func SaveUser(user User) error {
       err := db.Insert("users", user)
       if err != nil {
           return DatabaseError{
               Operation: "insert",
               Table:     "users",
               Err:       err,
           }
       }
       return nil
   }

   3. ERROR WITH STATUS CODE:

   type HTTPError struct {
       StatusCode int
       Message    string
   }

   func (e HTTPError) Error() string {
       return fmt.Sprintf("HTTP %d: %s", e.StatusCode, e.Message)
   }

   // Usage
   func GetUser(id int) (User, error) {
       if id < 0 {
           return User{}, HTTPError{
               StatusCode: 400,
               Message:    "invalid user ID",
           }
       }
       // ...
   }

   4. TYPE ASSERTION TO ACCESS FIELDS:

   func HandleError(err error) {
       if err == nil {
           return
       }

       // Type assertion to access custom fields
       if httpErr, ok := err.(HTTPError); ok {
           fmt.Printf("HTTP Error: Status %d\n", httpErr.StatusCode)
           return
       }

       if dbErr, ok := err.(DatabaseError); ok {
           fmt.Printf("DB Error: Operation %s, Table %s\n",
               dbErr.Operation, dbErr.Table)
           return
       }

       fmt.Println("Unknown error:", err)
   }

   5. COMPLETE WORKING EXAMPLE:

   type AppError struct {
       Code    int
       Message string
       Details string
   }

   func (e AppError) Error() string {
       return fmt.Sprintf("[Error %d] %s: %s", e.Code, e.Message, e.Details)
   }

   func ProcessData(data string) error {
       if data == "" {
           return AppError{
               Code:    1001,
               Message: "Invalid input",
               Details: "data cannot be empty",
           }
       }
       return nil
   }

   func main() {
       err := ProcessData("")
       if err != nil {
           fmt.Println(err)  // [Error 1001] Invalid input: data cannot be empty

           // Access custom fields
           if appErr, ok := err.(AppError); ok {
               fmt.Printf("Error code: %d\n", appErr.Code)
           }
       }
   }

   Key Points:
   ✅ Implement Error() string method
   ✅ Can add any fields you need
   ✅ Use type assertion to access custom fields
   ✅ Can wrap other errors
   ✅ Provides rich error context


KEY TAKEAWAYS:
==============
- Interfaces define behavior contracts (what), not implementation (how)
- Implicit implementation: no "implements" keyword needed
- Empty interface (interface{}/any) accepts any type but loses type safety
- Use interfaces for polymorphism, abstraction, and testability
- Real-world uses: I/O, databases, testing, payment gateways, logging
- Avoid overusing empty interface - prefer specific interfaces or generics
- Interface internally: (type info, data pointer) pair
- Type assertions and type switches extract underlying values safely
- Custom errors: implement Error() string method
- Best practices: small interfaces, dependency injection, accept interfaces/return structs

// Example types for demonstration
type User struct {
	ID    int
	Name  string
	Email string
}

// Example 1: Interface vs Struct
type Speaker interface {
	Speak() string
}

type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return fmt.Sprintf("%s says: Woof!", d.Name)
}

type Cat struct {
	Name string
}

func (c Cat) Speak() string {
	return fmt.Sprintf("%s says: Meow!", c.Name)
}

// Example 2: Empty Interface
func PrintAnything(v interface{}) {
	fmt.Printf("Value: %v, Type: %T\n", v, v)
}

// Example 3: Type Assertion and Type Switch
func Describe(data interface{}) {
	switch v := data.(type) {
	case int:
		fmt.Printf("Integer: %d\n", v)
	case string:
		fmt.Printf("String: %s\n", v)
	case bool:
		fmt.Printf("Boolean: %v\n", v)
	case []int:
		fmt.Printf("Int slice: %v\n", v)
	case nil:
		fmt.Println("Nil value")
	default:
		fmt.Printf("Unknown type: %T\n", v)
	}
}

// Example 4: Custom Error Type
type AppError struct {
	Code    int
	Message string
	Details string
}

func (e AppError) Error() string {
	return fmt.Sprintf("[Error %d] %s: %s", e.Code, e.Message, e.Details)
}

func ValidateUser(user User) error {
	if user.Name == "" {
		return AppError{
			Code:    1001,
			Message: "Validation failed",
			Details: "name cannot be empty",
		}
	}
	if user.Email == "" {
		return AppError{
			Code:    1002,
			Message: "Validation failed",
			Details: "email cannot be empty",
		}
	}
	return nil
}

// Example 5: Dependency Injection for Testing
type Database interface {
	SaveUser(user User) error
	GetUser(id int) (User, error)
}

type EmailService interface {
	Send(to, subject, body string) error
}

type UserService struct {
	db    Database
	email EmailService
}

func NewUserService(db Database, email EmailService) *UserService {
	return &UserService{db: db, email: email}
}

func (s *UserService) RegisterUser(user User) error {
	if err := s.db.SaveUser(user); err != nil {
		return err
	}
	return s.email.Send(user.Email, "Welcome", "Thanks for joining!")
}

// Mock implementations for testing
type MockDB struct {
	users map[int]User
}

func (m *MockDB) SaveUser(user User) error {
	if m.users == nil {
		m.users = make(map[int]User)
	}
	m.users[user.ID] = user
	return nil
}

func (m *MockDB) GetUser(id int) (User, error) {
	return m.users[id], nil
}

type MockEmail struct {
	sentEmails []string
}

func (m *MockEmail) Send(to, subject, body string) error {
	m.sentEmails = append(m.sentEmails, to)
	fmt.Printf("Mock: Email sent to %s\n", to)
	return nil
}

func main() {
	fmt.Println("=== GO QUIZ 5: Interfaces ===")
	fmt.Println("See comments above for all questions and answers!")
	fmt.Println()

	// Example 1: Polymorphism
	fmt.Println("--- 1. Polymorphism with Interfaces ---")
	var speaker Speaker
	speaker = Dog{Name: "Buddy"}
	fmt.Println(speaker.Speak())
	speaker = Cat{Name: "Whiskers"}
	fmt.Println(speaker.Speak())
	fmt.Println()

	// Example 2: Empty Interface
	fmt.Println("--- 2. Empty Interface ---")
	PrintAnything(42)
	PrintAnything("hello")
	PrintAnything(true)
	PrintAnything([]int{1, 2, 3})
	fmt.Println()

	// Example 3: Type Switch
	fmt.Println("--- 3. Type Switch ---")
	Describe(42)
	Describe("hello")
	Describe(true)
	Describe([]int{1, 2, 3})
	Describe(nil)
	fmt.Println()

	// Example 4: Custom Error
	fmt.Println("--- 4. Custom Error Type ---")
	user := User{ID: 1, Name: "", Email: "test@example.com"}
	err := ValidateUser(user)
	if err != nil {
		fmt.Println(err)
		if appErr, ok := err.(AppError); ok {
			fmt.Printf("Error code: %d\n", appErr.Code)
		}
	}
	fmt.Println()

	// Example 5: Dependency Injection
	fmt.Println("--- 5. Dependency Injection & Testing ---")
	mockDB := &MockDB{}
	mockEmail := &MockEmail{}
	service := NewUserService(mockDB, mockEmail)

	user = User{ID: 1, Name: "Alice", Email: "alice@example.com"}
	err = service.RegisterUser(user)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("User registered successfully!")
		fmt.Printf("Emails sent: %d\n", len(mockEmail.sentEmails))
	}
}
*/
