package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"is_adult"`
}

func main() {

	man := Person{Name: "Sparsh", Age: 22, IsAdult: true}

	fmt.Println(man)

	//converting to Json (marshalling)
	jsonData, err := json.Marshal(man)
	if err != nil {
		fmt.Println("Error Marshalling", err)
		return
	}
	fmt.Println(string(jsonData))

	//Converting from json(unmarshalling)
	var personData Person
	err = json.Unmarshal(jsonData, &personData)
	if err != nil {
		fmt.Println("Error Decoding", err)
		return
	}
	fmt.Println(personData)

}
