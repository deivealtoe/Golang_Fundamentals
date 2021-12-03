package main

import "fmt"

func main() {
	fmt.Println("1- Good morning!")
	fmt.Println("2- Good evening!")
	fmt.Println("3- Good night!")

	fmt.Print("Number from 1 to 3: ")

	var choice int
	fmt.Scan(&choice)

	switch choice {
		case 1:
			fmt.Println("Good morning to you too!")
		case 2:
			fmt.Println("Good evening miss!")
		case 3:
			fmt.Println("Time to sleep...")
		default:
			fmt.Println("Not a valid option...")
	}
}
