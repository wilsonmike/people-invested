package main

import "fmt"

func main() {
	appName := "People Invested"
	const activeSubLimit = 50
	var remainingSubs uint = 50
	
	fmt.Printf("Welcome to the %v data stream\n", appName)
	fmt.Printf("We have a limit of %v total subs and %v are still available\n", activeSubLimit, remainingSubs)
	fmt.Println("Get your access here")

	var userName string
	var userSubs int
	// ask user for their name
	// fmt.Scan(&userName)

	fmt.Println(remainingSubs)
	fmt.Println(&remainingSubs)
	
	userSubs = 2
	fmt.Printf("User %v booked %v subs.\n", userName, userSubs)
}