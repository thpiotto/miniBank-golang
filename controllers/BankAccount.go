package class

import "fmt"

type BankAccount struct {
	Agency  int
	Account int
	Holder  string
	Balance float64
}

func (acc *BankAccount) Deposit(value float64) bool {
	var canDeposit bool = value > 0
	if !canDeposit {
		fmt.Println("Algo inesperado ocorreu")
		return false
	}
	acc.Balance += value
	fmt.Println("Deposito realizado com sucesso!")
	return true
}

func (acc *BankAccount) Withdraw(value float64) bool {
	var canWithdraw bool = value > 0 && value < acc.Balance
	if !canWithdraw {
		fmt.Println("Saldo insuficiente")
		return false
	}
	acc.Balance -= value
	fmt.Println("Saque realizado com sucesso!")
	return true
}

func (acc *BankAccount) Transfer(value float64, accToTransfer BankAccount) bool {
	if acc.Withdraw(value) {
		accToTransfer.Deposit(value)
		return true
	}

	return false
}
