package main

import (
	"GolangApp/helper"
	"fmt"
	"strings"
)

func main() {
	appName := "Frank"
	const tickets int = 50
	var remainingTickets uint = 50
	bookings := []string{}

	greetUsers(appName, tickets, remainingTickets)

	for {
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

		// 驗證
		isValidName := len(userName) >= 2
		isValidEmail := strings.Contains(userEmail, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		// 條件確認
		if isValidName && isValidEmail && isValidTicketNumber {

			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, userName+" "+userEmail)

			fmt.Printf("Thank you '%v' for booking '%v' tickets. You will receive a confirmation email at '%v'\n", userName, userTickets, userEmail)

			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, appName)

			printName(bookings)

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

		helper.PrintUserInfo(userName, userEmail, userTickets)
	}
}

func greetUsers(appName string, tickets int, remainingTickets uint) {
	fmt.Println("Welcome ", appName, " app")
	fmt.Printf("We have total of %v tickets and %v are still available.\n", tickets, remainingTickets)
	fmt.Println("Get you ticket")
}

func printName(bookings []string) {
	names := []string{}
	for _, booking := range bookings {
		var fieldData = strings.Fields(booking)
		names = append(names, fieldData[0])
	}
	fmt.Printf("The name of bookings are: %v\n", names)

}
