package wallet

import (
	"errors"
	"testing"
)

func TestWithdraw(t *testing.T) {
	t.Run("Empty wallet case", func(t *testing.T) {
		wallet := Wallet{}
		err := wallet.Withdraw(Bitcoin(10))

		AssertError(t, err, ErrInsufficientFunds)
		checkBalance(t, wallet, Bitcoin(0))
	})

	t.Run("Non-empty wallet", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))

		AssertError(t, err, nil)
		checkBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw more than balance", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(30))

		AssertError(t, err, ErrInsufficientFunds)
		checkBalance(t, wallet, Bitcoin(20))
	})

	t.Run("WIthdraw 0", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(0))

		AssertError(t, err, nil)
		checkBalance(t, wallet, Bitcoin(20))
	})

	t.Run("Withdraw negative amount", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(-20))

		AssertError(t, err, ErrNegativeInput)
		checkBalance(t, wallet, Bitcoin(20))
	})

	t.Run("Withdraw whole balance", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(wallet.balance)

		AssertError(t, err, nil)
		checkBalance(t, wallet, Bitcoin(0))
	})

}

func AssertError(t *testing.T, got, expected error) {
	t.Helper()

	if !errors.Is(got, expected) {
		t.Errorf("got error '%v', expected error '%v'", got, expected)

	}

}
