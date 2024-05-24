package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		fmt.Println("Error", err)
		return
	}
	defer res.Body.Close()

	fmt.Printf("%T\n", res)
	//fmt.Println("Response : ", res)  this could not return data

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Println("response :", string(data))

}
