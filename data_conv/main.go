package main

import (
	"fmt"
	"strconv"
)

func main() {

	num := 5
	fmt.Println(num)
	fmt.Printf("%T\n", num)

	data := float64(num)
	data = data + 33
	fmt.Println(data)
	fmt.Printf("%T\n", data)

	num_string := strconv.Itoa(num)
	fmt.Println(num_string)
	fmt.Printf("%T\n", num_string)

	str := "70666"
	string_number, _ := strconv.Atoi(str)
	fmt.Println(string_number)
	fmt.Printf("%T\n", string_number)

	str2 := "3.14"
	string_float, _ := strconv.ParseFloat(str2, 64)
	fmt.Println(string_float)
	fmt.Printf("%T\n", string_float)

	// str3 := "9090"
	// str_Int, _ := strconv.ParseInt(str3, 0, 0)
	// fmt.Println(str_Int)
	// fmt.Printf("%T\n", str_Int)

}
