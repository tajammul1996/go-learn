package helper

import (
	"fmt"
	"strings"
)

func PrintFirstNames(bookings *[]string) {
	fmt.Println(bookings)
	var firstNames []string
	for _, booking := range *bookings {
		names := strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}

	for _, name := range firstNames {
		fmt.Println(name)
	}
}
