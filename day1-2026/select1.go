package main

import (
	"fmt"
	"time"
)

/*
SELECT STATEMENT NOTES:
1. select = Switch statement for channels
2. Waits on multiple channel operations simultaneously
3. Executes whichever case is ready first
4. If multiple cases ready, picks one randomly
5. Blocks until at least one case can proceed

SYNTAX:
select {
case msg := <-ch1:
    // Handle ch1
case msg := <-ch2:
    // Handle ch2
default:
    // Optional: runs if no channel ready (non-blocking)
}

ANONYMOUS FUNCTION CALL:
- func() { }() - Define AND call immediately (parentheses at end)
- myFunc := func() { } - Define but don't call (store for later)
- myFunc() - Call the stored function later
- go func() { }() - Must have () to call the goroutine

KEY POINTS:
- select blocks until a channel is ready
- Use in loop to receive multiple messages
- default case makes select non-blocking
- Useful for timeouts: case <-time.After(1 * time.Second)
*/

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Hello 1 World"
	}()
	go func() {
		time.Sleep(3 * time.Second)
		ch2 <- "Hello World 2"
	}()
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		}
	}
}
