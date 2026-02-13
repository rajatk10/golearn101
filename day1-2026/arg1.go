package main

import (
	"fmt"
	"os"
	"reflect"
)

/*
COMMAND-LINE ARGUMENTS NOTES:

os.Args:
1. os.Args = []string - slice of strings containing all arguments
2. os.Args[0] = Program name (or temp path with 'go run')
3. os.Args[1:] = Actual arguments passed by user
4. ALL arguments are strings, even if they look like numbers

GETTING TYPE:
- fmt.Printf("%T", variable) - Print type using %T format verb
- reflect.TypeOf(variable) - Get type using reflect package
- Go has NO type() function like Python

CONVERTING ARGUMENTS:
- strconv.Atoi(str) - String to int
- strconv.ParseFloat(str, 64) - String to float64
- strconv.ParseBool(str) - String to bool

EXAMPLE:
go run arg1.go 10 20 30
os.Args = []string{"arg1", "10", "20", "30"}
           [0]      [1]   [2]   [3]

CHECKING ARGUMENTS:
if len(os.Args) < 2 {
    fmt.Println("Usage: program <args>")
    return
}

BETTER APPROACH (flag package):
import "flag"
name := flag.String("name", "default", "description")
flag.Parse()
// Run: go run prog.go -name=value
*/

func main() {
	args := os.Args //Get all arguments
	fmt.Println("Args is of type ", reflect.TypeOf(args))
	for i := range len(args) {
		fmt.Printf("Argument is %s \n", args[i])
		fmt.Println(reflect.TypeOf(args[i]))
	}
}

/*
	Args is of type  []string
	Argument is /var/folders/1v/gtp66rwn76j69n0mw8cqft7m0000gn/T/go-build2272869841/b001/exe/arg1
	string
	Argument is 1
	string
	Argument is 2
	string
	Argument is 3
	string
*/
