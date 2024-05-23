package main

import "fmt"

func main() {

	studentGrades := make(map[string]int)

	studentGrades["Sparsh"] = 90
	studentGrades["Bob"] = 60
	studentGrades["Alice"] = 50
	studentGrades["Charlie"] = 40

	fmt.Println("Marks of Bob :", studentGrades["Bob"])
	studentGrades["Bob"] = 85
	fmt.Println("Updated grade of Bob", studentGrades["Bob"])

	delete(studentGrades, "Charlie")

	grades, exists := studentGrades["Charlie"]
	fmt.Println("Marks of Charlie is ", grades)
	fmt.Println("Marks of Charlie is ", studentGrades["Charlie"])
	fmt.Println("Charlie existane in map ", exists)

	studentGrades["Navneet"] = 70

	for index, marks := range studentGrades {
		fmt.Printf("key is %s and Grades is %d\n", index, marks)
	}

	person := map[string]int{
		"Joy":   9,
		"Minch": 7,
		"Soy":   8,
	}
	for index, marks := range person {
		fmt.Printf("_____Key is %s and Grades is %d\n", index, marks)
	}

}
