package main

import (
	"fmt"
	"math"
)

const IMTPow = 2

var userHeight float64
var userWeight float64
var answer string

func main() {
	fmt.Println("Welcome to IMT calculator! Let`s begin!")
	for {
		getUserInput()
		IMT := calculateIMT(userWeight, userHeight, IMTPow)
		outputResult(IMT)
		if repeatCalculation() {
			break
		}
	}

}

func outputResult(IMT float64) {
	switch {
	case IMT < 16:
		fmt.Println("Too low weight")
	case (16 <= IMT) && (IMT < 18.5):
		fmt.Println("Low weight")
	case (18.5 <= IMT) && (IMT < 25):
		fmt.Println("Normal weight")
	case (25 <= IMT) && (IMT < 30):
		fmt.Println("High weight")
	default:
		fmt.Println("Too high weight")
	}

	fmt.Printf("Your IMT is %.1f\n", IMT)
}

func calculateIMT(userKG, userHeight, pow float64) float64 {
	return userKG / math.Pow(userHeight/100, pow)
}

func getUserInput() {
	fmt.Print("Type your height in cm: ")
	fmt.Scanln(&userHeight)
	fmt.Print("Type your weight in kg: ")
	fmt.Scanln(&userWeight)
}

func repeatCalculation() bool {
	fmt.Println("Do you want to continue IMT calculations? y/n")
	fmt.Scanln(&answer)
	if answer != "y" && answer != "Y" || answer == "" {
		fmt.Println("Exiting...")
		return true
	}
	return false
}
