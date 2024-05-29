package main

import (
	"fmt"
	"time"
)

func first() {
	fmt.Println("Namaste")

	time.Sleep(1000 * time.Millisecond)

	fmt.Println("Nice to meet you")
}

func second() {
	fmt.Println("Mucho Gusto")

	time.Sleep(2000 * time.Millisecond)

	fmt.Println("Adios")
}

func main() {
	fmt.Println("GoRoutines")
	go first()
	go second()

	time.Sleep(5000 * time.Millisecond)

}
