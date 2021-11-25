package main

import (
	"fmt"
)

func main() {
	InteractingWithSlice()
}

func InteractingWithSlice() {
	var mySlice []int

	ShowLenAndCap(mySlice)
	mySlice = append(mySlice, 0)
	ShowLenAndCap(mySlice)
	mySlice = append(mySlice, 1)
	ShowLenAndCap(mySlice)
	mySlice = append(mySlice, 2)
	ShowLenAndCap(mySlice)
	mySlice = append(mySlice, 3)
	ShowLenAndCap(mySlice)
	mySlice = append(mySlice, 4)
	ShowLenAndCap(mySlice)
	mySlice = append(mySlice, 5)
	ShowLenAndCap(mySlice)
	mySlice = append(mySlice, 6)
	ShowLenAndCap(mySlice)
	mySlice = append(mySlice, 7)
	ShowLenAndCap(mySlice)
	mySlice = append(mySlice, 8)
	ShowLenAndCap(mySlice)
	mySlice = append(mySlice, 9)
	ShowLenAndCap(mySlice)
	
	fmt.Println("Elements of mySlice:", mySlice)
}

func ShowLenAndCap(mySlice []int) {
	fmt.Println("The length is", len(mySlice), "Capacity is", cap(mySlice))
}
