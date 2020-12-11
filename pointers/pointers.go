package pointers

import (
	"errors"
	"fmt"
)

// ErrInsufficientFunds is Error for insufficient funds
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// Bitcoin is the currency in which we will be dealing
type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet is a structure for holding money
type Wallet struct {
	balance Bitcoin
}

// Balance is used to return the balance of the wallet
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// Deposit is used to deposit some amount into the bank
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Withdraw is used to withdraw bitcoin from the bank
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}
