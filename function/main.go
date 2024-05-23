package main

import "fmt"

func simplefunction() {
	fmt.Println("Simple Function")
}

func add(a, b int) int {
	return a + b
}

func multiply(a int, b int) (result int) {
	result = a * b
	return
}

func main() {
	fmt.Println("learning function")
	simplefunction()

	ans := add(2, 8)
	fmt.Println(ans)

	product := multiply(2, 8)
	fmt.Println(product)
}
