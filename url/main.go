package main

import (
	"fmt"
	"net/url"
)

func main() {

	myurl := "https://github.com/sparsh1119/GoLang/blob/main/main.go/?12121"

	fmt.Printf("Url Type : %T\n", myurl)

	parsedUrl, err := url.Parse(myurl)
	if err != nil {
		fmt.Println("Can't parse URL", err)
		return
	}
	fmt.Printf("Url Type : %T\n", parsedUrl)

	fmt.Println("Scheme: ", parsedUrl.Scheme)
	fmt.Println("Host : ", parsedUrl.Host)
	fmt.Println("Path :", parsedUrl.Path)
	fmt.Println("RawQuery :", parsedUrl.RawQuery)

	//modifying url

	parsedUrl.Path = "/newPath"
	parsedUrl.RawQuery = "/username=sparsh"

	newUrl := parsedUrl.String()

	fmt.Println("New Url : ", newUrl)

}
