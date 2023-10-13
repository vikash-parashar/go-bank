package main

import (
	"database/sql"
	"time"

	uuid "github.com/google/uuid"
)

var (
	Credit  string = "Credit"
	Debit   string = "Debit"
	Admin   string = "Admin"
	General string = "user"
)

type (
	APIServer struct {
		Version string
		Env     string
		Port    string
		DB      *sql.DB
	}

	APIError struct {
		Error string
	}
)

type BankAccount struct {
	ID            uuid.UUID     `json:"account_id"`
	AccountNumber string        `json:"account_number"`
	Balance       string        `json:"account_balance"`
	User          User          `json:"user_details"`
	Transactions  []Transaction `json:"transactions"`
	CreatedAt     time.Time     `json:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at"`
}

func NewBankAccount(u *User) *BankAccount {
	uuid.New()
	return &BankAccount{
		User: *u,
	}
}

type User struct {
	ID               uuid.UUID
	FirstName        string    `json:"first_name"`
	LastName         string    `json:"last_name"`
	PanCardNumber    string    `json:"pan_card"`
	AadharCardNumber string    `json:"aadhar_card"`
	DateOfBirth      time.Time `json:"date_of_birth"`
	Phone            string    `json:"phone"`
	Email            string    `json:"email"`
	Type             string    `json:"user_type"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type Transaction struct {
	ID          uuid.UUID `json:"transaction_id"`
	TxnType     string    `json:"transaction_type"`
	Amount      float64   `json:"transaction_amount"`
	Source      string    `json:"transaction_from"`
	Destination string    `json:"transaction_to"`
	Status      string    `json:"transaction_status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
