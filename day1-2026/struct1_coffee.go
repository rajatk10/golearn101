package main

import (
	"bufio"
	"fmt"
	"os"
)

// Exploring different types of struct in Go
var REFILL int = 100

// 1. Define a struct

type CoffeShop struct {
	Name     string
	Location string
	//Stock int
	Greet func(name string, cs CoffeShop)
}

func greetUser(username string, cs CoffeShop) {
	fmt.Printf("Hello %s, Welcome to %s \n", username, cs.Name)
}

type CoffeeOrder struct {
	Name     string
	Size     string
	Price    float64
	Quantity int
}

//It follows pascal case, declared at root of file so can be exported and used by any function

func OrderCoffee() {
	fmt.Println("May I please know your order")
	name := bufio.NewScanner(os.Stdin)
	name.Scan()
	order := &CoffeeOrder{
		Name:     name.Text(),
		Size:     "small",
		Price:    2.99,
		Quantity: 1,
	}
	fmt.Println(order)
	order.AddToCart()
}

// Struct by Refrence
func (c *CoffeeOrder) AddToCart() {
	fmt.Println("Adding To Cart")
	remaining := REFILL - c.Quantity
	if remaining == 0 {
		fmt.Println("Sorry, we are out of stock")
		return
	}
	fmt.Println("SUCCESS : Added to cart")

}

func main() {
	myShop := CoffeShop{
		Name:     "Capgemini - Brew & Beans",
		Location: "Bangalore",
		Greet:    greetUser,
	}
	//myShop.Greet("John")
	greetUser("John", myShop) //Hello John, Welcome to Capgemini - Brew & Beans
	OrderCoffee()
	//shop.greetUser("John")
	//shop.Greet(shop)
	//shop.Greet("John")
	//OrderCoffee("John")
}
