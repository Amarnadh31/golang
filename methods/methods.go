package methods

import "fmt"

type Bank interface {
	Deposit( amount int)
	Withdraw(amount int)
	CheckBalance()
}

type BankAccount struct{
	Balance int
	Account []int
}

// type WithdrawDetails struct{
// 	DepositDetails
// 	Balance int
// }




func (d *BankAccount) Deposit(amount int){
	
	d.Balance +=amount
	fmt.Println("Deposited amount : ", amount)
	
	fmt.Println("Available amount Rs: ", d.Balance)
}
func (w *BankAccount) Withdraw(amount int){

	if amount > w.Balance {
		fmt.Println("Insufficient balance, available balance is :", w.Balance)
	}else {
		w.Balance -= amount

		fmt.Println("withdrawl amount Rs: ",amount )
		fmt.Println("Available amount Rs: ", w.Balance)
	}

	
}

func (c *BankAccount) CheckBalance(){

	fmt.Println("current available balance is :", c.Balance)
}



