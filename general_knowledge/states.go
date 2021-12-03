package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(GetStatesArray())

	ShowNames()
}

func GetStatesArray() [4]string {
	var statesArray [4]string

	statesArray[0] = "RJ"
	statesArray[1] = "SP"
	statesArray[2] = "MG"
	statesArray[3] = "ES"

	fmt.Println("Type of statesArray is:", reflect.TypeOf(statesArray))

	return statesArray
}

func ShowNames() {
	names := []string{"Alex", "Victoria", "Mary"}

	fmt.Println("Type of names is:", reflect.TypeOf(names))

	fmt.Println(names)
}
