package main

import (
	"fmt"
	"math/rand/v2"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_-")

type account struct {
	login    string
	password string
	url      string
}

func main() {
	login := promptData("Enter your login: ")
	url := promptData("Enter your URL: ")

	user1 := account{
		login: login,
		url:   url,
	}
	user1.generatePassword(50)
	user1.viewAccData()
}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var data string
	fmt.Scan(&data)
	return data
}

func (user *account) viewAccData() {
	fmt.Println(user.login, user.password, user.url)
}

func (user *account) generatePassword(lengthPass int) {
	password := make([]rune, lengthPass)
	for i := range password {
		password[i] = letters[rand.IntN(len(letters))]
	}
	user.password = string(password)
}
