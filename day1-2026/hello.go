package main

import "fmt"

func main() {
	fmt.Println("Hello World in 2026 !")
	a := 11 // var a int - this is short way of declaring variable in go
	b := 3.14
	fmt.Println("Integer", a)
	fmt.Print("Float: ", b)
	//a = "hello" - type cannot be changed for a variable in Go

	fmt.Println("String", a)

	var a1 int     // 0 is assigned as default
	var b1 string  //null
	var c1 float64 // 0.0
	var d1 bool    // false
	fmt.Println("a1 : ", a1)
	fmt.Println("b1 : ", b1)
	fmt.Println("c1 : ", c1)
	fmt.Println("d1 : ", d1)

	//Constatsn - typed vs untyped

	const pi float64 = 3.14
	const g float64 = 9.8
	const x = 10
	fmt.Println("Typed Constant pi : ", pi)
	fmt.Println("Typed Constant g : ", g)
	fmt.Println("Untyped Constant x : ", x)

	//pi = 56 Constants cannot be changed
	fmt.Println("sum of pi + g :  ", pi+g)
	fmt.Println("Multiply pi*g : ", pi*g)
	fmt.Println("Subtract pi-g : ", pi-g)
	fmt.Println("Divide pi/g : ", pi/g)
	fmt.Println("Remainder operation : ", 10%3)
}
