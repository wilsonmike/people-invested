package main

import (
	"fmt"
	"strings"
) 

func main() {
	appName := "People Invested"
	const activeSubLimit = 50
	var remainingSubs uint = 50
	subscriptions := []string{} 

	greetUsers(appName, activeSubLimit, remainingSubs)

	for {
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

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidSubscription := userSubs > 0 && userSubs <= remainingSubs

		if isValidName && isValidEmail && isValidSubscription {
			remainingSubs = remainingSubs - userSubs 
			subscriptions = append(subscriptions, firstName + " " + lastName)
			
			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userSubs, email)
			fmt.Printf("%v subs remaining for %v\n", remainingSubs, appName)

			firstNames := getFirstNames(subscriptions)
			fmt.Printf("The first names of bookings are: %v\n", firstNames)
			
			if remainingSubs == 0 {
				fmt.Println("Oops, looks like we are out of available subscriptions this month.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("The first name or last name you entered is too short.")
			}
			if !isValidEmail {
				fmt.Println("The email you entered is invalid please re-enter.")
			}
			if !isValidSubscription {
				fmt.Println("The number of subscriptions you entered is invalid.")
			}
		}
	}
}

func greetUsers(greetAppName string, subTickets int, subRemaining uint) {
	fmt.Printf("Welcome to %v\n", greetAppName)
	fmt.Printf("We have a limit of %v total subs and %v are still available\n", subTickets, subRemaining)
	fmt.Println("Get your access here")
}

func getFirstNames(subscriptions []string) []string {
	firstNames := []string{}
	for _, booking := range subscriptions {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	} 
	return firstNames
}