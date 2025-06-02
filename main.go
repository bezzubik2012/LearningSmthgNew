package main

import (
	"LearningSmthgNew/account"
	"fmt"
)

func main() {
	login := promptData("Enter your login: ")
	password := promptData("Enter your password: ")
	url := promptData("Enter your URL: ")

	user1, err := account.NewAccount(login, password, url)
	if err != nil {
		println("Invalid parameters")
		return
	}
	
	user1.ViewAccData()
}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var data string
	fmt.Scanln(&data)
	return data
}
