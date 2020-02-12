package main

import (
	"bufio"
	"fmt"
	"os"
)

// This is just making a collection of fields with a name and a type
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
			priceWithOutTaxes := inventoryItem.price
			taxes := inventoryItem.price * .07
			priceWithTaxes := priceWithOutTaxes + taxes
			fmt.Println("okay, great!")
			numberOfQuantityChange--

			fmt.Println(priceWithOutTaxes)
			fmt.Printf("%.1f", priceWithTaxes)
			fmt.Print(numberOfQuantityChange)
			break
		}
	}
	if chosenItem == (inventory{}) {
		fmt.Println("not an option")
	}

}
