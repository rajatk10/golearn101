package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

/*
GO GOROUTINES, CHANNELS & CONCURRENCY QUIZ - Questions & Answers

WHAT Questions:
================

Q1: What is a goroutine in Go, and how does it differ from a traditional thread?
A: Goroutine:
   - Lightweight concurrent function execution
   - Managed by Go runtime (not OS)
   - Started with 'go' keyword: go myFunction()
   
   Differences from OS Threads:
   ┌──────────────┬─────────────┬────────────┐
   │ Aspect       │ Goroutine   │ OS Thread  │
   ├──────────────┼─────────────┼────────────┤
   │ Stack Size   │ ~2KB        │ ~1-2MB     │
   │ Management   │ Go runtime  │ OS kernel  │
   │ Creation     │ ~1μs        │ ~10-100μs  │
   │ Switching    │ ~200ns      │ ~1-2μs     │
   │ Scalability  │ Millions    │ Thousands  │
   │ Memory       │ User-space  │ Kernel     │
   └──────────────┴─────────────┴────────────┘
   
   Key Points:
   - Goroutines are multiplexed onto OS threads (M:N model)
   - Go runtime scheduler manages goroutines efficiently
   - Can create millions of goroutines without crashing
   - Context switching happens in user-space (faster)
   
   Example:
   // Create 10,000 goroutines easily
   for i := 0; i < 10000; i++ {
       go func(id int) {
           fmt.Println("Goroutine", id)
       }(i)
   }
   // Would crash with OS threads!

Q2: What are channels in Go, and how do they facilitate communication between goroutines?
A: Channels:
   - Typed conduits for sending/receiving data between goroutines
   - Created with make(): ch := make(chan Type)
   - Thread-safe communication mechanism
   
   Types:
   1. Unbuffered (synchronous):
      ch := make(chan int)
      - Direct handoff
      - Sender blocks until receiver ready
      - Receiver blocks until sender sends
   
   2. Buffered (asynchronous):
      ch := make(chan int, 5)
      - Temporary storage (size 5)
      - Sender blocks only when buffer full
      - Receiver blocks only when buffer empty
   
   Operations:
   ch <- value    // Send to channel
   value := <-ch  // Receive from channel
   close(ch)      // Close channel
   
   How They Facilitate Communication:
   - Synchronization: Automatic coordination between goroutines
   - Data Transfer: Safe passing of data without locks
   - Ownership Transfer: Clear ownership semantics
   - Signaling: Notify completion or events
   
   Example:
   func worker(jobs <-chan int, results chan<- int) {
       for job := range jobs {
           results <- job * 2  // Process and send result
       }
   }
   
   func main() {
       jobs := make(chan int, 5)
       results := make(chan int, 5)
       
       go worker(jobs, results)
       
       // Send jobs
       for i := 1; i <= 3; i++ {
           jobs <- i
       }
       close(jobs)
       
       // Receive results
       for i := 1; i <= 3; i++ {
           fmt.Println(<-results)
       }
   }
   
   Go Proverb: "Don't communicate by sharing memory; share memory by communicating."

Q3: What is the purpose of the select statement in Go?
A: Select Statement:
   - Multiplexes multiple channel operations
   - Like switch, but for channels
   - Waits on multiple channels simultaneously
   
   Purpose:
   1. Wait on multiple channels at once
   2. Handle whichever channel is ready first
   3. Non-blocking operations (with default)
   4. Implement timeouts
   5. Coordinate complex channel operations
   
   Syntax:
   select {
   case msg1 := <-ch1:
       // Handle ch1
   case msg2 := <-ch2:
       // Handle ch2
   case ch3 <- value:
       // Send to ch3
   default:
       // Non-blocking (optional)
   }
   
   Behavior:
   - Checks ALL cases simultaneously (not sequential)
   - Executes first ready case
   - Random choice if multiple ready
   - Blocks if none ready (unless default present)
   
   Use Cases:
   1. Multiple channel operations:
      select {
      case msg := <-ch1:
          fmt.Println("From ch1:", msg)
      case msg := <-ch2:
          fmt.Println("From ch2:", msg)
      }
   
   2. Timeout:
      select {
      case result := <-ch:
          fmt.Println("Got:", result)
      case <-time.After(2 * time.Second):
          fmt.Println("Timeout!")
      }
   
   3. Non-blocking:
      select {
      case msg := <-ch:
          fmt.Println("Got:", msg)
      default:
          fmt.Println("No message available")
      }
   
   Key Points:
   - Order in code doesn't matter
   - All cases evaluated simultaneously
   - Pseudo-random selection when multiple ready
   - Efficient (no polling, no CPU waste)


WHY Questions:
==============

Q1: Why are goroutines considered lightweight compared to OS threads?
A: Goroutines are Lightweight Because:
   
   1. Small Stack Size:
      OS Thread:  ~1-2 MB (fixed, pre-allocated)
      Goroutine:  ~2 KB (starts small, grows dynamically)
      Result: 500x less memory per goroutine!
   
   2. User-Space Scheduling:
      OS Thread:
      - Context switch requires kernel mode
      - Save/restore CPU registers
      - Expensive system calls
      - ~1-2 microseconds per switch
      
      Goroutine:
      - Context switch in user space
      - Go runtime handles scheduling
      - No kernel involvement
      - ~200 nanoseconds per switch (10x faster!)
   
   3. M:N Scheduling Model:
      - M goroutines multiplexed on N OS threads
      - Example: 100,000 goroutines on 4 OS threads
      - Runtime manages mapping efficiently
      - OS sees only a few threads
   
   4. Fast Creation/Destruction:
      OS Thread:
      - System call to kernel
      - Memory allocation
      - TLB flush
      - ~10-100 microseconds
      
      Goroutine:
      - Simple function call
      - Stack allocation
      - ~1 microsecond
   
   5. Efficient Memory:
      // 100,000 goroutines
      Memory: ~200 MB (2KB × 100,000)
      
      // 100,000 OS threads
      Memory: ~200 GB! (Would crash)
   
   6. Growing Stacks:
      - Start with 2KB
      - Grow automatically when needed
      - Shrink when not needed
      - No wasted memory
   
   Summary:
   ✅ Smaller: 2KB vs 1-2MB
   ✅ Faster: User-space switching
   ✅ Scalable: Millions vs thousands
   ✅ Efficient: Go runtime optimizes scheduling
   ✅ Cheap: Fast creation/destruction

Q2: Why are channels preferred over shared memory for communication between goroutines?
A: Channels Preferred Over Shared Memory Because:
   
   1. Prevents Race Conditions:
      Shared Memory (Risky):
      var counter int  // Shared variable
      var mu sync.Mutex
      
      go func() {
          mu.Lock()
          counter++  // Must remember to lock!
          mu.Unlock()
      }()
      // Easy to forget lock → race condition
      
      Channels (Safe):
      ch := make(chan int)
      go func() {
          ch <- 1  // No race possible
      }()
      val := <-ch  // Synchronized automatically
   
   2. Clear Ownership:
      Shared Memory:
      var sharedData []int  // Who owns this? Unclear!
      // Multiple goroutines access it
      // Need locks everywhere
      
      Channels:
      ch := make(chan []int)
      data := []int{1, 2, 3}
      ch <- data  // Ownership transferred to receiver
      // Sender shouldn't use data anymore
   
   3. Easier to Reason About:
      Shared Memory:
      mu1.Lock()
      mu2.Lock()
      // Do work
      mu2.Unlock()
      mu1.Unlock()
      // Deadlock risk! Lock order matters
      
      Channels:
      ch1 <- data      // Send
      result := <-ch2  // Receive
      // Clear flow, no deadlock concerns
   
   4. Built-in Synchronization:
      Shared Memory:
      var ready bool
      var mu sync.Mutex
      
      go func() {
          mu.Lock()
          ready = true
          mu.Unlock()
      }()
      
      // Wait for ready (manual polling)
      for {
          mu.Lock()
          if ready { break }
          mu.Unlock()
      }
      
      Channels:
      done := make(chan bool)
      
      go func() {
          // Do work
          done <- true  // Signal completion
      }()
      
      <-done  // Wait automatically
   
   5. Go Philosophy:
      "Don't communicate by sharing memory;
       share memory by communicating."
   
   When to Use Each:
   ┌─────────────────┬──────────────────┐
   │ Use Channels    │ Use Mutex        │
   ├─────────────────┼──────────────────┤
   │ Data transfer   │ Shared state     │
   │ Coordination    │ Caching          │
   │ Events          │ Counters         │
   │ Pipelines       │ Short critical   │
   │ Ownership       │ sections         │
   └─────────────────┴──────────────────┘
   
   Benefits:
   ✅ No explicit locks needed
   ✅ Clear communication flow
   ✅ Prevents race conditions by design
   ✅ Easier to understand and maintain
   ✅ Composable patterns

Q3: Why would you use a select statement instead of polling multiple channels?
A: Select vs Polling:
   
   Polling (Bad Approach):
   // ❌ Inefficient polling
   for {
       select {
       case msg := <-ch1:
           fmt.Println(msg)
       default:
           // No message
       }
       
       select {
       case msg := <-ch2:
           fmt.Println(msg)
       default:
           // No message
       }
       
       time.Sleep(10 * time.Millisecond)  // Waste CPU
   }
   
   Problems with Polling:
   🔴 Wastes CPU - constantly checking
   🔴 Delays - might miss messages between checks
   🔴 Complex code - multiple checks needed
   🔴 Battery drain on mobile/embedded
   🔴 Inefficient - busy-waiting
   
   Select (Good Approach):
   // ✅ Efficient with select
   for {
       select {
       case msg := <-ch1:
           fmt.Println("From ch1:", msg)
       case msg := <-ch2:
           fmt.Println("From ch2:", msg)
       }
   }
   
   Benefits of Select:
   ✅ Blocks efficiently - no CPU waste
   ✅ Instant response - wakes immediately when ready
   ✅ Clean code - single statement
   ✅ Fair - all channels checked simultaneously
   ✅ Go runtime optimized - built-in scheduler support
   
   Comparison:
   Polling:
   - Check ch1 → Nothing? Wait...
   - Check ch2 → Nothing? Wait...
   - Check ch1 again...
   - Repeat forever (CPU spinning!)
   
   Select:
   - Wait on BOTH channels
   - Wake up when ANY is ready
   - Process immediately
   - No wasted cycles
   
   Think of it like:
   - Polling = Constantly checking mailbox every minute
   - Select = Mailbox notifies you when mail arrives
   
   Reasons to Use Select:
   1. Efficiency - No busy-waiting
   2. Responsiveness - Instant reaction
   3. Simplicity - One statement handles multiple channels
   4. Fairness - All channels treated equally
   5. Scalability - Works with many channels


HOW Questions:
==============

Q1: How does Go's runtime manage goroutines efficiently?
A: Go Runtime Management (G-M-P Model):
   
   The G-M-P Model:
   G = Goroutine (the task to execute)
   M = Machine (OS thread)
   P = Processor (scheduling context, logical CPU)
   
   Structure:
   [G1][G2][G3]...[Gn]  ← Many goroutines (user code)
        ↓    ↓    ↓
       [P1] [P2] [P3]    ← GOMAXPROCS processors
        ↓    ↓    ↓
       [M1] [M2] [M3]    ← Few OS threads
   
   Key Mechanisms:
   
   1. Work Stealing:
      P1: [G1][G2][G3][G4]  ← Busy processor
      P2: [empty]           ← Idle processor
      
      P2 steals from P1:
      P1: [G1][G2]
      P2: [G3][G4]          ← Balanced!
   
   2. Cooperative Scheduling:
      Goroutine yields at:
      - Channel operations (ch <- x, <-ch)
      - System calls (file I/O, network)
      - Function calls (compiler inserts checks)
      - time.Sleep()
      - runtime.Gosched() (explicit yield)
   
   3. Growing Stacks:
      Start: 2KB stack
      ↓
      Needs more? Allocate larger stack
      ↓
      Copy old stack to new
      ↓
      Continue execution
   
   4. Blocking Optimization:
      Goroutine blocks on I/O:
      G1 blocks → M1 detaches from P1
                → New M2 created for P1
                → Other goroutines continue
                → When G1 unblocks, reattaches
   
   5. Preemption (Go 1.14+):
      Long-running goroutine?
      → Runtime sends signal
      → Goroutine preempted
      → Other goroutines get CPU time
   
   Example:
   func main() {
       runtime.GOMAXPROCS(4)  // 4 processors
       
       for i := 0; i < 1000; i++ {
           go func(id int) {
               // 1000 goroutines
               // Multiplexed on 4 OS threads
               fmt.Println(id)
           }(i)
       }
       
       time.Sleep(1 * time.Second)
   }
   
   Efficiency Features:
   ✅ User-space scheduling - No kernel involvement
   ✅ Small stacks - 2KB vs 1-2MB
   ✅ Fast context switch - ~200ns vs ~1-2μs
   ✅ Work stealing - Automatic load balancing
   ✅ Cooperative + preemptive - Fair scheduling
   ✅ Scalable - Millions of goroutines possible
   ✅ Blocking optimization - Doesn't waste threads

Q2: How can you synchronize multiple goroutines without using explicit locks?
A: Synchronization Without Locks:
   
   Methods:
   
   1. Channels (Primary Method):
      func main() {
          done := make(chan bool)
          
          go func() {
              fmt.Println("Working...")
              time.Sleep(1 * time.Second)
              done <- true  // Signal completion
          }()
          
          <-done  // Wait for completion (synchronized!)
          fmt.Println("Done!")
      }
   
   2. WaitGroup:
      func main() {
          var wg sync.WaitGroup
          
          for i := 0; i < 3; i++ {
              wg.Add(1)
              go func(id int) {
                  defer wg.Done()
                  fmt.Println("Goroutine", id)
              }(i)
          }
          
          wg.Wait()  // Wait for all (no mutex!)
      }
   
   3. Channel Pipeline:
      func producer(ch chan<- int) {
          for i := 0; i < 5; i++ {
              ch <- i
          }
          close(ch)
      }
      
      func consumer(ch <-chan int, done chan<- bool) {
          for val := range ch {
              fmt.Println(val)
          }
          done <- true
      }
      
      func main() {
          ch := make(chan int)
          done := make(chan bool)
          
          go producer(ch)
          go consumer(ch, done)
          
          <-done  // Synchronized without locks!
      }
   
   4. sync.Once (One-time initialization):
      var once sync.Once
      var config Config
      
      func getConfig() Config {
          once.Do(func() {
              config = loadConfig()  // Runs only once, thread-safe
          })
          return config
      }
   
   5. Atomic Operations:
      var counter int64
      
      func increment() {
          atomic.AddInt64(&counter, 1)  // No mutex needed!
      }
      
      func getCounter() int64 {
          return atomic.LoadInt64(&counter)
      }
   
   Key Point:
   Channels = Built-in synchronization
   - No explicit Lock()/Unlock()
   - Synchronization happens automatically
   - Follows Go's philosophy: "Share memory by communicating"
   
   Comparison:
   ┌─────────────────┬──────────────────┐
   │ With Locks      │ Without Locks    │
   ├─────────────────┼──────────────────┤
   │ mu.Lock()       │ ch <- data       │
   │ counter++       │ val := <-ch      │
   │ mu.Unlock()     │                  │
   │                 │                  │
   │ Manual sync     │ Automatic sync   │
   │ Error-prone     │ Safer            │
   └─────────────────┴──────────────────┘

Q3: How can you use the select statement to handle a timeout in a program?
A: Timeout with Select:
   
   Use time.After() for timeouts:
   
   Basic Timeout:
   select {
   case result := <-ch:
       fmt.Println("Got result:", result)
   case <-time.After(2 * time.Second):
       fmt.Println("Timeout! Took too long")
   }
   
   How it Works:
   - time.After(duration) returns a channel
   - Sends current time after duration
   - select waits for EITHER channel
   - Whichever comes first wins
   
   Complete Example:
   func slowOperation(ch chan string) {
       time.Sleep(3 * time.Second)  // Takes 3 seconds
       ch <- "Done!"
   }
   
   func main() {
       ch := make(chan string)
       
       go slowOperation(ch)
       
       select {
       case result := <-ch:
           fmt.Println("Success:", result)
       case <-time.After(2 * time.Second):
           fmt.Println("Timeout after 2 seconds!")
       }
   }
   // Output: Timeout after 2 seconds!
   
   default vs time.After():
   
   default - Non-blocking (immediate):
   select {
   case msg := <-ch:
       fmt.Println("Got:", msg)
   default:
       fmt.Println("No message RIGHT NOW")
       // Executes immediately if ch not ready
   }
   
   time.After() - Timeout (wait with limit):
   select {
   case msg := <-ch:
       fmt.Println("Got:", msg)
   case <-time.After(5 * time.Second):
       fmt.Println("Gave up after 5 seconds")
       // Waits up to 5 seconds
   }
   
   More Timeout Patterns:
   
   1. Context with Timeout (Modern Approach):
      ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
      defer cancel()
      
      select {
      case result := <-ch:
          fmt.Println("Got:", result)
      case <-ctx.Done():
          fmt.Println("Timeout:", ctx.Err())
      }
   
   2. Multiple Operations with Timeout:
      select {
      case msg1 := <-ch1:
          fmt.Println("From ch1:", msg1)
      case msg2 := <-ch2:
          fmt.Println("From ch2:", msg2)
      case <-time.After(1 * time.Second):
          fmt.Println("Neither channel responded in time")
      }
   
   3. Periodic Timeout:
      timeout := time.After(5 * time.Second)
      tick := time.Tick(1 * time.Second)
      
      for {
          select {
          case <-tick:
              fmt.Println("Tick")
          case <-timeout:
              fmt.Println("Timeout!")
              return
          }
      }
   
   Key Points:
   ✅ time.After() creates timeout channel
   ✅ select picks first ready channel
   ✅ Clean timeout handling
   ✅ No manual timer management
   ✅ Composable with other operations


KEY TAKEAWAYS:
==============
- Goroutines are lightweight, managed by Go runtime (not OS)
- Channels provide safe communication between goroutines
- Unbuffered channels = synchronous, Buffered = asynchronous
- select multiplexes multiple channel operations
- G-M-P model: goroutines multiplexed on OS threads
- Work stealing balances load across processors
- Channels preferred over shared memory (Go philosophy)
- Synchronize without locks using channels, WaitGroup, atomic
- time.After() for timeouts in select
- User-space scheduling makes goroutines efficient
*/

// Example 1: Basic Goroutine
func example1() {
	fmt.Println("\n=== Example 1: Basic Goroutine ===")

	go func() {
		fmt.Println("Hello from goroutine!")
	}()

	time.Sleep(100 * time.Millisecond)
	fmt.Println("Main function")
}

// Example 2: Channel Communication
func example2() {
	fmt.Println("\n=== Example 2: Channel Communication ===")

	ch := make(chan string)

	go func() {
		ch <- "Message from goroutine"
	}()

	msg := <-ch
	fmt.Println(msg)
}

// Example 3: Buffered Channel
func example3() {
	fmt.Println("\n=== Example 3: Buffered Channel ===")

	ch := make(chan int, 3)

	ch <- 1
	ch <- 2
	ch <- 3

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

// Example 4: Select Statement
func example4() {
	fmt.Println("\n=== Example 4: Select Statement ===")

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "from ch1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "from ch2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Received:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Received:", msg2)
		}
	}
}

// Example 5: WaitGroup
func example5() {
	fmt.Println("\n=== Example 5: WaitGroup ===")

	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Goroutine %d working\n", id)
			time.Sleep(100 * time.Millisecond)
		}(i)
	}

	wg.Wait()
	fmt.Println("All goroutines completed")
}

// Example 6: Timeout with Select
func example6() {
	fmt.Println("\n=== Example 6: Timeout with Select ===")

	ch := make(chan string)

	go func() {
		time.Sleep(3 * time.Second)
		ch <- "Done!"
	}()

	select {
	case result := <-ch:
		fmt.Println("Success:", result)
	case <-time.After(2 * time.Second):
		fmt.Println("Timeout after 2 seconds!")
	}
}

// Example 7: Worker Pool
func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep(100 * time.Millisecond)
		results <- job * 2
	}
}

func example7() {
	fmt.Println("\n=== Example 7: Worker Pool ===")

	jobs := make(chan int, 5)
	results := make(chan int, 5)

	// Start 3 workers
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Send 5 jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// Collect results
	for a := 1; a <= 5; a++ {
		fmt.Println("Result:", <-results)
	}
}

// Example 8: Atomic Operations
func example8() {
	fmt.Println("\n=== Example 8: Atomic Operations ===")

	var counter int64
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&counter, 1)
		}()
	}

	wg.Wait()
	fmt.Println("Final counter:", atomic.LoadInt64(&counter))
}

// Example 9: Context with Timeout
func example9() {
	fmt.Println("\n=== Example 9: Context with Timeout ===")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	ch := make(chan string)

	go func() {
		time.Sleep(3 * time.Second)
		ch <- "Done!"
	}()

	select {
	case result := <-ch:
		fmt.Println("Success:", result)
	case <-ctx.Done():
		fmt.Println("Timeout:", ctx.Err())
	}
}

func main() {
	fmt.Println("=== GO QUIZ 6: Goroutines, Channels & Concurrency ===")
	fmt.Println("See comments above for all questions and answers!")

	example1()
	example2()
	example3()
	example4()
	example5()
	example6()
	example7()
	example8()
	example9()

	fmt.Println("\n=== All Examples Completed ===")
}
