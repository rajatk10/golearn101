package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Year   int
	Pages  int
}

func (b Book) String() string {
	return fmt.Sprintf("%s by %s (%d) - %d pages", b.Title, b.Author, b.Year, b.Pages)
}

type Product struct {
	Name    string
	Price   float64
	InStock bool
}

func (p Product) String() string {
	stockStatus := "No"
	if p.InStock {
		stockStatus = "Yes"
	}
	return fmt.Sprintf("%s - $%.2f (In Stock: %s)", p.Name, p.Price, stockStatus)
}

type Temperature struct {
	Value float64
	Unit  string
}

func (t Temperature) String() string {
	return fmt.Sprintf("%.1f°%s", t.Value, t.Unit)
}

func (t Temperature) ToFahrenheit() Temperature {
	if t.Unit == "C" {
		return Temperature{
			Value: t.Value*9/5 + 32,
			Unit:  "F",
		}
	}
	return t
}

func (t Temperature) ToCelsius() Temperature {
	if t.Unit == "F" {
		return Temperature{
			Value: (t.Value - 32) * 5 / 9,
			Unit:  "C",
		}
	}
	return t
}

func main() {
	book := Book{
		Title:  "The Go Programming Language",
		Author: "Donovan & Kernighan",
		Year:   2015,
		Pages:  380,
	}
	fmt.Println(book)

	product1 := Product{
		Name:    "Laptop",
		Price:   999.99,
		InStock: true,
	}
	fmt.Println(product1)

	product2 := Product{
		Name:    "Phone",
		Price:   699.99,
		InStock: false,
	}
	fmt.Println(product2)

	tempC := Temperature{Value: 25.5, Unit: "C"}
	fmt.Println("Temperature in Celsius:", tempC)
	fmt.Println("Temperature in Fahrenheit:", tempC.ToFahrenheit())

	tempF := Temperature{Value: 77.9, Unit: "F"}
	fmt.Println("Temperature in Fahrenheit:", tempF)
	fmt.Println("Temperature in Celsius:", tempF.ToCelsius())
}
