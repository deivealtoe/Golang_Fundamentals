package main

import "fmt"

func main() {
	fmt.Println("Input your age: ")

	var age int
	fmt.Scan(&age)

	if age < 18 {
		fmt.Println("Minor")
	} else if age >= 18 && age < 60 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Elderly")
	}
}
