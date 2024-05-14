package main

import "time"

type Transaction struct {
	Type     string
	Amount   float64
	Account  *Account
	Date     time.Time
	Category string
	Notes    string
}

func NewTransaction(transactionType string, amount float64, account *Account, category string, notes string) Transaction {
	return Transaction{
		Type:     transactionType,
		Amount:   amount,
		Account:  account,
		Date:     time.Now(),
		Category: category,
		Notes:    notes,
	}
}
