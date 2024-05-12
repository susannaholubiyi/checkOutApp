package main

import (
	"fmt"
	"time"
)

func main() {
	var checkOut CheckOutApp
	checkOut.checkOutInformation()
}

type CheckOutApp struct {
	customersName string
	cashiersName  string
	discount      float64
	amountPaid    float64
	subTotal      float64
	itemBought    []string
	itemQuantity  []int
	itemPrice     []float64
}

func (c *CheckOutApp) getDateAndTime() string {
	currentTime := time.Now()
	currentDateTime := currentTime.Format("2006-01-02 03:04:05 PM")
	return currentDateTime
}
func (c *CheckOutApp) checkOutInformation() {
	var item string
	var numberOfPieces int
	var pricePerUnit float64

	fmt.Println("what is the customer's name?")
	_, err := fmt.Scanf("%s", &c.customersName)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println("what did the user buy?")
	_, err2 := fmt.Scanf("%s", &item)
	if err2 != nil {
		fmt.Println("Error: ", err2)
	}
	c.itemBought = append(c.itemBought, item)

	fmt.Println("How many pieces?")
	_, err3 := fmt.Scanf("%d", &numberOfPieces)
	if err3 != nil {
		fmt.Println("Error: ", err3)
	}
	for numberOfPieces <= 0 {
		fmt.Println("You've not entered the number of pieces the user got.")
		fmt.Print("How many pieces?")
		_, err3 := fmt.Scanf("%d", &numberOfPieces)
		if err3 != nil {
			fmt.Println("Error: ", err3)
		}
	}
	c.itemQuantity = append(c.itemQuantity, numberOfPieces)

	fmt.Println("How much per unit?")
	_, err4 := fmt.Scanf("%2f", &pricePerUnit)
	if err4 != nil {
		fmt.Println("Error: ", err4)
	}
	for pricePerUnit <= 0 {
		fmt.Println("Price is invalid, please enter a valid price! ")
		fmt.Print("How much per unit?: ")
		_, err4 := fmt.Scanf("%2f", &pricePerUnit)
		if err4 != nil {
			fmt.Println("Error: ", err4)
		}
	}
	c.itemPrice = append(c.itemPrice, pricePerUnit)

	fmt.Println("Add more items?")
	var response string
	_, err5 := fmt.Scanf("%s", &response)
	if err5 != nil {
		fmt.Println("Error: ", err5)
	}
	for response == "yes" {

	}
}
