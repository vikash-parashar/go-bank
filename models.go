package main

type BankAccount struct {
	Id            int    `json:"account-id"`
	FirstName     string `json:"first-name"`
	LastName      string `json:"last-name"`
	Age           string `json:"age"`
	Phone         string `json:"phone"`
	AccountNumber string `json:"account-number"`
	Balance       string `json:"account-balance"`
}

func NewBankAccount(first, last string) *BankAccount {
	return &BankAccount{
		FirstName: first,
		LastName:  last,
	}
}
