package methods

import (
	"fmt"
)

type BankAccount struct {
	AccountNumber int
	HolderName    string
	Balance       float64
}

func (b *BankAccount) Passbook() {

	var userInput uint
	fmt.Println("Enter your name: ")
	fmt.Scan(&b.HolderName)

	fmt.Println("Enter your Account Number:")
	fmt.Scan(&b.AccountNumber)
	for {

		fmt.Print("1. Deposit\n2. Withdraw \n3. Exit \n")

		fmt.Println("Select one service by entering one number: ")
		fmt.Scan(&userInput)

		switch userInput {

		case 1:

			var depositAmount float64

			fmt.Println("Enter Deposit Amount: ")
			fmt.Scan(&depositAmount)
			b.Balance += depositAmount
			fmt.Println("Your total amount is Rs.", b.Balance)

		case 2:
			var withdrawAmount float64

			fmt.Println("Enter withdrawl Amount: ")
			fmt.Scan(&withdrawAmount)

			if withdrawAmount < b.Balance {
				b.Balance -= withdrawAmount
				fmt.Println("Your remaining balance is Rs.", b.Balance)
			} else {
				fmt.Println("Insufficient Funds")
			}

		case 3:
			fmt.Println("Thank you..Exiting")
			return

		default:
			fmt.Println("Invalid Input!, Please try again")

		}

	}

}
