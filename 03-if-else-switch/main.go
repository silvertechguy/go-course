package main

import (
	"fmt"
	"time"
)

func main() {
	playerName := "Mr. x"
	playerName = "Eslam"
	if playerName == "Eslam" {
		fmt.Println("Hi", playerName)
	}

	// if time.Now().Weekday() == time.Friday || time.Now().Weekday() == time.Saturday {}
	switch time.Now().Weekday() {
	case time.Friday, time.Saturday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

}
