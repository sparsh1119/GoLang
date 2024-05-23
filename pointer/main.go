package main

import "fmt"

func modifyValue(num *int) {
	*num = *num + 20
}

func main() {

	// var num int
	// num = 5

	num := 5

	// var ptr *int
	// ptr = &num

	ptr := &num

	fmt.Println(ptr)
	fmt.Println(*ptr)
	fmt.Println(num)

	var pointer *int
	if pointer == nil {
		fmt.Println("Pointer is null")
	}

	value := 5
	modifyValue(&value)
	fmt.Println(value)

}
