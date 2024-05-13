package main

import (
	"fmt"
	"strings"
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

	var response string

	for {
		fmt.Println("Add more items?")
		_, err5 := fmt.Scanf("%s", &response)
		if err5 != nil {
			fmt.Println("Error: ", err5)
			break
		}
		if strings.ToLower(response) != "no" {
			break
		}
	}

	fmt.Println("What is your name?")
	_, err6 := fmt.Scanf("%s", &c.cashiersName)
	if err6 != nil {
		fmt.Println("Error: ", err6)
	}
}

func (c *CheckOutApp) displayOutput() {
	fmt.Println("=====================================================")
	fmt.Println("\t\tITEM\tQTY\tPRICE\t  TOTAL(NGN)")
	fmt.Println("-----------------------------------------------------")

}
