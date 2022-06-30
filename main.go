package main

import (
	ctrl "bank/controllers"
	"fmt"
)

func main() {
	fmt.Println("Creating a new obj BankAccount: Thales Lira")
	var acc1 = ctrl.BankAccount{}
	acc1.Holder = "Thales Lira"
	acc1.Agency = 1
	acc1.Account = 1
	fmt.Println("New BankAccount Created:", acc1)

	fmt.Println("\nCreating a new obj BankAccount: Maria Vitória")
	var acc2 = ctrl.BankAccount{}
	acc2.Holder = "Maria Vitória"
	acc2.Agency = 1
	acc2.Account = 2
	fmt.Println("New BankAccount Created:", acc2)

	fmt.Println("\nDepositing R$350.00 to acc1:", acc1.Holder)
	acc1.Deposit(350)
	fmt.Println("Updated acc1:", acc1)

	fmt.Println("\nDepositing R$735.45 to acc2:", acc2.Holder)
	acc2.Deposit(735.45)
	fmt.Println("Updated acc2:", acc2)

	acc1.Transfer(450, acc2)
}
