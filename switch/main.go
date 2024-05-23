package main

import (
	"fmt"
)

func main() {

	day := 1

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Unknown day")
	}

	month := "November"

	switch month {
	case "March", "April", "May", "June":
		fmt.Println("Summer")

	case "July", "August", "September":
		fmt.Println("Moonsoon")

	case "November", "December", "January":
		fmt.Println("Winter")

	default:
		fmt.Println("Autom")
	}

	temperature := 30

	switch {
	case temperature < 0:
		fmt.Println("Freezing")
	case temperature >= 0 && temperature < 10:
		fmt.Println("Cold")
	case temperature >= 10 && temperature < 20:
		fmt.Println("Cool")
	case temperature >= 20 && temperature < 30:
		fmt.Println("Warm")
	default:
		fmt.Println("Hot")

	}
}
