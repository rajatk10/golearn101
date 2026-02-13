package main

import "fmt"

/*
CHANNELS NOTES:
1. Channel = Pipe for communication between goroutines
2. Syntax: ch := make(chan Type) - unbuffered channel
3. Send: ch <- value (blocks until received)
4. Receive: value := <-ch (blocks until sent)
5. Channels provide synchronization - no need for time.Sleep()

UNBUFFERED vs BUFFERED:
Unbuffered (default):
- make(chan int) - no buffer
- Send blocks until receive happens
- Receive blocks until send happens

Buffered:
- make(chan int, 3) - buffer size 3
- Send blocks only when buffer is full
- Receive blocks only when buffer is empty

COMPARISON TABLE:
┌─────────────────────┬──────────────────┬──────────────────┐
│ Feature             │ Unbuffered       │ Buffered         │
├─────────────────────┼──────────────────┼──────────────────┤
│ Capacity            │ 0                │ > 0              │
│ Send blocks when    │ No receiver ready│ Buffer full      │
│ Receive blocks when │ No sender ready  │ Buffer empty     │
│ Synchronization     │ Guaranteed       │ Not guaranteed   │
│ Use for             │ Coordination,    │ Decoupling,      │
│                     │ signaling        │ throughput       │
└─────────────────────┴──────────────────┴──────────────────┘

USE CASES:
1. Unbuffered (Synchronization):
   done := make(chan bool)
   go func() {
       // Do work...
       done <- true  // Signal completion
   }()
   <-done  // Wait for signal

2. Buffered (Producer-Consumer):
   jobs := make(chan int, 100)
   go func() {
       for i := 0; i < 1000; i++ {
           jobs <- i  // Won't block until buffer full
       }
       close(jobs)
   }()
   for job := range jobs {
       process(job)
   }

3. Buffered (Rate Limiting):
   semaphore := make(chan struct{}, 3)  // Max 3 concurrent
   for i := 0; i < 10; i++ {
       semaphore <- struct{}{}  // Acquire
       go func(id int) {
           defer func() { <-semaphore }()  // Release
           // Do work...
       }(i)
   }

CLOSING CHANNELS:
- close(ch) - signals no more data will be sent
- value, ok := <-ch - ok is false if closed
- Only sender should close, never receiver
- Closing is optional but good practice

RANGE OVER CHANNELS:
for value := range ch {
    // Receives until channel is closed
}

SELECT STATEMENT (multiple channels):
select {
case msg := <-ch1:
    // Handle ch1
case msg := <-ch2:
    // Handle ch2
case <-time.After(1 * time.Second):
    // Timeout
}

DIRECTIONAL CHANNELS:
- chan<- Type - send-only channel
- <-chan Type - receive-only channel
- Used in function parameters for safety

KEY POINTS:
- Channels block by default (synchronization)
- Send and receive must happen in different goroutines (or buffered)
- Deadlock occurs if no goroutine can proceed
- Close channel when done sending (sender's responsibility)

make() FUNCTION:
1. Built-in function to create slices, maps, and channels
2. Returns initialized, ready-to-use value (not pointer)
3. Required for channels and maps (nil ones are unusable)

SYNTAX:
- Channel: make(chan Type) or make(chan Type, bufferSize)
- Slice: make([]Type, length) or make([]Type, length, capacity)
- Map: make(map[KeyType]ValueType) or make(map[KeyType]ValueType, capacity)

make() vs new():
- make(): Returns initialized value, only for slice/map/channel
- new(): Returns pointer to zero value, works for any type
- Example: make(chan int) returns chan int (usable)
           new(chan int) returns *chan int (nil, NOT usable)

WHY make() IS NEEDED:
var ch chan int  // nil channel - CANNOT use
ch := make(chan int)  // initialized channel - ready to use
*/

func main() {
	ch := make(chan int)
	go func() {
		ch <- 1
		close(ch)
	}()
	fmt.Println("Use channel to pass data between goroutines")
	fmt.Println(<-ch)
	chs := make(chan string)
	go func() {
		chs <- "Hello World Channel"
		chs <- "Hello 2 World"
		chs <- "Hello 3 World"
		close(chs)
	}()
	//fmt.Printf("Another channel %s \n", <-chs)

	for value := range chs {
		fmt.Println(value)
	}
}
