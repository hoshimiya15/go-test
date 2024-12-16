package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Monday?")
	today := time.Now().Weekday()
	switch time.Monday {
		case today + 0:
			fmt.Println("Today")
		case today + 1:
			fmt.Println("Tommorow")
		case today + 2:
			fmt.Println("In two days")
		default:
			fmt.Printf("Too far away")
	}
}