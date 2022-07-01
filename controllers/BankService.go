package controllers

import "fmt"

type BankBillet interface {
	Withdraw(value float64) bool
}

func PayBillet(value float64, acc BankBillet) {
	acc.Withdraw(value)
	fmt.Println("Billet paid.")
}
