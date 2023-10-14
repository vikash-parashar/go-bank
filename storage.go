package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Storage interface {
	GetAccountByID(int) (*Account, error)
	CreateAccount(*Account) error
	UpdateAccount(*Account) error
	DeleteAccount(int) error
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore(dsn string) (*PostgresStore, error) {
	if dsn == "" {
		dsn = "user=postgres password=postgres dbname=go bank sslmode=disable"
	}
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	return &PostgresStore{db: db}, nil
}

func (s *PostgresStore) init() error {
	return s.CreateAccountTable()
}
func (s *PostgresStore) CreateAccountTable() error {
	query := `create table account if not exists(
		id serial primary key,
		first_name varchar(50),
		last_name varchar(50),
		number serial,
		balance serial,
		created_at timestamp
	)`
	res, err := s.db.Exec(query)
	if err != nil {
		return err
	}
	log.Println(res)
	return nil
}

func (s *PostgresStore) GetAccountByID(id int) (*Account, error) {
	return nil, nil
}

func (s *PostgresStore) CreateAccount(a *Account) error {
	return nil
}

func (s *PostgresStore) UpdateAccount(a *Account) error {
	return nil
}

func (s *PostgresStore) DeleteAccount(id int) error {
	return nil
}
