package wallet

import "testing"

func TestWithdraw(t *testing.T) {
	t.Run("Empty wallet case", func(t *testing.T) {
		wallet := Wallet{}
		err := wallet.Withdraw(Bitcoin(10))

		withdrawError(t, err, ErrInsufficientFunds)
		checkBalance(t, wallet, Bitcoin(0))
	})

	t.Run("Non-empty wallet", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))

		withdrawError(t, err, nil)
		checkBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw more than balance", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(30))

		withdrawError(t, err, ErrInsufficientFunds)
		checkBalance(t, wallet, Bitcoin(20))
	})
}

func withdrawError(t *testing.T, got error, expected error) {
	t.Helper()

	if got == nil && expected == nil {
		return
	}

	if got == nil || expected == nil {
		t.Fatalf("got error '%v', expected error '%v'", got, expected)
	}

	if got.Error() != expected.Error() {
		t.Errorf("got error '%v', expected error '%v'", got, expected)
	}
}