package main

import (
	"math/rand"
)

var (
	Credit  string = "Credit"
	Debit   string = "Debit"
	Admin   string = "Admin"
	General string = "user"
)

type Account struct {
	ID            int    `json:"id"`
	FirstName     string `json:"first-name"`
	LastName      string `json:"last-name"`
	AccountNumber int64  `json:"account-number"`
	Balance       int64  `json:"account-balance"`
}

func NewAccount(first, last string) *Account {
	return &Account{
		ID:            rand.Intn(1000000),
		FirstName:     first,
		LastName:      last,
		AccountNumber: int64(rand.Intn(1000000000)),
	}
}
