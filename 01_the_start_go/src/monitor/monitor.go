package main

import "fmt"
import "os"
import "net/http"

func main() {
	for {
		ShowMenu()

		var sites [4]string
		sites[0] = "https://random-status-code.herokuapp.com/"
		sites[1] = "https://random-status-code.herokuapp.com/"
		sites[2] = "https://random-status-code.herokuapp.com/"

		fmt.Println(sites[3])

		choice := ReadChoice()
	
		switch choice {
			case 1:
				StartMonitor()
			case 2:
				fmt.Println("Logs...")
			case 0:
				fmt.Println("See ya...")
				os.Exit(0)
			default:
				fmt.Println("Invalid choice!")
				os.Exit(-1)
		}
	}
}

func ShowMenu() {
	fmt.Println("1- Start")
	fmt.Println("2- Show logs")
	fmt.Println("0- Exit")
}

func ReadChoice() int {
	fmt.Print("Choose: ")

	var choice int
	fmt.Scan(&choice)

	fmt.Println("The choice was", choice)

	return choice
}

func StartMonitor() {
	fmt.Println("Starting...")

	site := "https://random-status-code.herokuapp.com/"
	response, _ := http.Get(site)

	responseStatusCode := response.StatusCode

	if responseStatusCode == 200 {
		fmt.Println("Site:", site, "was loaded with success!")
	} else {
		fmt.Println("There is a problem with the site:", site)
	}

	fmt.Println("The status code was", responseStatusCode)
}
