package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Animal interface {
	Speaker
	Move() string
}

type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return "Woof! Woof!"
}

func (d Dog) Move() string {
	return "Running on four legs"
}

type Cat struct {
	Name string
}

func (c Cat) Speak() string {
	return "Meow!"
}

func (c Cat) Move() string {
	return "Sneaking silently"
}

type Robot struct {
	Model string
}

func (r Robot) Speak() string {
	return "Beep boop!"
}

func MakeItSpeak(s Speaker) {
	fmt.Printf("It says: %s\n", s.Speak())
}

func DescribeAnimal(a Animal) {
	fmt.Printf("Sound: %s\n", a.Speak())
	fmt.Printf("Movement: %s\n", a.Move())
}

func main() {
	dog := Dog{Name: "Buddy"}
	cat := Cat{Name: "Whiskers"}
	robot := Robot{Model: "R2D2"}

	fmt.Println("=== Speaker Interface ===")
	MakeItSpeak(dog)
	MakeItSpeak(cat)
	MakeItSpeak(robot)

	fmt.Println("\n=== Animal Interface ===")
	DescribeAnimal(dog)
	fmt.Println()
	DescribeAnimal(cat)
}
