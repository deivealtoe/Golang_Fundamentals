package main

import "fmt"
import "os"
import "net/http"

func main() {
	showMenu()

	choice := readChoice()

	switch choice {
		case 1:
			startMonitor()
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

func showMenu() {
	fmt.Println("1- Start")
	fmt.Println("2- Show logs")
	fmt.Println("0- Exit")
}

func readChoice() int {
	fmt.Print("Choose: ")

	var choice int
	fmt.Scan(&choice)

	fmt.Println("The choice was", choice)

	return choice
}

func startMonitor() {
	fmt.Println("Starting...")

	site := "https://www.alura.com.br"
	response, error := http.Get(site)
	
}