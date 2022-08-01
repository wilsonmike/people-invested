package main

import "fmt"

func main() {
	appName := "People Invested"
	const activeSubLimit = 50
	var remainingSubs uint = 50
	
	fmt.Printf("Welcome to the %v data stream\n", appName)
	fmt.Printf("We have a limit of %v total subs and %v are still available\n", activeSubLimit, remainingSubs)
	fmt.Println("Get your access here")

	var firstName string
	var lastName string
	var email string
	var userSubs int
	// ask user for their name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	userSubs = 2
	fmt.Printf("User %v booked %v subs.\n", firstName, userSubs)
}