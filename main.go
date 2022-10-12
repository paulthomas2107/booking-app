package main

import "fmt"

func main() {

	const conferenceName string = "Go Conference"
	const conferenceTickets uint = 500
	const basicTicketPrice float32 = 23.45
	someValue := 1222.22

	var remainingTickets uint = conferenceTickets
	var userName string
	var userTickets uint
	var bundleCost float32

	fmt.Printf("Welcome to our %v booking app\n", conferenceName)
	fmt.Println("Get your tickets here....")
	fmt.Printf("We have %v tickets still remaining.\n", remainingTickets)
	fmt.Printf("somevalue is %T\n", someValue)

	// Buy a ticket
	userName = "Caitlin"
	userTickets = 2
	bundleCost = float32(userTickets) * basicTicketPrice
	remainingTickets -= userTickets
	fmt.Printf("User: %s has bought %d tickets, costing £%.2f, leaving %d tickets remaining.\n", userName, userTickets, bundleCost, remainingTickets)

	// Get input data name and qty
	fmt.Printf("Enter username: ")
	fmt.Scan(&userName)
	fmt.Printf("Enter quantity: ")
	fmt.Scan(&userTickets)
	bundleCost = float32(userTickets) * basicTicketPrice
	remainingTickets -= userTickets
	fmt.Printf("User: %s has bought %d tickets, costing £%.2f, leaving %d tickets remaining.\n", userName, userTickets, bundleCost, remainingTickets)

}
