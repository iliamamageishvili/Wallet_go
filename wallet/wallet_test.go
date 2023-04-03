package wallet

import (
	"sync"
	"testing"
)

func TestConcurrentSafe(t *testing.T) {
	w := Wallet{}

	var wg sync.WaitGroup
	n := 10
	wg.Add(n * 2)

	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()

			for j := 0; j < 100; j++ {
				err := w.Deposit(Bitcoin(1))
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
			}
		}()
	}

	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()

			for j := 0; j < 50; j++ {
				err := w.Withdraw(Bitcoin(1))
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
			}
		}()
	}

	wg.Wait()

	balance := w.Balance()
	if balance != Bitcoin(500) {
		t.Errorf("Unexpected balance: %v", balance)
	}
}
