package main

import (
	"fmt"
	"time"
)

func main() {

	currentTime := time.Now()
	fmt.Println("Current time :", currentTime)
	fmt.Printf("%T\n", currentTime)

	formatted := currentTime.Format("02/01/2006, Monday , 3:04 PM")
	fmt.Println("Formatted time : ", formatted)

	layout_str := "02/01/2006"
	date_str := "21/05/2024"
	formatted_time, _ := time.Parse(layout_str, date_str)
	fmt.Println(formatted_time)

	new_date := currentTime.Add(48 * time.Hour)
	fmt.Println(new_date)

	formatted_new_date := new_date.Format("2006/01/02 Monday")
	fmt.Println(formatted_new_date)
}
