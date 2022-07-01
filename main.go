package main

import (
	ctrl "bank/controllers"
	"fmt"
)

func main() {
	fmt.Println("----------------------------------------------")

	var holder1 = ctrl.Client{}
	holder1.Name = "Thales"
	holder1.LastName = "Lira"
	holder1.CPF = "111.111.111-11"
	holder1.Ocupation = "Computer Scientist"
	holder1.Role = "Software Engineer"

	var holder2 = ctrl.Client{}
	holder2.Name = "Maria"
	holder2.LastName = "Vitória"
	holder2.CPF = "222.222.222-22"
	holder2.Ocupation = "Computer Scientist"
	holder2.Role = "Database Administrator"

	fmt.Println("\nCreating a new obj BankAccount:", holder1.CompletedName())
	var acc1 = ctrl.BankAccount{}
	acc1.Holder = holder1
	acc1.Agency = 1
	acc1.Account = 1
	fmt.Println("New BankAccount Created:")
	acc1.GeneralInfo()

	fmt.Println("\nCreating a new obj BankAccount:", holder2.CompletedName())
	var acc2 = ctrl.BankAccount{}
	acc2.Holder = holder2
	acc2.Agency = 1
	acc2.Account = 2
	fmt.Println("New BankAccount Created:")
	acc2.GeneralInfo()

	fmt.Println("\nDepositing R$500.00 to acc1:", acc1.Holder.CompletedName())
	acc1.Deposit(500)
	fmt.Println("Updated acc1:")
	acc1.GeneralInfo()

	fmt.Println("\nDepositing R$1000.00 to acc2:", acc2.Holder.CompletedName())
	acc2.Deposit(1000)
	fmt.Println("Updated acc2:")
	acc2.GeneralInfo()

	acc1.Transfer(200, &acc2)
	acc1.GeneralInfo()
	acc2.GeneralInfo()
}
