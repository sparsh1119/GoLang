package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "apple,orange,banana"
	parts := strings.Split(data, ",")
	fmt.Println(parts)

	str := "one two three four five six two two two"
	count := strings.Count(str, "two")
	fmt.Println("Count :", count)

	str = "               Hello     goo!   "
	trimmed := strings.TrimSpace(str)
	fmt.Println(trimmed)

	str1 := "Sparsh"
	str2 := "Jaiswal"
	result := strings.Join([]string{str1, "Kumar", str2}, " ")
	fmt.Println("result:", result)

}
