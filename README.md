# 🎓 GoLearn101

**A comprehensive 7-day intermediate Go programming learning journey**

---

## 📖 About

This repository contains hands-on exercises, quizzes, and practical projects for mastering intermediate Go concepts. Designed for developers with basic Go knowledge who want to level up their skills through structured, daily practice.

**Learning Duration:** 7 days × 3 hours/day  
**Level:** Intermediate  
**Approach:** Interactive quizzes, code examples, and real-world projects

---

## 🗂️ Repository Structure

```
golearn101/
├── day1-2026/          # Day 1: Goroutines, Channels, Concurrency
│   ├── goroutine1.go   # Goroutine examples
│   ├── mutex1.go       # Mutex and synchronization
│   └── ...
├── day2-exercise/      # Day 2: Error Handling, File I/O, Testing
│   ├── quiz6.go        # Concurrency quiz with answers
│   ├── error-file.go   # Error handling examples
│   └── ...
├── 7_DAY_LEARNING_PLAN.md    # Complete learning roadmap
├── START_HERE.md             # Getting started guide
└── GO_LEARNING_RULES.md      # Best practices and guidelines
```

---

## 🚀 Quick Start

### Prerequisites
- Go 1.22+ installed
- Basic Go knowledge (syntax, functions, structs)
- 3 hours daily commitment

### Getting Started

1. **Clone the repository:**
   ```bash
   git clone git@github.com:yourusername/golearn101.git
   cd golearn101
   ```

2. **Read the learning plan:**
   ```bash
   cat 7_DAY_LEARNING_PLAN.md
   ```

3. **Start with Day 1:**
   ```bash
   cd day1-2026
   go run goroutine1.go
   ```

---

## 📚 Learning Path

### **Day 1:** Advanced Data Structures & Methods
- Structs, methods, interfaces
- Pointer vs value receivers
- Type embedding and composition

### **Day 2:** Error Handling & Package Design
- Custom errors and error wrapping
- File I/O with bufio
- Unit testing with `testing` package

### **Day 3:** Goroutines & Concurrency
- Goroutines and channels
- Select statement and timeouts
- Mutexes and synchronization
- **Quiz 6:** Comprehensive concurrency quiz

### **Day 4:** Advanced Concurrency Patterns
- Worker pools and pipelines
- Context package
- Race condition detection

### **Day 5:** Web Development Basics
- HTTP servers and routing
- JSON encoding/decoding
- Middleware patterns

### **Day 6:** Database & Persistence
- SQL with `database/sql`
- Connection pooling
- Migrations and transactions

### **Day 7:** Testing & Best Practices
- Table-driven tests
- Benchmarking
- Code organization

---

## 🎯 Key Features

✅ **Interactive Quizzes** - Test your understanding with "What, Why, How" questions  
✅ **Runnable Examples** - Every concept includes working code  
✅ **Progressive Learning** - Build on previous day's knowledge  
✅ **Real-world Projects** - Practical applications of concepts  
✅ **Best Practices** - Learn idiomatic Go patterns  

---

## 📝 Quiz System

Each topic includes comprehensive quizzes with:
- **What questions:** Understanding core concepts
- **Why questions:** Reasoning behind design decisions
- **How questions:** Practical implementation details

**Example:** `day2-exercise/quiz6.go` - Complete concurrency quiz with 9 questions and answers

---

## 🛠️ Running Examples

```bash
# Run individual examples
go run day1-2026/goroutine1.go
go run day1-2026/mutex1.go

# Run quiz files
go run day2-exercise/quiz6.go

# Run tests
go test ./...
```

---

## 📖 Key Topics Covered

### Concurrency
- Goroutines vs OS threads
- Channels (buffered/unbuffered)
- Select statement
- Mutexes and WaitGroups
- G-M-P scheduling model

### Error Handling
- Custom error types
- Error wrapping with `fmt.Errorf`
- Interface satisfaction for errors
- Panic and recover

### File I/O
- Buffer-based reading/writing
- `bufio` package for efficiency
- File operations and cleanup

### Testing
- Unit tests with `*testing.T`
- Table-driven tests
- Subtests and parallel execution

---

## 🎓 Learning Resources

- **7_DAY_LEARNING_PLAN.md** - Detailed daily curriculum
- **START_HERE.md** - Setup and orientation guide
- **GO_LEARNING_RULES.md** - Coding standards and best practices
- **progress.txt** - Track your learning progress

---

## 🤝 Contributing

This is a personal learning repository. Feel free to fork and adapt for your own learning journey!

---

## 📄 License

MIT License - Feel free to use this for your own learning.

---

## 🌟 Progress Tracking

Track your progress in `progress.txt` or create your own learning journal.

**Current Status:** Day 2-3 (Concurrency & Error Handling)

---

## 💡 Tips for Success

1. **Code along** - Don't just read, type and run the examples
2. **Complete quizzes** - Test your understanding before moving forward
3. **Experiment** - Modify examples to see what happens
4. **Build projects** - Apply concepts to real problems
5. **Review regularly** - Revisit previous days' concepts

---

**Happy Learning! 🚀**

*Master Go one day at a time.*
