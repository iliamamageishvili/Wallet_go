package main

import (
	"Wallet_go/wallet"
	"fmt"
)

func main() {
	myWallet := wallet.Wallet{}

	myWallet.Deposit(10.5)
	fmt.Println(myWallet.Balance())

	err := myWallet.Withdraw(1.0)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(myWallet.Balance())
}
