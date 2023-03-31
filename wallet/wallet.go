package wallet

import "errors"

type Bitcoin float64

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) error {
	if amount < 0 {
		return ErrNegativeInput
	}
	w.balance += amount
	return nil
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")
var ErrNegativeInput = errors.New("negative inputs are not allowed")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	if amount < 0 {
		return ErrNegativeInput
	}
	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
