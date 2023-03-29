package main

import (
	"Wallet_go/wallet"
	"fmt"
)

func main() {
	wallet := wallet.Wallet{}

	wallet.Deposit(10.5)
	fmt.Println(wallet.Balance())

	err := wallet.Withdraw(1.0)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(wallet.Balance())
}
