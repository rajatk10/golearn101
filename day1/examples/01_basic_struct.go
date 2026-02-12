package main

import "fmt"

type Person struct {
	Name  string
	Age   int
	Email string
}

func (p Person) Greet() string {
	return fmt.Sprintf("Hello, I'm %s and I'm %d years old", p.Name, p.Age)
}

func (p Person) GetAge() int {
	return p.Age
}

func (p *Person) SetAge(age int) {
	p.Age = age
}

func (p *Person) HaveBirthday() {
	p.Age++
	fmt.Printf("%s is now %d years old!\n", p.Name, p.Age)
}

func main() {
	person1 := Person{
		Name:  "Alice",
		Age:   25,
		Email: "alice@example.com",
	}

	fmt.Println(person1.Greet())

	person1.HaveBirthday()

	person1.SetAge(30)
	fmt.Printf("After setting age: %d\n", person1.GetAge())

	person2 := Person{"Bob", 35, "bob@example.com"}
	fmt.Println(person2.Greet())
}
