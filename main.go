package main

import (
	"fmt"
	"learn/myUtil"
)

func main() {
	fmt.Println("Namaste goLang")
	myutil.PrintMessage("Hello")

	var name string = "Sparsh"
	fmt.Println(name)

	var age = 22
	fmt.Println(age)

	var money int = 6000
	fmt.Println(money)

	var dimension float64 = 12.22232
	fmt.Println(dimension)

	var decide bool = true
	fmt.Println(decide)

	var person = "Sparsh Jaiswal"
	fmt.Println("Name :", person)

	const pi = 3.14
	//pi = 102.22 cannot be done
	fmt.Println(pi)

	//also variable can be declare like this
	person1 := "Naveen Raj"
	fmt.Println(person1)

	var Public = "data is public"   //upercase
	var private = "data is private" //lowercase

	fmt.Println(Public, private)

}
