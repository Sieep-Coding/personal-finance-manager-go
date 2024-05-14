package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type FinanceManager struct {
	accounts     []Account
	transactions []Transaction
	categories   map[string][]string // Categories for transactions
}

func NewFinanceManager() *FinanceManager {
	return &FinanceManager{
		categories: make(map[string][]string),
	}
}

func (fm *FinanceManager) Run() {
	fm.loadCategories()
	for {
		fm.displayMenu()
		choice := fm.getUserChoice()
		switch choice {
		case 1:
			fm.createAccount()
		case 2:
			fm.listAccounts()
		case 3:
			fm.performTransaction()
		case 4:
			fm.displayTransactions()
		case 5:
			fm.searchTransactions()
		case 6:
			fm.displayAccountSummary()
		case 7:
			fm.manageCategories()
		case 8:
			fmt.Println("Exiting program...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func (fm *FinanceManager) displayMenu() {
	fmt.Println("\nPersonal Finance Manager")
	fmt.Println("1. Create Account")
	fmt.Println("2. List Accounts")
	fmt.Println("3. Perform Transaction")
	fmt.Println("4. Display All Transactions")
	fmt.Println("5. Search Transactions")
	fmt.Println("6. Display Account Summary")
	fmt.Println("7. Manage Categories")
	fmt.Println("8. Exit")
}

func (fm *FinanceManager) getUserChoice() int {
	var choice int
	fmt.Print("Enter your choice: ")
	fmt.Scanln(&choice)
	return choice
}

func (fm *FinanceManager) createAccount() {
	var name, accountType string
	var balance float64
	fmt.Print("Enter account name: ")
	fmt.Scanln(&name)
	fmt.Print("Enter account type (Checking, Savings, Credit Card, etc.): ")
	fmt.Scanln(&accountType)
	fmt.Print("Enter initial balance: ")
	fmt.Scanln(&balance)
	account := NewAccount(name, balance, accountType)
	fm.accounts = append(fm.accounts, account)
	fmt.Println("Account created successfully.")
}

func (fm *FinanceManager) listAccounts() {
	fmt.Println("\nList of Accounts:")
	for _, account := range fm.accounts {
		fmt.Printf("Name: %s, Type: %s, Balance: %.2f\n", account.Name, account.Type, account.Balance)
	}
}

func (fm *FinanceManager) performTransaction() {
	if len(fm.accounts) == 0 {
		fmt.Println("No accounts found. Please create an account first.")
		return
	}

	fmt.Println("\nSelect an account:")
	for i, account := range fm.accounts {
		fmt.Printf("%d. %s\n", i+1, account.Name)
	}
	var choice int
	fmt.Print("Enter your choice: ")
	fmt.Scanln(&choice)

	if choice < 1 || choice > len(fm.accounts) {
		fmt.Println("Invalid choice.")
		return
	}

	account := fm.accounts[choice-1]
	var transactionType string
	var amount float64
	fmt.Print("Enter transaction type (Deposit or Withdrawal): ")
	fmt.Scanln(&transactionType)
	fmt.Print("Enter amount: ")
	fmt.Scanln(&amount)

	var category, notes string
	fmt.Print("Enter transaction category: ")
	fmt.Scanln(&category)
	fmt.Print("Enter transaction notes (optional): ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		notes = scanner.Text()
	}

	if strings.EqualFold(transactionType, "Deposit") {
		account.Deposit(amount)
		transaction := NewTransaction("Deposit", amount, &account, category, notes)
		fm.transactions = append(fm.transactions, transaction)
		fmt.Println("Deposit successful.")
	} else if strings.EqualFold(transactionType, "Withdrawal") {
		if amount > account.Balance {
			fmt.Println("Insufficient funds.")
			return
		}
		account.Withdraw(amount)
		transaction := NewTransaction("Withdrawal", amount, &account, category, notes)
		fm.transactions = append(fm.transactions, transaction)
		fmt.Println("Withdrawal successful.")
	} else {
		fmt.Println("Invalid transaction type.")
	}
}

func (fm *FinanceManager) displayTransactions() {
	fmt.Println("\nAll Transactions:")
	for _, transaction := range fm.transactions {
		fmt.Printf("Account: %s, Type: %s, Amount: %.2f, Category: %s, Notes: %s\n",
			transaction.Account.Name, transaction.Type, transaction.Amount, transaction.Category, transaction.Notes)
	}
}

func (fm *FinanceManager) searchTransactions() {
	var searchType, searchText string
	fmt.Print("Enter search type (Account, Type, Amount, Category, or Notes): ")
	fmt.Scanln(&searchType)
	fmt.Print("Enter search text: ")
	fmt.Scanln(&searchText)

	fmt.Println("\nSearch Results:")
	for _, transaction := range fm.transactions {
		match := false
		switch strings.ToLower(searchType) {
		case "account":
			match = strings.Contains(strings.ToLower(transaction.Account.Name), strings.ToLower(searchText))
		case "type":
			match = strings.Contains(strings.ToLower(transaction.Type), strings.ToLower(searchText))
		case "amount":
			amount, _ := strconv.ParseFloat(searchText, 64)
			match = transaction.Amount == amount
		case "category":
			match = strings.Contains(strings.ToLower(transaction.Category), strings.ToLower(searchText))
		case "notes":
			match = strings.Contains(strings.ToLower(transaction.Notes), strings.ToLower(searchText))
		default:
			fmt.Println("Invalid search type.")
			return
		}

		if match {
			fmt.Printf("Account: %s, Type: %s, Amount: %.2f, Category: %s, Notes: %s\n",
				transaction.Account.Name, transaction.Type, transaction.Amount, transaction.Category, transaction.Notes)
		}
	}
}

func (fm *FinanceManager) displayAccountSummary() {
	if len(fm.accounts) == 0 {
		fmt.Println("No accounts found.")
		return
	}

	fmt.Println("\nAccount Summary:")
	for _, account := range fm.accounts {
		fmt.Printf("Account: %s\n", account.Name)
		fmt.Printf("Balance: %.2f\n", account.Balance)
		fmt.Println("Transactions:")
		var transactions []Transaction
		for _, transaction := range fm.transactions {
			if transaction.Account == &account {
				transactions = append(transactions, transaction)
			}
		}
		sort.Slice(transactions, func(i, j int) bool {
			return transactions[i].Amount > transactions[j].Amount
		})
		for _, transaction := range transactions {
			fmt.Printf("Type: %s, Amount: %.2f, Category: %s, Notes: %s\n", transaction.Type, transaction.Amount, transaction.Category, transaction.Notes)
		}
		fmt.Println()
	}
}

func (fm *FinanceManager) manageCategories() {
	// Category management implementation
}

func (fm *FinanceManager) loadCategories() {
	// Load categories from a file or set default categories
}
