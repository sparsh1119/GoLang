package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	count := 0
	for {
		fmt.Println("Infinite loop")
		count++
		if count == 5 {
			break
		}
	}

	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Printf("index: %d , value: %d\n", index, value)
	}

	data := "Namaste"
	for index, value := range data {
		fmt.Printf("index: %d , value: %c\n", index, value)
	}

}
