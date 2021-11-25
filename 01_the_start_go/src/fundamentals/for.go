package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	randomNumbers := GenerateSliceWithRandomIntegers(10)
	fmt.Println(randomNumbers)

	randomNumbers = BubbleSort(randomNumbers)
	fmt.Println(randomNumbers)
}

func GenerateSliceWithRandomIntegers(sliceSize int) []int {
	slice := []int{}

	for i := 0; i < sliceSize; i++ {
		slice = append(slice, GetRandomInt())
	}

	return slice
}

func GetRandomInt() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(1000)
}

func BubbleSort(arr []int) []int {
	size := len(arr)

	for i := 0; i < size - 1; i++ {
		for j := 0; j < size - 1; j++ {
			if arr[j] > arr[j + 1] {
				temp := arr[j]
				arr[j] = arr[j + 1]
				arr[j + 1] = temp
			}
		}
	}

	return arr
}
