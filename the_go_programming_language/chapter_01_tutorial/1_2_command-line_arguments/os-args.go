package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {
	fmt.Println("First argument: ", os.Args[1])
	fmt.Println("Second argument: ", os.Args[2])
	fmt.Println(reflect.TypeOf(os.Args[0]))
	fmt.Println(reflect.TypeOf(os.Args[1]))
	fmt.Println(reflect.TypeOf(os.Args[2]))
}
