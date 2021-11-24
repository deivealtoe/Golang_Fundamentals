package main

import "fmt"
import "reflect"

func main() {
	var name = "Alex"
	var version = 1.1

	fmt.Println("Hello,", name)
	fmt.Println("This app is running on version", version)
	fmt.Println("The name variable is a", reflect.TypeOf(name))
	fmt.Println("The version variable is a", reflect.TypeOf(version))
}
