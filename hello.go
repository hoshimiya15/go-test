package main

import (
	"fmt"

	"github.com/tenntenn/greeting"
)

func main() {
	greet := greeting.Do()
	fmt.Println(greet)
}
