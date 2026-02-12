package main

/*
GO CONTROL FLOW QUIZ - Questions & Answers

WHAT Questions:
================

Q1: What is the purpose of the if statement in Go?
A: Purpose of if statement:
   - Conditional execution based on boolean expression
   - Controls program flow by branching logic
   - Executes code block only when condition is true
   
   Syntax:
   if condition {
       // Execute if true
   }
   
   if condition {
       // Execute if true
   } else {
       // Execute if false
   }
   
   if condition1 {
       // Execute if condition1 true
   } else if condition2 {
       // Execute if condition2 true
   } else {
       // Execute if all false
   }
   
   Go-specific feature - Short statement:
   if x := getValue(); x > 10 {
       // x is scoped to this if block
   }
   // x not accessible here
   
   Key Points:
   - Condition must be boolean (no implicit conversion)
   - Braces {} are mandatory
   - Can initialize variable in if statement

Q2: What is the significance of the defer keyword in Go, and how does it manage execution flow?
A: Significance of defer:
   - Defers function execution until surrounding function returns
   - Executes in LIFO order (Last In, First Out)
   - Runs even if function panics (before panic propagates)
   - Arguments evaluated immediately, but execution delayed
   
   How it manages execution flow:
   
   1. Execution timing:
      func example() {
          defer fmt.Println("Last")
          fmt.Println("First")
          fmt.Println("Second")
      }
      // Output: First, Second, Last
   
   2. LIFO order (stack):
      defer fmt.Println("1")
      defer fmt.Println("2")
      defer fmt.Println("3")
      // Output: 3, 2, 1
   
   3. Guaranteed cleanup:
      file, _ := os.Open("file.txt")
      defer file.Close()  // Always closes, even if panic
      // Use file...
   
   4. Works with panic:
      defer func() {
          if r := recover(); r != nil {
              fmt.Println("Recovered:", r)
          }
      }()
      panic("error")  // Defer still executes
   
   Key Points:
   - Defers execute in reverse order (LIFO)
   - Perfect for cleanup (close files, unlock mutexes)
   - Executes even during panic
   - Arguments evaluated immediately

Q3: What is a switch statement, and how does it differ from if-else constructs in Go?
A: What is switch:
   - Multi-way conditional statement
   - Compares expression against multiple values
   - Cleaner alternative to multiple if-else chains
   
   Syntax:
   switch expression {
   case value1:
       // Execute if expression == value1
   case value2:
       // Execute if expression == value2
   default:
       // Execute if no match
   }
   
   Differences from if-else:
   ┌──────────────────┬─────────────────┬──────────────────────┐
   │ Feature          │ switch          │ if-else              │
   ├──────────────────┼─────────────────┼──────────────────────┤
   │ Readability      │ Cleaner         │ Verbose with many    │
   │ Break            │ Auto-breaks     │ N/A                  │
   │ Multiple values  │ case 1, 2, 3:   │ Needs || operators   │
   │ No expression    │ switch { }      │ Standard usage       │
   │ Type switch      │ switch v.(type) │ Not possible         │
   └──────────────────┴─────────────────┴──────────────────────┘
   
   Examples:
   
   1. Basic switch:
      switch day {
      case 1:
          fmt.Println("Monday")
      case 2:
          fmt.Println("Tuesday")
      default:
          fmt.Println("Other")
      }
   
   2. Multiple values per case:
      switch month {
      case 1, 3, 5, 7, 8, 10, 12:
          fmt.Println("31 days")
      case 4, 6, 9, 11:
          fmt.Println("30 days")
      }
   
   3. Tagless switch (like if-else):
      switch {
      case age < 18:
          fmt.Println("Minor")
      case age < 65:
          fmt.Println("Adult")
      default:
          fmt.Println("Senior")
      }
   
   Key Points:
   - No break needed (auto-breaks)
   - Can have multiple values per case
   - Cleaner than long if-else chains

Q4: What is the only looping construct in Go, and how can it replace while loops?
A: The only loop: for
   
   Go has only one looping keyword: for, but it's versatile!
   
   1. Traditional for loop:
      for i := 0; i < 5; i++ {
          fmt.Println(i)
      }
   
   2. While-style loop (condition only):
      i := 0
      for i < 5 {
          fmt.Println(i)
          i++
      }
   
   3. Infinite loop:
      for {
          // Runs forever
          if condition {
              break  // Exit loop
          }
      }
   
   4. Range loop (iterate collections):
      nums := []int{1, 2, 3}
      for index, value := range nums {
          fmt.Println(index, value)
      }
   
   How it replaces while:
   ┌─────────────────────┬──────────────────────┐
   │ Other Languages     │ Go Equivalent        │
   ├─────────────────────┼──────────────────────┤
   │ while (condition)   │ for condition { }    │
   │ do-while            │ for { ... break }    │
   │ for (;;)            │ for { }              │
   └─────────────────────┴──────────────────────┘
   
   Loop control:
   - break - exit loop
   - continue - skip to next iteration
   
   Key Point: One keyword (for), multiple styles!

Q5: What does the range keyword do in a loop, and what types does it support?
A: What range does:
   - Iterates over elements in a collection
   - Returns index/key and value for each element
   - Simplifies iteration (no manual index management)
   
   Syntax:
   for index, value := range collection {
       // Use index and value
   }
   
   Supported Types:
   
   1. Array/Slice:
      nums := []int{10, 20, 30}
      for index, value := range nums {
          fmt.Println(index, value)
      }
      // Output: 0 10, 1 20, 2 30
   
   2. Map:
      m := map[string]int{"a": 1, "b": 2}
      for key, value := range m {
          fmt.Println(key, value)
      }
      // Output: a 1, b 2 (order not guaranteed)
   
   3. String (iterates runes):
      for index, char := range "Hello" {
          fmt.Printf("%d: %c\n", index, char)
      }
      // Output: 0: H, 1: e, 2: l, 3: l, 4: o
   
   4. Channel:
      ch := make(chan int)
      for value := range ch {
          fmt.Println(value)  // Receives until channel closed
      }
   
   5. Integer (Go 1.22+):
      for i := range 5 {
          fmt.Println(i)  // 0, 1, 2, 3, 4
      }
   
   Ignoring values:
   for index := range nums { }        // Ignore value
   for _, value := range nums { }     // Ignore index
   for range nums { }                 // Just iterate
   
   Key Point: range works with arrays, slices, maps, strings, channels, and integers!


WHY Questions:
==============

Q1: Why does Go enforce braces {} for conditional and loop code blocks?
A: Why Go enforces braces:
   
   1. Consistency and readability:
      // ✅ Always clear where block starts/ends
      if x > 10 {
          doSomething()
      }
      
      // ❌ Not allowed in Go (prevents ambiguity)
      if x > 10
          doSomething()
   
   2. Prevents bugs (like Apple's goto fail):
      // Famous bug in C (without braces):
      if (error)
          goto fail;
          goto fail;  // Always executes! Bug!
      
      // Go prevents this:
      if error {
          goto fail
      }  // Clear scope
   
   3. Simplifies parsing:
      - No ambiguity for compiler
      - Automatic semicolon insertion works correctly
      - One standard formatting style
   
   4. Enforces Go's philosophy:
      // ❌ Not allowed - opening brace must be on same line
      if x > 10 
      {
          doSomething()
      }
      
      // ✅ Required style
      if x > 10 {
          doSomething()
      }
   
   5. Prevents indentation debates:
      - Python: Indentation is syntax
      - Go: Braces are syntax, gofmt handles indentation
      - No "tabs vs spaces" debates
   
   Key Point: Mandatory braces prevent bugs, ensure consistency, and eliminate formatting debates!

Q2: Why does the switch statement in Go not require break statements by default?
A: Why no break needed:
   
   1. Prevents common bug (fallthrough):
      // C/C++ - Common bug (forgot break):
      switch (x) {
          case 1:
              doA();
              // Bug! Falls through to case 2
          case 2:
              doB();
              break;
      }
      
      // Go - Auto-breaks (safe by default):
      switch x {
      case 1:
          doA()  // Automatically exits
      case 2:
          doB()
      }
   
   2. Cleaner, less boilerplate:
      // Go (concise):
      switch day {
      case 1:
          fmt.Println("Monday")
      case 2:
          fmt.Println("Tuesday")
      case 3:
          fmt.Println("Wednesday")
      }
      
      // vs C (verbose):
      switch (day) {
          case 1:
              printf("Monday");
              break;  // Repetitive!
          case 2:
              printf("Tuesday");
              break;
          case 3:
              printf("Wednesday");
              break;
      }
   
   3. Explicit fallthrough when needed:
      switch x {
      case 1:
          doA()
          fallthrough  // Explicit opt-in
      case 2:
          doB()
      }
   
   4. Safer default behavior:
      - Fallthrough is rare in practice
      - Making it explicit prevents accidents
      - 99% of cases don't need fallthrough
   
   Key Point: Auto-break is safer and cleaner - use fallthrough only when explicitly needed!

Q3: Why is the recover function useful when dealing with panic in Go?
A: Why recover is useful:
   
   1. Prevents program crash:
      func riskyOperation() {
          defer func() {
              if r := recover(); r != nil {
                  fmt.Println("Recovered from:", r)
                  // Program continues!
              }
          }()
          
          panic("something went wrong!")
          // Without recover, program would crash here
      }
   
   2. Graceful error handling:
      func handleRequest() {
          defer func() {
              if r := recover(); r != nil {
                  log.Println("Request failed:", r)
                  // Send error response instead of crashing server
              }
          }()
          
          // Risky code that might panic
      }
   
   3. Cleanup before exit:
      defer func() {
          if r := recover(); r != nil {
              closeConnections()
              saveState()
              log.Fatal("Panic:", r)
          }
      }()
   
   4. Convert panic to error:
      func safeDivide(a, b int) (result int, err error) {
          defer func() {
              if r := recover(); r != nil {
                  err = fmt.Errorf("division error: %v", r)
              }
          }()
          
          return a / b, nil  // May panic if b == 0
      }
   
   5. Protect goroutines:
      go func() {
          defer func() {
              if r := recover(); r != nil {
                  log.Println("Goroutine panic:", r)
                  // Prevents one goroutine crash from killing entire program
              }
          }()
          
          // Risky work
      }()
   
   Key Point: recover allows graceful degradation instead of catastrophic failure!

Q4: Why does Go provide only one looping construct (for) instead of multiple options like while and do-while?
A: Why only for:
   
   1. Simplicity - one keyword, multiple uses:
      // Traditional for
      for i := 0; i < 5; i++ { }
      
      // While-style
      for condition { }
      
      // Infinite
      for { }
      
      // Range
      for i, v := range slice { }
   
   2. Reduces language complexity:
      - Fewer keywords to learn
      - Less cognitive overhead
      - One construct does everything
   
   3. Prevents confusion:
      // C/C++ - Multiple ways to do same thing:
      while (i < 5) { i++; }
      for (; i < 5; ) { i++; }
      do { i++; } while (i < 5);
      
      // Go - One clear way:
      for i < 5 { i++ }
   
   4. Consistency with Go philosophy:
      - "There should be one obvious way to do it"
      - Simplicity over flexibility
      - Less is more
   
   5. Easier to read/maintain:
      // Always starts with 'for' - easy to scan
      for i := 0; i < 10; i++ { }
      for condition { }
      for { }
   
   Key Point: One versatile for keyword is simpler and clearer than multiple loop types!

Q5: Why is defer important for resource management in Go?
A: Why defer is important:
   
   1. Guaranteed cleanup:
      file, err := os.Open("file.txt")
      if err != nil {
          return err
      }
      defer file.Close()  // Always closes, even if panic
      
      // Multiple return paths - defer still executes
      if someCondition {
          return nil  // file.Close() called
      }
      return processFile(file)  // file.Close() called
   
   2. Prevents resource leaks:
      // ❌ Without defer - easy to forget cleanup
      func readFile() error {
          file, _ := os.Open("file.txt")
          if err := process(); err != nil {
              return err  // Forgot to close! Leak!
          }
          file.Close()
          return nil
      }
      
      // ✅ With defer - automatic cleanup
      func readFile() error {
          file, _ := os.Open("file.txt")
          defer file.Close()  // Guaranteed!
          return process()
      }
   
   3. Keeps cleanup code near acquisition:
      mutex.Lock()
      defer mutex.Unlock()  // Clear pairing
      
      conn, _ := net.Dial("tcp", "localhost:8080")
      defer conn.Close()  // Easy to see relationship
   
   4. Works even during panic:
      defer db.Close()  // Closes even if panic occurs
      defer conn.Close()
      // Risky code that might panic
   
   5. Multiple resources handled correctly:
      file1, _ := os.Open("file1.txt")
      defer file1.Close()
      
      file2, _ := os.Open("file2.txt")
      defer file2.Close()  // Both close in reverse order (LIFO)
   
   Key Point: defer ensures resources are always released, preventing leaks and making code safer!


HOW Questions:
==============

Q1: How can you use a for loop in Go to iterate through a slice of integers?
A: How to iterate through a slice:
   
   Method 1: Using range (most common):
   nums := []int{10, 20, 30, 40, 50}
   
   // With index and value
   for index, value := range nums {
       fmt.Printf("Index: %d, Value: %d\n", index, value)
   }
   
   // Only value (ignore index)
   for _, value := range nums {
       fmt.Println(value)
   }
   
   // Only index (ignore value)
   for index := range nums {
       fmt.Println(index)
   }
   
   Method 2: Traditional for loop:
   nums := []int{10, 20, 30, 40, 50}
   
   for i := 0; i < len(nums); i++ {
       fmt.Println(nums[i])
   }
   
   Method 3: While-style loop:
   nums := []int{10, 20, 30, 40, 50}
   i := 0
   
   for i < len(nums) {
       fmt.Println(nums[i])
       i++
   }
   
   Comparison:
   ┌─────────────┬──────────────────────────────────┐
   │ Method      │ Use Case                         │
   ├─────────────┼──────────────────────────────────┤
   │ range       │ Most idiomatic, need values      │
   │ Traditional │ Custom index manipulation        │
   │ While-style │ Rare, complex condition          │
   └─────────────┴──────────────────────────────────┘
   
   Key Point: range is the idiomatic Go way!

Q2: How does the recover function help in handling panics in a Go application?
A: How recover handles panics:
   
   Pattern: Always use recover inside defer:
   func safeFunction() {
       defer func() {
           if r := recover(); r != nil {
               fmt.Println("Recovered from panic:", r)
               // Handle the panic gracefully
           }
       }()
       
       // Code that might panic
       panic("something went wrong!")
   }
   
   Step-by-step:
   
   1. Catch and log panic:
      defer func() {
          if r := recover(); r != nil {
              log.Printf("Panic occurred: %v", r)
          }
      }()
   
   2. Convert panic to error:
      func safeDivide(a, b int) (result int, err error) {
          defer func() {
              if r := recover(); r != nil {
                  err = fmt.Errorf("panic: %v", r)
              }
          }()
          
          result = a / b  // Panics if b == 0
          return result, nil
      }
      
      // Usage:
      result, err := safeDivide(10, 0)
      if err != nil {
          fmt.Println("Error:", err)
      }
   
   3. Cleanup and re-panic:
      defer func() {
          if r := recover(); r != nil {
              cleanup()  // Do cleanup
              panic(r)   // Re-panic after cleanup
          }
      }()
   
   4. Protect HTTP handlers:
      func handler(w http.ResponseWriter, r *http.Request) {
          defer func() {
              if err := recover(); err != nil {
                  http.Error(w, "Internal Server Error", 500)
                  log.Printf("Handler panic: %v", err)
              }
          }()
          
          // Handler code that might panic
      }
   
   5. Protect goroutines:
      go func() {
          defer func() {
              if r := recover(); r != nil {
                  log.Printf("Goroutine panic: %v", r)
              }
          }()
          
          // Goroutine work
      }()
   
   Key Point: Always use recover() inside a deferred function to catch and handle panics!

Q3: How do you implement a switch statement in Go to handle multiple related cases with a single code block?
A: How to handle multiple cases with one code block:
   
   Method 1: Comma-separated values (most common):
   switch month {
   case 1, 3, 5, 7, 8, 10, 12:
       fmt.Println("31 days")
   case 4, 6, 9, 11:
       fmt.Println("30 days")
   case 2:
       fmt.Println("28 or 29 days")
   default:
       fmt.Println("Invalid month")
   }
   
   Method 2: Multiple cases (fallthrough style):
   switch day {
   case "Saturday":
       fallthrough
   case "Sunday":
       fmt.Println("Weekend!")
   default:
       fmt.Println("Weekday")
   }
   
   Method 3: Tagless switch with conditions:
   switch {
   case age < 13:
       fmt.Println("Child")
   case age < 20:
       fmt.Println("Teenager")
   case age < 65:
       fmt.Println("Adult")
   default:
       fmt.Println("Senior")
   }
   
   With nested logic in cases:
   switch x {
   case 1, 2, 3:
       if x == 1 {
           fmt.Println("One")
       } else {
           fmt.Println("Two or Three")
       }
   case 4, 5:
       fmt.Println("Four or Five")
   default:
       fmt.Println("Other")
   }
   
   Real-world example:
   switch httpStatus {
   case 200, 201, 202, 204:
       fmt.Println("Success")
   case 400, 401, 403, 404:
       fmt.Println("Client Error")
   case 500, 502, 503:
       fmt.Println("Server Error")
   default:
       fmt.Println("Unknown Status")
   }
   
   Key Point: Use comma-separated values in a single case to handle multiple related values!

Q4: How do you calculate the sum of all elements in a slice using the range keyword?
A: How to calculate sum using range:
   
   Method 1: Using range with index:
   nums := []int{10, 20, 30, 40, 50}
   total := 0
   
   for i := range nums {
       total += nums[i]
   }
   
   fmt.Println("Sum:", total)  // Sum: 150
   
   Method 2: Using range with value (more idiomatic):
   nums := []int{10, 20, 30, 40, 50}
   total := 0
   
   for _, value := range nums {
       total += value
   }
   
   fmt.Println("Sum:", total)  // Sum: 150
   
   Method 3: Using range with both index and value:
   nums := []int{10, 20, 30, 40, 50}
   total := 0
   
   for index, value := range nums {
       fmt.Printf("Adding nums[%d] = %d\n", index, value)
       total += value
   }
   
   fmt.Println("Sum:", total)  // Sum: 150
   
   Method 4: Traditional for loop:
   nums := []int{10, 20, 30, 40, 50}
   total := 0
   
   for i := 0; i < len(nums); i++ {
       total += nums[i]
   }
   
   fmt.Println("Sum:", total)  // Sum: 150
   
   Best Practice: Use for _, value := range nums - it's cleaner and more idiomatic!

Q5: How does Go's for loop handle infinite loops, and how can you implement one?
A: How to implement infinite loops:
   
   Method 1: Empty for (true infinite loop):
   for {
       // Runs forever until break
       fmt.Println("Running...")
       
       if someCondition {
           break  // Exit loop
       }
   }
   
   Method 2: Explicit true condition:
   for true {
       // Also runs forever
       if someCondition {
           break
       }
   }
   
   Real-world examples:
   
   1. Server loop:
      for {
          conn, err := listener.Accept()
          if err != nil {
              log.Println("Error:", err)
              continue
          }
          go handleConnection(conn)
      }
   
   2. Event loop:
      for {
          select {
          case msg := <-channel:
              process(msg)
          case <-quit:
              return  // Exit loop
          }
      }
   
   3. Retry loop:
      for {
          err := tryOperation()
          if err == nil {
              break  // Success, exit
          }
          time.Sleep(time.Second)
          fmt.Println("Retrying...")
      }
   
   4. Game loop:
      for {
          input := getInput()
          update(input)
          render()
          
          if input == "quit" {
              break
          }
      }
   
   Exit strategies:
   - break - exit loop
   - return - exit function (and loop)
   - os.Exit() - exit program
   
   Key Point: for { } creates an infinite loop - use break or return to exit!


KEY TAKEAWAYS:
==============
- if statements control conditional execution with mandatory braces
- defer executes in LIFO order, perfect for cleanup and resource management
- switch auto-breaks by default, cleaner than if-else chains
- for is the only loop, versatile enough to replace while/do-while
- range simplifies iteration over arrays, slices, maps, strings, channels
- Mandatory braces prevent bugs and ensure consistency
- Auto-break in switch is safer, use fallthrough explicitly when needed
- recover prevents crashes, enables graceful error handling
- One for loop reduces complexity and confusion
- defer guarantees cleanup even during panics
- Use range for idiomatic iteration
- recover must be inside defer to work
- Comma-separated case values handle multiple conditions
- for { } creates infinite loops, exit with break/return
*/

import (
	"fmt"
	"log"
	"time"
)

func main() {
	fmt.Println("=== GO QUIZ 4: Control Flow ===")
	fmt.Println("See comments above for all questions and answers!")
	fmt.Println()

	// Demonstrate key concepts
	fmt.Println("--- if Statement ---")
	x := 15
	if x > 10 {
		fmt.Println("x is greater than 10")
	}

	// if with short statement
	if y := x * 2; y > 20 {
		fmt.Println("y is greater than 20")
	}

	fmt.Println("\n--- defer (LIFO) ---")
	deferDemo()

	fmt.Println("\n--- switch Statement ---")
	switchDemo()

	fmt.Println("\n--- for Loop Variations ---")
	forLoopDemo()

	fmt.Println("\n--- range Keyword ---")
	rangeDemo()

	fmt.Println("\n--- panic and recover ---")
	recoverDemo()

	fmt.Println("\n--- Sum with range ---")
	sumDemo()

	fmt.Println("\n--- Infinite Loop (limited) ---")
	infiniteLoopDemo()
}

func deferDemo() {
	defer fmt.Println("Defer 1 (executes last)")
	defer fmt.Println("Defer 2 (executes second)")
	defer fmt.Println("Defer 3 (executes first)")
	fmt.Println("Normal execution")
}

func switchDemo() {
	month := 3
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		fmt.Println("31 days")
	case 4, 6, 9, 11:
		fmt.Println("30 days")
	case 2:
		fmt.Println("28 or 29 days")
	default:
		fmt.Println("Invalid month")
	}
}

func forLoopDemo() {
	// Traditional
	fmt.Print("Traditional: ")
	for i := 0; i < 3; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// While-style
	fmt.Print("While-style: ")
	j := 0
	for j < 3 {
		fmt.Print(j, " ")
		j++
	}
	fmt.Println()
}

func rangeDemo() {
	nums := []int{10, 20, 30}
	for index, value := range nums {
		fmt.Printf("nums[%d] = %d\n", index, value)
	}
}

func recoverDemo() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	fmt.Println("Before panic")
	panic("intentional panic for demo")
	fmt.Println("This won't print")
}

func sumDemo() {
	nums := []int{10, 20, 30, 40, 50}
	total := 0

	for _, value := range nums {
		total += value
	}

	fmt.Printf("Sum of %v = %d\n", nums, total)
}

func infiniteLoopDemo() {
	count := 0
	for {
		fmt.Println("Loop iteration:", count)
		count++
		if count >= 3 {
			fmt.Println("Breaking out of infinite loop")
			break
		}
		time.Sleep(100 * time.Millisecond)
	}
}

// Real-world example: Safe division with recover
func safeDivide(a, b int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("division error: %v", r)
		}
	}()

	result = a / b
	return result, nil
}

// Example usage of safeDivide
func init() {
	result, err := safeDivide(10, 0)
	if err != nil {
		log.Println("Safe division caught error:", err)
	} else {
		log.Println("Result:", result)
	}
}
