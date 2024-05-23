package main

import "fmt"

func main() {

	// numbers := []int{1, 2, 3, 4, 5}
	// fmt.Println(numbers)
	// numbers = append(numbers, 6, 7, 8, 9, 10)
	// fmt.Println(numbers)
	// fmt.Printf("Data type : %T\n", numbers)
	// fmt.Println("Length : ", len(numbers))

	// name := []string{}
	// fmt.Println("name : ", name)

	numbers := make([]int, 3, 5) //type size and capacity
	fmt.Println(numbers)

	numbers = append(numbers, 4)
	numbers = append(numbers, 98)
	numbers = append(numbers, 6)
	numbers = append(numbers, 4)
	numbers = append(numbers, 98)
	numbers = append(numbers, 6)
	numbers = append(numbers, 6)
	numbers = append(numbers, 6)

	fmt.Println("Slice:", numbers)
	fmt.Println("Length:", len(numbers))
	fmt.Println("Capacity:", cap(numbers))

	stock := make([]int, 0)
	fmt.Println("Slice:", stock)
	fmt.Println("Length:", len(stock))
	fmt.Println("Capacity:", cap(stock))

}
