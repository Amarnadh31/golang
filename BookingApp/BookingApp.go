package bookingapp

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

type UserDetails struct {
	firstName string
	lastName string
	email string
	city string
	bookings uint
}

var bookings []UserDetails

func GreetUser(){
	fmt.Println("###########Welcome to our Booking service############")
}

func (u *UserDetails)GettingUser(){

	const totalTicktes uint = 100
	availableTickets := totalTicktes

	
	
	for{
		for {
			fmt.Println("Please enter your first name:")
			fmt.Scan(&u.firstName)
			alphaRegex := regexp.MustCompile(`^[a-zA-Z]+$`)

			if alphaRegex.MatchString(u.firstName){
				break
			}else{
				fmt.Println("Invalid Input!, Please Enter Alphabets only")
			}

		}

		for {
			fmt.Println("Please enter your Last name:")
			fmt.Scan(&u.lastName)
			alphaRegex := regexp.MustCompile(`^[a-zA-Z]+$`)

			if alphaRegex.MatchString(u.lastName){
				break
			}else{
				fmt.Println("Invalid Input!, Please Enter Alphabets only")
			}

		}

		for {
			fmt.Println("Please enter your email address(Must contain @ and .com)")
			fmt.Scan(&u.email)

			if strings.Contains(u.email, "@") && strings.Contains(u.email, ".com"){
					break
			}else{
				fmt.Println("Invalid Inpu!, Please your email must contain @ and .com")
			}
		}

		for {

			fmt.Println("Please enter your City name:")
			fmt.Scan(&u.city)
			alphaRegex := regexp.MustCompile(`^[a-zA-Z]+$`)

			if alphaRegex.MatchString(u.city){
				break
			}else{
				fmt.Println("Invalid Input!, Please Enter Alphabets only")
			}
			
		}

		for{

			fmt.Println("Please enter no.of tickets: ")
			_, err := fmt.Scan(&u.bookings)
			if err != nil {
				fmt.Println("Please enter valid Input(1-100)")

				var discard string
				fmt.Scanln(&discard)
				continue
			} else if u.bookings > availableTickets || u.bookings < uint (0) {
					fmt.Println("Invalid input!, Check the available tickets please")
			}else {

				availableTickets -= u.bookings 

				fmt.Printf("Thank you!. You have booked %v ticktes \n", u.bookings)

				fmt.Printf("Available tickets are %v \n", availableTickets)
				break
		
			}

		}
	
		bookings = append(bookings, *u)
		// for _, b := range bookings {
		 	fmt.Println(bookings)
		// }

		fmt.Printf("We will shortly send you booking details to %v\n", u.email)



		
		go func (email string){
			time.Sleep(10 * time.Second)
			fmt.Printf("*************Cograts! your tickets are successfully sent to %v ******************\n", email)
		}(u.email)
		
		var confirmation string
		fmt.Println("Do you want to book more? (yes/no):")
		fmt.Scan(&confirmation)
		alphaRegex := regexp.MustCompile(`^[a-zA-Z]+$`)
		for {
			if alphaRegex.MatchString(confirmation){
				break
			}else{
				fmt.Println("Invalid Input!, Please Enter Alphabets only")
			}
		}
		if confirmation == "no"{
			return
		}
		for {
			if confirmation == "yes"{
				break
			}else{
				fmt.Println("Kindly enter either yes or no")
			}
		}
	
	}

}


// func sendingEmailBookingTickets() {

// 	emailDetails := UserDetails{}

// 	fmt.Printf("We will shortly send you booking details to %v", emailDetails.email)
// 	time.Sleep(10 * time.Second)

// 	fmt.Printf("Cograts! your tickets are successfully sent to %v", emailDetails.email)

// }

// func (t *UserDetails) ticketBooking(){

// 	const totalTicktes = 100
// 	for {

// 		fmt.Println("Please enter no.of tickets: ")
// 		fmt.Scan(&t.bookings)
// 	}

	
// }