package main

import (
	"fmt"
	"math"
)

func main() {
	const IMTPow = 2
	var userHeight float64
	var userWeight float64
	fmt.Print("Type your height in cm: ")
	fmt.Scanln(&userHeight)
	fmt.Print("Type your weight in kg: ")
	fmt.Scanln(&userWeight)
	IMT := userWeight / math.Pow(userHeight/100, IMTPow)
	fmt.Printf("Your IMT is %.1f", IMT)

}
