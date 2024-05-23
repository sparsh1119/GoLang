package main

import "fmt"

func main() {
	age := 22
	name := "alice"
	height := 5.4442211

	fmt.Println("age :", age, "height :", height, "name :", name)

	//f is fomator
	fmt.Printf("age is %d\n", age)
	fmt.Printf("height is %2f\n", height)
	fmt.Printf("Type of age is %T\n", age)
	fmt.Printf("Type of height is %T\n", height)

	fmt.Printf("Name is %s\n", name)
	fmt.Printf("Name: %s, Age %d , Height %f\n", name, age, height)

}
