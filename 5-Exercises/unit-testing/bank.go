package main

type BankAccount interface {
	Deposit(amount int)
	Balance() int
}

func Deposit(account BankAccount, amount int) int {
	account.Deposit(amount)
	return account.Balance()
}
