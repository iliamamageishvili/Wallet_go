package wallet

import "testing"

func TestBalance(t *testing.T) {
	t.Run("Empty wallet case", func(t *testing.T) {
		wallet := Wallet{}
		expected := Bitcoin(0)

		checkBalance(t, wallet, expected)
	})

	t.Run("Non-empty wallet case", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}
		expected := Bitcoin(10)

		checkBalance(t, wallet, expected)
	})

	t.Run("Negative wallet case", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(-10)}
		expected := Bitcoin(-10)

		checkBalance(t, wallet, expected)
	})
}
