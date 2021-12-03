package main

import (
	"fmt"
	"reflect"
)

func main() {
	ShowSliceInfos()
}

func ShowSliceInfos() {
	var mySlice []int
	mySlice = []int{ 1, 2, 3, 4 }

	fmt.Println("The content of my slice is:", mySlice)
	fmt.Println("The variable mySlice is of type:", reflect.TypeOf(mySlice))

	fmt.Println("The length of mySlice is:", len(mySlice))
	fmt.Println("The capacity of mySlice is:", cap(mySlice))

	fmt.Println("Appending one more element")
	mySlice = append(mySlice, 5)

	fmt.Println("The length of mySlice is:", len(mySlice))
	fmt.Println("The capacity of mySlice is:", cap(mySlice))
}
