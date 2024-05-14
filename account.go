package main

import "fmt"

type Account struct {
	Name         string
	Balance      float64
	Type         string
	Currency     string
	Transactions []Transaction
}

func NewAccount(name string, balance float64, accountType string) Account {
	return Account{
		Name:     name,
		Balance:  balance,
		Type:     accountType,
		Currency: "USD", // Default currency
	}
}

func (a *Account) Deposit(amount float64) {
	a.Balance += amount
}

func (a *Account) Withdraw(amount float64) {
	if amount > a.Balance {
		// Handle insufficient funds
		fmt.Println("Insufficient funds!")
		return
	}
	a.Balance -= amount
}

func (a *Account) AddTransaction(transaction Transaction) {
	a.Transactions = append(a.Transactions, transaction)
}

func (a *Account) DisplayTransactions() {
	fmt.Printf("Transactions for account '%s':\n", a.Name)
	for _, transaction := range a.Transactions {
		fmt.Printf("Type: %s, Amount: %.2f\n", transaction.Type, transaction.Amount)
	}
}
