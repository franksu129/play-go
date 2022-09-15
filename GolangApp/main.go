package main

import (
	"GolangApp/helper"
	"fmt"
	"sync"
	"time"
)

const tickets int = 50

var remainingTickets uint = 50
var appName = "Frank"
var bookings = make([]UserData, 0)

type UserData struct {
	name            string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	for {

		// 輸入資料
		userName, userEmail, userTickets := getUserInput()
		// 驗證
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(userName, userEmail, userTickets, remainingTickets)

		// 條件確認
		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTickets, userName, userEmail)
			go sendTicket(userTickets, userName, userEmail)

			names := getNames()
			fmt.Printf("The name of bookings are: %v\n", names)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email address you entered doesn't contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}
		}
	}
}

func greetUsers() {
	fmt.Println("Welcome ", appName, " app")
	fmt.Printf("We have total of %v tickets and %v are still available.\n", tickets, remainingTickets)
	fmt.Println("Get you ticket")
}

func getNames() []string {
	names := []string{}
	for _, booking := range bookings {
		names = append(names, booking.name)
	}

	return names
}

func getUserInput() (string, string, uint) {
	var userName string
	var userEmail string
	var userTickets uint

	// 輸入
	fmt.Print("Enter your name: ")
	fmt.Scan(&userName)

	fmt.Print("Enter your email : ")
	fmt.Scan(&userEmail)

	fmt.Print("Enter number of tickets : ")
	fmt.Scan(&userTickets)

	return userName, userEmail, userTickets
}

func bookTicket(userTickets uint, userName string, userEmail string) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		name:            userName,
		email:           userEmail,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)

	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you '%v' for booking '%v' tickets. You will receive a confirmation email at '%v'\n", userName, userTickets, userEmail)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, appName)
}

func sendTicket(userTickets uint, name string, email string) {
	time.Sleep(10 * time.Second)
	var v = fmt.Sprintf("%v tickets for %v", userTickets, name)
	fmt.Println("###############")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", v, email)
	fmt.Println("###############")
}
