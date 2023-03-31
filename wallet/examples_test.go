package wallet

import (
	"fmt"
)

var myWallet Wallet = Wallet{}

func ExampleBalance() {
	fmt.Println(myWallet.balance)
	// Output: 0
}

func ExampleDeposit() {
	myWallet.Deposit(10.5)
	fmt.Println(myWallet.Balance())
	// Output: 10.5
}

func ExamlpeWithdraw() {
	err := myWallet.Withdraw(1.0)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(myWallet.Balance())
	// Output: 9.5
}
