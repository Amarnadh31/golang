package main

import (
	// "fmt"
	"fmt"
	"golang/golang/methods"
	// "golang/golang/filecopy"
	// "golang/golang/BookingApp"
	// "fmt"
	// "golang/golang/structs-methods"
	// "golang/golang/functions"
	// "golang/golang/loops" // importing loops func to main
)

func main(){

	// filecopy.Copyfile() // calling Copyfile function from filecopy package

	// bookingapp.GreetUser() // calling Greetuser function from BookingApp package

	// userInputs := bookingapp.UserDetails{} //calling UserDetails struct from BookingApp package ad storing in userInputs

	// userInputs.GettingUserDetails() //Calling GettingUserDetails function from BookingApp package

	// for i := 0; i < 11 ; i++{
	// 	fmt.Print(functions.Fibonacci(i), " ")
	// }
//////////////structs&methods///////////////////////
	// account := methods.BankAccount{}
	// account.Passbook()
	////////////////loops/////////////////////
	// loops.LoopsFunc() //calling the loop func

	/////////////////////////functions//////////////////////
	// var temp float64
	
	// for {
	// 	fmt.Println("Enter the temperature: ")
	// 	_, err := fmt.Scan(&temp)
	// 	if err != nil {
	// 		fmt.Println("Invalid Input!, Please enter numbers.")

	// 		var discard string
	// 		fmt.Scanln(&discard)
	// 		continue
	// 	}

	// 	normaltemp := functions.Weather(temp)
	// 	if temp < 0{
	// 		normaltemp = 0
	// 	}

	// 	fmt.Println("Farenheit Temp : ", normaltemp)
	// }

	///////////////////////////////////////////interfaces///////////////////
	// account := methods.BankAccount{Balance: 60000}
		var account methods.Bank = &methods.BankAccount{Balance: 6000}
		

			var deposit int
			var widraw int
			var choose int

	fmt.Println("Please select option \n1. Deposit\n2. Withdraw\n3. Bank Balance\n4. Exit")
	fmt.Scan(&choose)

	for {
		
	switch choose{

		case 1: 
			fmt.Println("Please Enter Deposit mount :")
			fmt.Scan(&deposit)
			account.Deposit(deposit)
			return
		case 2:
			fmt.Println("Please Enter withdrawl mount :")
			fmt.Scan(&widraw)
			account.Withdraw(widraw)
			return
		case 3:

			account.CheckBalance()
			return

		case 4:
			fmt.Println("Thank you for visiting")
			return
		default:
			fmt.Println("Invalid input. Kindly enter 1,2,3 and 4 options")
			return
		
	}
		
	}

	

}