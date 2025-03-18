package main

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Wallet struct {
	// * if the variables starts with lowercase, it is private outside this package.
	balance Bitcoin
}

// * This interface can be implemented by any value that has a String method.
type Stringer interface {
	String() string
}

// * Here we use it to modify the "native" format of the value.
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
func (w *Wallet) Deposit(amount Bitcoin) {
	// * Struct Pointers are automatically dereferenced
	w.balance += amount
}
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}
