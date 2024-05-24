package main

import (
	"fmt"
	//"io"
	"os"
)

func main() {

	// file, err := os.Create("example.txt")
	// if err != nil {
	// 	fmt.Println("Error while creating file", err)
	// 	return
	// }
	// defer file.Close()

	// content := "Namaste !!"
	// byte, error := io.WriteString(file, content+"\n")
	// fmt.Println(byte)
	// if error != nil {
	// 	fmt.Println("Error", error)
	// 	return
	// }

	// fmt.Println("File Created ")

	// file, err := os.Open("example.txt")
	// if err != nil {
	// 	fmt.Println("Error while creating file", err)
	// 	return
	// }
	// defer file.Close()

	// buffer := make([]byte, 1024)

	// for {
	// 	n, err := file.Read(buffer)
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	if err != nil {
	// 		fmt.Println("Error reading file", err)
	// 		return
	// 	}

	// 	// Process to read file content
	// 	fmt.Println(string(buffer[:n])) // read from 0 to n

	// }

	// read file into byte slice not recommeneded for large files
	content, err := os.ReadFile("example.txt")
	if err != nil {
		fmt.Println("Error ", err)
		return
	}
	fmt.Println(string(content))

}
