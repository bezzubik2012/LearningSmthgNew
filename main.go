package main

import (
	"errors"
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
		err := getUserInput()
		if err != nil {
			fmt.Println(err)
			continue
		}
		IMT := calculateIMT(userWeight, userHeight, IMTPow)
		outputResult(IMT)
		userWeight = 0
		userHeight = 0
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
	IMT := userKG / math.Pow(userHeight/100, pow)
	return IMT
}

func getUserInput() error {
	if userHeight == 0 {
		fmt.Print("Type your height in cm: ")
		fmt.Scan(&userHeight)
		if userHeight <= 0 {
			return errors.New("Invalid height, please try again")
		}
	}

	if userWeight == 0 {
		fmt.Print("Type your weight in kg: ")
		fmt.Scan(&userWeight)
		if userWeight <= 0 {
			return errors.New("Invalid weight, please try again")
		}
	}

	return nil
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
