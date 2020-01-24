package main

import (
	"fmt"
)

/// This is just making a collection of fields with a name and a type
type inventory struct {
	code        string
	name        string
	description string
	price       float64
	quantity    int
}

func main() {
	fmt.Println("Welcome to our grocery store!\nPlease review our inventory and make your selection.")

	inventory_items := []inventory{
		{code: "001", name: "apple", description: "fresh red apple", price: 0.4, quantity: 14},
		{code: "002", name: "dounut", description: "donut with sprinkle", price: 0.99, quantity: 11},
		{code: "003", name: "coke", description: "bottle of coca cola", price: 1.49, quantity: 9},
		{code: "004", name: "bacon", description: "farm fresh bacon", price: 4.99, quantity: 5},
	}
	fmt.Println(inventory_items)
}
