package main

import "testing"

type MockBankAccount struct {
	balance int
}

func (m *MockBankAccount) Deposit(amount int) {
	m.balance += amount
}

func (m *MockBankAccount) Balance() int {
	return m.balance
}

func TestDeposit(t *testing.T) {
	account := &MockBankAccount{}
	newBalance := Deposit(account, 100)

	want := 100
	if newBalance != want {
		t.Errorf("got %d, want %d", newBalance, want)
	}
}
