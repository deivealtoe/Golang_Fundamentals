package main

import "fmt"

func main() {
	fmt.Print("Inform your name: ")
	
	var name string
	fmt.Scan(&name)

	fmt.Println("Hello there,", name)
	fmt.Println("The address of variable name is", &name)
}