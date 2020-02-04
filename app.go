package main

import (
	"bufio"
	"fmt"
	"os"
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
	fmt.Println(inventory_items[0].code, inventory_items[0].name, ",", inventory_items[0].description, ",", inventory_items[0].price, ",", inventory_items[0].quantity)
	fmt.Println(inventory_items[1].code, inventory_items[1].name, ",", inventory_items[1].description, ",", inventory_items[1].price, ",", inventory_items[1].quantity)
	fmt.Println(inventory_items[2].code, inventory_items[2].name, ",", inventory_items[2].description, ",", inventory_items[2].price, ",", inventory_items[2].quantity)
	fmt.Println(inventory_items[3].code, inventory_items[3].name, ",", inventory_items[3].description, ",", inventory_items[3].price, ",", inventory_items[3].quantity)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter name of the item you wanna buy: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
}
