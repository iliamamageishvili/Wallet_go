package wallet

import (
	"errors"
	"sync"
)

type Bitcoin float64

type Wallet struct {
	balance Bitcoin
	mu      sync.Mutex
}

func (w *Wallet) Deposit(amount Bitcoin) error {
	if amount < 0 {
		return ErrNegativeInput
	}
	w.mu.Lock()
	defer w.mu.Unlock()
	w.balance += amount
	return nil
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")
var ErrNegativeInput = errors.New("negative inputs are not allowed")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount < 0 {
		return ErrNegativeInput
	}
	w.mu.Lock()
	defer w.mu.Unlock()
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.balance
}
