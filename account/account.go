package account

import (
	"errors"
	"fmt"
	"math/rand/v2"
	url2 "net/url"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_-")

type Account struct {
	login    string
	password string
	url      string
}

func (user *Account) ViewAccData() {
	fmt.Println(user.login, user.password, user.url)
}

func (user *Account) generatePassword(lengthPass int) {
	password := make([]rune, lengthPass)
	for i := range password {
		password[i] = letters[rand.IntN(len(letters))]
	}
	user.password = string(password)
}

func NewAccount(login, password, url string) (*Account, error) {
	newAcc := Account{
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
