package main

import "fmt"

func main() {

	x := 20

	if x > 10 {
		fmt.Println("x is greater than 10")
	} else if x > 5 {
		fmt.Println("x is greater than 5")
	} else {
		fmt.Println("x is smaller than 5")
	}

	y := 25

	if x < y && y < 20 {
		fmt.Println("Nice")
	} else {
		fmt.Println("Nicer")
	}
}
