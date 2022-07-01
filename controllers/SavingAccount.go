package controllers

import "fmt"

type SavingAccount struct {
	Agency  int
	Account int
	Holder  Client
	balance float64
}

func (acc *SavingAccount) GeneralInfo() {
	fmt.Println("[Agency Number]:", acc.Agency)
	fmt.Println("[Account]:", acc.Account)
	fmt.Println("[Holder]:", acc.Holder.CompletedName())
	fmt.Println("[Balance]: R$", acc.balance)
	fmt.Println()
}

func (acc *SavingAccount) Deposit(value float64) bool {
	var canDeposit bool = value > 0
	if !canDeposit {
		fmt.Println("Something unexpected happened.")
		return false
	}
	acc.balance += value
	fmt.Println("Deposit accomplished successfully!")
	return true
}

func (acc *SavingAccount) Withdraw(value float64) bool {
	var canWithdraw bool = value > 0 && value < acc.balance
	if !canWithdraw {
		fmt.Println("Insufficient balance.")
		return false
	}
	acc.balance -= value
	fmt.Println("Withdraw accomplished successfully!")
	return true
}

func (acc *SavingAccount) Transfer(value float64, accToTransfer *SavingAccount) bool {
	if acc.Withdraw(value) {
		accToTransfer.Deposit(value)
		fmt.Println("Transfer accomplished successfully!")
		return true
	}
	fmt.Println("Transfer failed.")
	return false
}
