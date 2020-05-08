package main

import (
	"fmt"
	"time"
)

func weekend(d time.Weekday) bool {
	if int(d) == 5 {
		return true
	} else {
		return false
	}
}

func main() {
	today := time.Now().Weekday()
	fmt.Printf("Today is: %s\n", today)
	if weekend(today) {
		fmt.Println("TGIF!!")
	}
}
