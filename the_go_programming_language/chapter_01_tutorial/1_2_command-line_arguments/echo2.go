package main

import (
	"fmt"
	"os"
)

func main() {
	a := ""
	var b string
	var c = ""
	var d string = ""
	
	completeLine, separator := "", ""

	for _, argument := range os.Args[1:] {
		completeLine += separator + argument
		separator = " "
	}

	fmt.Println(completeLine)
}
