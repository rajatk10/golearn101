package main

import "fmt"

type Address struct {
	Street string
	City   string
	State  string
	Zip    string
}

func (a Address) FullAddress() string {
	return fmt.Sprintf("%s, %s, %s %s", a.Street, a.City, a.State, a.Zip)
}

type Person struct {
	Name  string
	Email string
	Address
}

type Employee struct {
	Person
	EmployeeID int
	Department string
}

func (e Employee) DisplayInfo() {
	fmt.Printf("Employee: %s (ID: %d)\n", e.Name, e.EmployeeID)
	fmt.Printf("Department: %s\n", e.Department)
	fmt.Printf("Email: %s\n", e.Email)
	fmt.Printf("Address: %s\n", e.FullAddress())
}

func main() {
	emp := Employee{
		Person: Person{
			Name:  "John Doe",
			Email: "john@company.com",
			Address: Address{
				Street: "123 Main St",
				City:   "New York",
				State:  "NY",
				Zip:    "10001",
			},
		},
		EmployeeID: 12345,
		Department: "Engineering",
	}

	fmt.Println("=== Accessing Embedded Fields ===")
	fmt.Printf("Name: %s\n", emp.Name)
	fmt.Printf("City: %s\n", emp.City)
	fmt.Printf("Full Address: %s\n", emp.FullAddress())

	fmt.Println("\n=== Display Full Info ===")
	emp.DisplayInfo()
}
