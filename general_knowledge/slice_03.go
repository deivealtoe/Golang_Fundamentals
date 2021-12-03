package main

import "fmt"

func main() {
	MySlice()
}

func MySlice() {
	var names []string = []string { "Dave", "Alex", "Maria", "Victoria", "Foo" }

	ShowInfoOfSlice(names)

	names = append(names, "Bar")

	ShowInfoOfSlice(names)
}

func ShowInfoOfSlice(slice []string) {
	fmt.Println("Length:", len(slice), "- Capacity:", cap(slice))
}