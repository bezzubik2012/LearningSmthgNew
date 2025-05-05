package main

import (
	"fmt"
	"strconv"
)

var transactions []float64

func main() {
	fmt.Println("Hello, lets calculate your transactions!\nType positive or negative value of transaction, with point if its needed:\nAfter last input enter empty string")
	for {
		inputTransaction()
		fmt.Println("Do you want to repeat calculations?")
		if repeatInput() == true {
			continue
		}
		break
	}
	fmt.Println(transactions)

}

func inputTransaction() {
	var transaction string
	for {
		fmt.Println("Enter your transaction (n for exit): ")
		fmt.Scanln(&transaction)
		if transaction == "n" || transaction == "N" {
			fmt.Println("End of input")
			break
		} else if number, err := strconv.ParseFloat(transaction, 64); err == nil {
			transactions = append(transactions, number)
		} else {
			fmt.Println("Invalid input, please try again")
			continue
		}
	}
}

func repeatInput() bool {
	var input string
	fmt.Scanln(&input)
	if input == "y" || input == "Y" {
		return true
	} else {
		return false
	}
}
