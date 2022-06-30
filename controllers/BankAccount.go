package class

import "fmt"

type BankAccount struct {
	Agency  int
	Account int
	Holder  string
	Balance float64
}

func (acc *BankAccount) getGeneralInfo() {
	fmt.Println("[Agency Number]:", acc.Agency)
	fmt.Println("[Account]:", acc.Account)
	fmt.Println("[Holder]:", acc.Holder)
	fmt.Print("[Balance]: R$", acc.Balance)
}

func (acc *BankAccount) Deposit(value float64) bool {
	var canDeposit bool = value > 0
	if !canDeposit {
		fmt.Println("Something unexpected happened.")
		return false
	}
	acc.Balance += value
	fmt.Println("Deposit accomplished successfully!")
	return true
}

func (acc *BankAccount) Withdraw(value float64) bool {
	var canWithdraw bool = value > 0 && value < acc.Balance
	if !canWithdraw {
		fmt.Println("Insufficient balance.")
		return false
	}
	acc.Balance -= value
	fmt.Println("Withdraw accomplished successfully!")
	return true
}

func (acc *BankAccount) Transfer(value float64, accToTransfer BankAccount) bool {
	if acc.Withdraw(value) {
		accToTransfer.Deposit(value)
		fmt.Println("Transfer accomplished successfully!")
		return true
	}
	fmt.Println("Transfer failed.")
	return false
}
