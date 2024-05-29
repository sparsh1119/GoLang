package main

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
	//"io/ioutil"
	"net/http"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func performGetRequest() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error getting", err)
		return
	}

	if res.StatusCode != http.StatusOK {
		fmt.Println("Error", res.Status)
		return
	}

	// data, err := io.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println("Error", err)
	// 	return
	// }
	// fmt.Println("Data", string(data))

	var todo Todo
	err = json.NewDecoder(res.Body).Decode(&todo)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	fmt.Println("Todo:", todo)
}

func performPostRequest() {

	todo := Todo{
		UserId:    11,
		Title:     "Work",
		Completed: true,
	}

	data, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	jsonString := string(data)
	fmt.Println(jsonString)

	jsonReader := strings.NewReader(jsonString)

	myUrl := "https://jsonplaceholder.typicode.com/todos"

	res, err := http.Post(myUrl, "application/json", jsonReader)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	defer res.Body.Close()

	listedData, _ := io.ReadAll(res.Body)
	fmt.Println("Response :", string(listedData))

}

func performUpdateRequest() {
	todo := Todo{
		UserId:    112121,
		Title:     "Work Hello world",
		Completed: false,
	}

	data, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	jsonString := string(data)
	jsonReader := strings.NewReader(jsonString)

	const myurl = "https://jsonplaceholder.typicode.com/todos/1"

	req, err := http.NewRequest(http.MethodPut, myurl, jsonReader)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	req.Header.Set("Content-type", "application/json")

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	defer res.Body.Close()

	upadatedData, _ := io.ReadAll(res.Body)
	fmt.Println("Response: ", string(upadatedData))
	fmt.Println("Status:", res.Status)
}

func performDeleteReqest() {

	const myurl = "https://jsonplaceholder.typicode.com/todos/1"

	req, err := http.NewRequest(http.MethodDelete, myurl, nil)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	defer res.Body.Close()

	fmt.Println("Status:", res.Status)

}

func main() {
	fmt.Println("CRUD")

	//performGetRequest()

	//performPostRequest()

	//performUpdateRequest()

	performDeleteReqest()

}
