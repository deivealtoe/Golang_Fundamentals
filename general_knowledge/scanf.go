package main

import "fmt"

func main() {
	fmt.Print("Please, inform your name: ")
	
	var name string
	fmt.Scanf("%s", &name)

	fmt.Println("Welcome,", name, "!")
}
