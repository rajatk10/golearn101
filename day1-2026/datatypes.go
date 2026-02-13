package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a int = 10
	var b float64 = 10.5
	var c string = "Rajat"
	var d bool = true
	var e rune = 'A'
	fmt.Println("Value of all above is  ", a, b, c, d, e)
	fmt.Println("Type of a is: ", reflect.TypeOf(a))

	//Composite types
	numbers := []int{1, 2, 3, 4, 5}
	person := map[string]string{"name": "Rajat"}
	fmt.Println("Above are composite types ", numbers, person)
	fmt.Println("Type of numbers is: ", reflect.TypeOf(numbers))

	//pointer
	var prt *int = &a
	fmt.Println("Value of prt is: ", prt)
	fmt.Println("Type of prt is: ", reflect.TypeOf(prt))

}

/*
FORMAT VERBS FOR fmt.Printf():
%t - bool (true/false)
%d - int (decimal)
%f - float (3.140000)
%.2f - float with 2 decimal places (3.14)
%s - string
%c - character (from rune or byte)
%v - any value (default format)
%T - type of value
%p - pointer address
%% - literal percent sign

EXAMPLES:
fmt.Printf("%t\n", true)           // true
fmt.Printf("%d\n", 42)             // 42
fmt.Printf("%.2f\n", 3.14159)      // 3.14
fmt.Printf("%s\n", "hello")        // hello
fmt.Printf("%c\n", 72)             // H
fmt.Printf("%T\n", 42)             // int
*/
