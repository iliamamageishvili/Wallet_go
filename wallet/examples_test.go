package wallet

import (
	"fmt"
)

func Example() {
	myWallet := Wallet{}

	myWallet.Deposit(10.5)

	if err := myWallet.Withdraw(1.0); err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(myWallet.Balance())
	// Output: 9.5
}
