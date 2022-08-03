package main

import "fmt"

func main() {
	appName := "People Invested"
	const activeSubLimit = 50
	var remainingSubs uint = 50
	subscriptions := []string{} 

	fmt.Printf("Welcome to the %v data stream\n", appName)
	fmt.Printf("We have a limit of %v total subs and %v are still available\n", activeSubLimit, remainingSubs)
	fmt.Println("Get your access here")

	var firstName string
	var lastName string
	var email string
	var userSubs uint
	// ask user for their name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("How many subscriptions would you like to purchase?: ")
	fmt.Scan(&userSubs)

	remainingSubs = remainingSubs - userSubs 
	subscriptions = append(subscriptions, firstName + " " + lastName)
	
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userSubs, email)
	fmt.Printf("%v subs remaining for %v\n", remainingSubs, appName)

	fmt.Printf("Thank you to all of our active users: %v", subscriptions)
}