package main

import "fmt"

func main() {
	fmt.Println("Array")

	var name [5]string
	name[0] = "sparsh"
	name[2] = "navneet"

	fmt.Println(name)

	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers)

	fmt.Println(len(numbers))

	fmt.Println(name[2])
	fmt.Println(name[1])

	var price [5]int
	fmt.Println(price) //00000 default valu

	var cars [6]string
	cars[2] = "bmw"
	cars[5] = "nexa"
	fmt.Println(cars)        //empty string
	fmt.Printf("%q\n", cars) //quoted

}
