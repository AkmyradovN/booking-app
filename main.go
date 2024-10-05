package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

const conferenceTickets int = 50

var conferenceName string = "Go conference"
var remainingTickets uint = 50
var bookings = []UserData{}

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	firstName, lastName, emailAdress, userTickets := getUserInput()

	isValidName, isValidEmail, isValidUserTickets := helper.ValidateUserInput(firstName, lastName, emailAdress, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidUserTickets {

		bookTicket(userTickets, firstName, lastName, emailAdress)
		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, emailAdress)

		firstNames := getFirstNames()
		fmt.Printf("The first names of bookings are %v\n", firstNames)

		if remainingTickets == 0 {
			fmt.Printf("Our conference is booked out. Come back next year\n")
		}
	} else {
		if !isValidName {
			fmt.Println("Firstname or Last name you entered too short")
		}
		if !isValidEmail {
			fmt.Println("The email adress you entered is not valid email")
		}
		if !isValidUserTickets {
			fmt.Println("Number of tickets you entered is invalid")
		}
	}
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of, %v, tickts and,%v are still available\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend\n")
}

func getFirstNames() []string {
	firstNames := []string{}

	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var emailAdress string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your Last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your EmailAdress: ")
	fmt.Scan(&emailAdress)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, emailAdress, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, emailAdress string) {
	remainingTickets = remainingTickets - userTickets
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           emailAdress,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recive a confirmation email at %v \n", firstName, lastName, userTickets, emailAdress)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, emailAdress string) {
	time.Sleep(50 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v, %v", userTickets, firstName, lastName)
	fmt.Println("#############")
	fmt.Printf("Sending ticket %v to emailadress %v\n", ticket, emailAdress)
	fmt.Println("#############")
	wg.Done()
}
