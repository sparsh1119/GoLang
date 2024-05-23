package main

import "fmt"

// func divide(a, b float64) float64 {
// 	return a / b
// }

// func divide2(a, b float64) (float64, error) {
// 	if b == 0 {
// 		return 0, fmt.Errorf("denominator must not to be zero")
// 	}
// 	return a / b, nil
// }

func divide2(a, b float64) (float64, string) {
	if b == 0 {
		return 0, "denominator must not to be zero"
	}
	return a / b, "nil"
}

func main() {
	ans, _ := divide2(10, 0)
	fmt.Println(ans)

	// ans, err := divide2(10, 4)
	// if err != nil {
	// 	fmt.Println("Error handling")
	// }
	// fmt.Println(ans)
}
