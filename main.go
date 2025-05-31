package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	url2 "net/url"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_-")

type account struct {
	login    string
	password string
	url      string
}

func main() {
	login := promptData("Enter your login: ")
	password := promptData("Enter your password: ")
	url := promptData("Enter your URL: ")

	user1, err := newAccount(login, password, url)
	if err != nil {
		println("Invalid parameters")
		return
	}
	user1.generatePassword(50)
	user1.viewAccData()
}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var data string
	fmt.Scanln(&data)
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

func newAccount(login, password, url string) (*account, error) {
	newAcc := account{
		login:    login,
		password: password,
		url:      url}

	if login == "" {
		return nil, errors.New("empty login")
	}
	if password == "" {
		newAcc.generatePassword(50)
		return &newAcc, nil
	}
	_, err := url2.ParseRequestURI(url)
	if err != nil {
		return nil, errors.New("invalid URL")
	}
	return &newAcc, nil
}
