package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("what's your name??")

	// var name string
	// fmt.Scan(&name)

	//for full name or reading space also for senctence
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n') //till it does'nt get new line

	fmt.Println("Hello , Mr.", name)

}
