// Prints it's command-line arguments
package main

import (
	"fmt"
	"os"
)

func main() {
	var completeLine, separator string

	for i := 1; i < len(os.Args); i++ {
		completeLine += separator + os.Args[i]
		separator = " "
	}

	fmt.Println(completeLine)
}
