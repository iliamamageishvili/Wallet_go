package wallet

import "testing"

func TestBalance(t *testing.T) {
	t.Run("Empty wallet case", func(t *testing.T) {
		wallet := Wallet{}
		got := wallet.Balance()
		expected := Bitcoin(0)

		assertBalance(t, got, expected)
	})

	t.Run("Non-empty wallet case", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}
		got := wallet.Balance()
		expected := Bitcoin(10)

		assertBalance(t, got, expected)
	})

	t.Run("Negative wallet case", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(-10)}
		got := wallet.Balance()
		expected := Bitcoin(-10)

		assertBalance(t, got, expected)
	})
}

func assertBalance(t *testing.T, got Bitcoin, expected Bitcoin) {
	t.Helper()

	if got != expected {
		t.Errorf("expected %v but got %v", expected, got)
	}
}
