package main

import (
	"fmt"

	"go-test/zoo/animals"
)

func main() {
	fmt.Println(AppName()) // run "go run main.go app.go"

	fmt.Println(animals.ElephantFeed())
	fmt.Println(animals.MonkeyFeed())
	fmt.Println(animals.RabbitFeed())
}