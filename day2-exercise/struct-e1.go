package main

import (
	"fmt"
)

type Book struct {
	title  string
	author string
	pages  int
}

type Person struct {
	age  int
	name string
}

func display(b Book) {
	fmt.Println("Title:", b.title)
	fmt.Println("Author:", b.author)
	fmt.Println("Pages:", b.pages)
}

func main() {
	b := Book{"The Go Programming Language", "Alan A. A. Donovan", 250}
	display(b)
	p1 := Person{34, "Ryan"}
	p2 := Person{56, "Ronny"}
	fmt.Printf("Difference in age of %v and %v is : %d \n", p1.name, p2.name, p2.age-p1.age)
}
