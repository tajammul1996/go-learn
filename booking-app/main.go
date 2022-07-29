package main

import (
	"booking-app/helper"
	"fmt"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	var bookings []string

	fmt.Println("Welcome to our", conferenceName, " booking application")
	fmt.Printf("Remaining tickets %v/%v \n", remainingTickets, conferenceTickets)
	fmt.Println("Get your tickets here to attend")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Print("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Print("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Print("Enter your email address: ")
		fmt.Scan(&email)

		fmt.Print("Enter the number of tickets to book: ")
		fmt.Scan(&userTickets)

		//bookings[0] = firstName + " " + lastName
		//fmt.Println(bookings, len(bookings))

		if remainingTickets >= userTickets {
			fmt.Printf("--------------------------------------\n\n Thank you %v %v for booking %v tickets. You will receive a confirmation on %v \n\n --------------------------------------\n\n", firstName, lastName, userTickets, email)
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)
			fmt.Printf("%v tickets remaining for the %v\n\n", remainingTickets, conferenceName)
			//return
		} else {
			fmt.Printf("\n====================\nSorry! Available tickets are only %v \n====================\n\n", remainingTickets)
		}

		helper.PrintFirstNames(&bookings)

		if remainingTickets == 0 {
			fmt.Println("=================================\n Our conference is fully booked! \n=================================")
			break
		}
	}

	//remainingTickets = remainingTickets - userTickets

}
