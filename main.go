package main

import (
	// "fmt"
	"golang/golang/BookingApp"
	// "fmt"
	// "golang/golang/structs-methods"
	// "golang/golang/functions"
	// "golang/golang/loops" // importing loops func to main
)

func main(){

	bookingapp.GreetUser()

	userInputs := bookingapp.UserDetails{}

	userInputs.GettingUser()


//////////////structs&methods///////////////////////
	// account := methods.BankAccount{}
	// account.Passbook()
	////////////////loops/////////////////////
	// loops.LoopsFunc() //calling the loop func
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
	// 	break
	// }
/////////////////////////functions//////////////////////
	// normaltemp := functions.Weather(temp)

	// fmt.Println("Farenheit Temp : ", normaltemp)

}