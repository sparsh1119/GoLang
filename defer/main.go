package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("First")
	data := add(2, 5)
	defer fmt.Println(data)
	defer fmt.Println("Second")
	fmt.Println("Third")

}

// defer execute in last and follow Lifo approuch when more one exits
