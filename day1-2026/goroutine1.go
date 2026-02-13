package main

import (
	"fmt"
	"time"
)

/*
GOROUTINES NOTES:
1. Goroutine = Lightweight thread managed by Go runtime
2. Syntax: go functionName() - starts function in new goroutine
3. Main goroutine exits = all other goroutines are killed
4. Goroutines run concurrently with main function
5. Use time.Sleep() temporarily to wait (not recommended for production)
6. Use sync.WaitGroup or channels for proper synchronization

CONCURRENCY vs PARALLELISM:
Concurrency:
- Dealing with multiple things at once (structure)
- Multiple tasks making progress (not necessarily simultaneously)
- Like a chef switching between cooking multiple dishes
- Go provides this via goroutines

Parallelism:
- Doing multiple things at once (execution)
- Multiple tasks running simultaneously on different CPU cores
- Like multiple chefs each cooking a dish
- Go runtime automatically uses available CPU cores
***

We used time.Sleep(1 * time.Second) in main() to give goroutines time to complete.
Without it, the main program might finish before the goroutines run.
Always assume the main function may exit before goroutines complete.
Use synchronization tools like sync.WaitGroup
(a WaitGroup is used to wait for a collection of goroutines to finish executing) or channels to coordinate goroutines.

***
KEY POINTS:
- Go gives you concurrency (goroutines)
- Go runtime decides parallelism based on CPU cores
- time.Sleep() gives goroutine time to execute before main exits
- Without wait mechanism, main exits and kills goroutines

PROPER SYNCHRONIZATION (WaitGroup):
var wg sync.WaitGroup
wg.Add(1)           // Expecting 1 goroutine
go func() {
    defer wg.Done() // Signal completion
    // work here
}()
wg.Wait()           // Wait for all goroutines
*/

func sayHello11() {
	fmt.Println("Hello World")
}
func main() {
	go sayHello11()
	time.Sleep(1 * time.Second)
	//code works because time.Sleep() gives the goroutine time to execute before main exits.
	//Without it, main would exit immediately and kill the goroutine!
	fmt.Println("Main function")
}
