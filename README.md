# 🎓 GoLearn101

**A hands-on Go programming learning journey with exercises, quizzes, and mini-projects**

---

## 📖 About

This repository documents my journey learning Go through practical exercises, comprehensive quizzes, and real-world mini-projects. It covers fundamental to intermediate concepts with runnable code examples and test-driven development.

**Level:** Beginner to Intermediate  
**Approach:** Learn by doing - code examples, quizzes, and projects

---

## 🗂️ Repository Structure

```
golearn101/
├── day1-2026/              # Core Go concepts & fundamentals
│   ├── goroutine1.go       # Goroutines and concurrency basics
│   ├── channel1.go         # Channel operations
│   ├── mutex1.go           # Mutex and synchronization
│   ├── struct1.go          # Structs and methods
│   ├── ptr1.go             # Pointers and memory
│   ├── file1.go            # File I/O operations
│   └── ... (30+ examples)
│
├── day2-exercise/          # Practice exercises & quizzes
│   ├── quiz-e1.go          # Basic syntax quiz
│   ├── quiz-e2.go          # Functions & methods quiz
│   ├── unique-e1.go        # Array/slice exercises
│   ├── swap-ptr.go         # Pointer exercises
│   └── ... (15+ exercises)
│
├── basic-projects/         # Mini-projects
│   └── password-gen/       # Password Generator CLI
│    # Best practices
└── progress.txt            # Learning progress tracker
```

---

## 🚀 Quick Start

### Prerequisites
- Go 1.22+ installed
- Basic programming knowledge
- Terminal/command line familiarity

### Getting Started

```bash
# Clone the repository
git clone git@github.com:rajatk10/golearn101.git
cd golearn101

# Run examples
go run day1-2026/goroutine1.go

# Run quizzes
go run day2-exercise/quiz-e6.go

# Try the password generator
cd basic-projects/password-gen
go run . --length 20 --special
```

---

### **Core Concepts (day1-2026/)**

**Data Types & Structures**
- Basic types, strings, arrays, slices, maps
- Structs, methods, interfaces
- Pointers and memory management

**Concurrency**
- Goroutines and the Go scheduler
- Channels (buffered/unbuffered)
- Select statement and timeouts
- Mutexes, WaitGroups, and synchronization

**Error Handling**
- Error interface and custom errors
- Error wrapping and unwrapping
- Panic and recover patterns

**File I/O**
- Reading and writing files
- Buffered I/O with `bufio`
- File operations and error handling

**Testing**
- Unit tests with `testing` package
- Test assertions and error reporting
- Running and organizing tests

---

### **Practice Exercises (day2-exercise/)**

**Comprehensive Quizzes:**
- `quiz-e1.go` - Basic syntax and control flow
- `quiz-e2.go` - Functions, methods, and closures
- `quiz-e3.go` - Interfaces and polymorphism
- `quiz-e4.go` - Concurrency patterns
- `quiz-e5.go` - Error handling strategies
- `quiz-e6.go` - Advanced concurrency (9 questions with detailed answers)

**Coding Exercises:**
- Array/slice manipulation
- Pointer operations
- Struct methods
- File operations
- Duplicate detection algorithms

---

### **Mini-Projects (basic-projects/)**

#### **Password Generator CLI**

**Usages Example:**
```bash
cd basic-projects/password-gen

# Generate default password (16 chars, all types)
go run .

# Custom length and options
go run . --length 20 --special=false

# Run tests
go test -v
go test -cover
```

---

## Key Learning Outcomes

 **Concurrency Mastery**
- Understanding goroutines vs threads
- Channel communication patterns

**Error Handling**
- Idiomatic error patterns
- Custom error types

**Testing Skills**
- Writing unit tests
- Test case design

**CLI Development**
- Flag parsing
- User input validation

**Web development**
- HTTP Requests and Methods
---

## 🛠️ Running Examples

```bash
# Run any example
go run day1-2026/<filename>.go

# Run with race detector
go run -race day1-2026/goroutine1.go

# Run tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Format code
go fmt ./...
gofmt -w .
```

---

## 📖 Key Resources

- **Go By Example** - https://gobyexample.com/
- **Go Tour** - https://go.dev/doc/
---

## 💡 Learning Approach

**1. Understand Concepts**
- Read examples in `day1-2026/`
- Study quiz questions and answers

**2. Practice**
- Complete exercises in `day2-exercise/`
- Modify examples to experiment

**3. Build Projects**
- Apply concepts in mini-projects
- Write tests for validation

**4. Review & Iterate**
- Run tests to verify understanding
- Refactor and improve code

---
## 📄 License

MIT License - Free to use for learning purposes.

---

**Learning Go, one project at a time! 🚀**
