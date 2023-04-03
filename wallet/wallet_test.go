package wallet

import (
	"sync"
	"testing"
)

func TestConcurrentSafe(t *testing.T) {

	w := Wallet{}

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for j := 0; j < 100; j++ {
				err := w.Deposit(Bitcoin(1))
				if err != nil {
					t.Errorf("Wrong error: %v", err)
				}
			}

			for j := 0; j < 50; j++ {
				err := w.Withdraw(Bitcoin(1))
				if err != nil {
					t.Errorf("Wrong error: %v", err)
				}
			}
		}()
	}

	wg.Wait()

	balance := w.Balance()
	if balance != Bitcoin(500) {
		t.Errorf("Wrong balance: %v", balance)
	}
}
