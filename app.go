package main

import (
	"bufio"
	"fmt"
	"os"
)

// This is a structure that is declared with strings field, floating-point field, and int field.
type inventory struct {
	code        string
	name        string
	description string
	price       float64
	quantity    int
}

func main() {
	fmt.Println("Welcome to our grocery store!\nPlease review our inventory and make your selection.")

	inventoryItems := []inventory{
		{code: "001", name: "apple", description: "fresh red apple", price: 0.4, quantity: 14},
		{code: "002", name: "dounut", description: "donut with sprinkle", price: 0.99, quantity: 11},
		{code: "003", name: "coke", description: "bottle of coca cola", price: 1.49, quantity: 9},
		{code: "004", name: "bacon", description: "farm fresh bacon", price: 4.99, quantity: 5},
	}
	fmt.Println(inventoryItems[0].code, inventoryItems[0].name, ",", inventoryItems[0].description, ",", inventoryItems[0].price, ",", inventoryItems[0].quantity)
	fmt.Println(inventoryItems[1].code, inventoryItems[1].name, ",", inventoryItems[1].description, ",", inventoryItems[1].price, ",", inventoryItems[1].quantity)
	fmt.Println(inventoryItems[2].code, inventoryItems[2].name, ",", inventoryItems[2].description, ",", inventoryItems[2].price, ",", inventoryItems[2].quantity)
	fmt.Println(inventoryItems[3].code, inventoryItems[3].name, ",", inventoryItems[3].description, ",", inventoryItems[3].price, ",", inventoryItems[3].quantity)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter name of the item you wanna buy: ")
	scanner.Scan()
	text := scanner.Text()
	fmt.Println(text)
	var chosenItem inventory
	for _, inventoryItem := range inventoryItems {
		if text == inventoryItem.name {
			chosenItem = inventoryItem
			numberOfQuantityChange := inventoryItem.quantity
			subTotal := inventoryItem.price
			taxes := inventoryItem.price * .07
			total := subTotal + taxes
			numberOfQuantityChange--
			fmt.Println("Subtotal: $", subTotal)
			fmt.Print("Totol: $")
			fmt.Printf("%.2f", total)
		}
	}
	if chosenItem == (inventory{}) {
		fmt.Println("not an option")
	}

}
