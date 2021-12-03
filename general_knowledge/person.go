package main

import "fmt"
import "reflect"

func main() {
	name := "Victoria"
	version := 1.1

	fmt.Println("Hello,", name)
	fmt.Println("This app version is", version)

	fmt.Println("The name variable is a", reflect.TypeOf(name))
	fmt.Println("The version variable is a", reflect.TypeOf(version))
}