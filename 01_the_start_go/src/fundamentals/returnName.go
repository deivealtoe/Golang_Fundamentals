package main

import "fmt"

func main() {
	name, age := ReturnNameAndAge()

	fmt.Println("Hello " + name + "! You are", age, "years old, right?")
}

func ReturnNameAndAge() (string, int) {
	name := "Silvia"
	age := 26

	return name, age
}
