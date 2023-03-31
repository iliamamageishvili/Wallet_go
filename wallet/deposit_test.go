package wallet

import "testing"

func TestDeposit(t *testing.T) {
	t.Run("Empty wallet case", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		checkBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Non-empty wallet case", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}
		wallet.Deposit(Bitcoin(20))

		checkBalance(t, wallet, Bitcoin(30))
	})

	t.Run("Negative wallet case", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(-10)}
		wallet.Deposit(Bitcoin(20))

		checkBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Depositing zero", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}
		wallet.Deposit(Bitcoin(0))

		checkBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Depositing negative value", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}
		wallet.Deposit(Bitcoin(-5))

		checkBalance(t, wallet, Bitcoin(5))
	})

}

func checkBalance(t *testing.T, wallet Wallet, expected Bitcoin) {
	t.Helper()

	if wallet.balance != expected {
		t.Errorf("expected %v but got %v", expected, wallet.balance)
	}
}
