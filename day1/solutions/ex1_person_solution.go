package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
	City      string
}

func (p Person) FullName() string {
	return p.FirstName + " " + p.LastName
}

func (p Person) IsAdult() bool {
	return p.Age >= 18
}

func (p Person) Introduce() string {
	return fmt.Sprintf("Hi, I'm %s, %d years old, from %s", p.FullName(), p.Age, p.City)
}

func (p *Person) HaveBirthday() {
	p.Age++
}

func (p *Person) MoveTo(newCity string) {
	p.City = newCity
}

func main() {
	person := Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       25,
		City:      "New York",
	}

	fmt.Println("Full Name:", person.FullName())
	fmt.Println("Is Adult:", person.IsAdult())
	fmt.Println("Introduction:", person.Introduce())

	person.HaveBirthday()
	fmt.Printf("After birthday: %d years old\n", person.Age)

	person.MoveTo("San Francisco")
	fmt.Printf("After moving: Lives in %s\n", person.City)

	youngPerson := Person{
		FirstName: "Alice",
		LastName:  "Smith",
		Age:       16,
		City:      "Boston",
	}
	fmt.Println("\nYoung person is adult:", youngPerson.IsAdult())
}
