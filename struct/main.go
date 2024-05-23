package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type Contact struct {
	Email string
	Phone string
}

type Address struct {
	House int
	Area  string
	State string
}

type Employee struct {
	Person_Details Person
	Person_Contact Contact
	Person_Address Address
}

func main() {

	var man Person
	fmt.Println(man) //__0
	man.FirstName = "Naveen"
	man.LastName = "Raj"
	man.Age = 30
	fmt.Println(man)

	person1 := Person{
		FirstName: "Sparsh",
		LastName:  "Jaiswal",
		Age:       22,
	}
	fmt.Println(person1)

	var person2 = new(Person)
	person2.FirstName = "Joy"
	person2.LastName = "Mich"
	person2.Age = 43
	fmt.Println(person2)

	var employee1 Employee
	employee1.Person_Details = Person{
		FirstName: "Rudra",
		LastName:  "Joshi",
		Age:       30,
	}

	employee1.Person_Contact.Email = "rudra11@gmail.com"
	employee1.Person_Contact.Phone = "986542312"

	employee1.Person_Address = Address{
		House: 89,
		Area:  "Surat",
		State: "Gujrat",
	}

	fmt.Println(employee1)

}
