package main

import (
	"fmt"
)

type buyTransaction struct {
	userName     string
	emailAddress string
	qty          uint
	cost         float32
}

func main() {

	const conferenceName string = "Go Conference"
	const conferenceTickets uint = 500
	const basicTicketPrice float32 = 23.45

	var remainingTickets uint = conferenceTickets
	var userName string
	var emailAddress string
	var userTickets uint

	welcomeMessage(conferenceName, remainingTickets)

	// Buy a ticket
	var trans buyTransaction = getTicket(userName, emailAddress, userTickets, basicTicketPrice)
	remainingTickets = calcRemaingTickets(trans.qty, conferenceTickets)
	fmt.Printf("User: %v(%v) has bought %v tickets, costing £%.2f leaving %d tickets remaining.\n", trans.userName, trans.emailAddress, trans.qty, trans.cost, remainingTickets)

}

func welcomeMessage(conferenceName string, remainingTickets uint) {
	fmt.Printf("Welcome to our %v booking app\n", conferenceName)
	fmt.Println("Get your tickets here....")
	fmt.Printf("We have %v tickets still remaining.\n", remainingTickets)
}

func calcRemaingTickets(userTickets uint, conferenceTickets uint) uint {
	return conferenceTickets - userTickets
}

func getTicket(userName string, emailAddress string, userTickets uint, basicTicketPrice float32) buyTransaction {
	fmt.Printf("Enter username: ")
	fmt.Scan(&userName)
	fmt.Printf("Enter email address: ")
	fmt.Scan(&emailAddress)
	fmt.Printf("Enter quantity: ")
	fmt.Scan(&userTickets)

	var trans = buyTransaction{
		userName:     userName,
		emailAddress: emailAddress,
		qty:          userTickets,
		cost:         (float32(userTickets) * basicTicketPrice),
	}

	return trans

}
