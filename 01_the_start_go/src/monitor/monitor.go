package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
	"strings"
)

const amountOfTry = 3
const secondsToWaitBeforeNewTest = 4

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

	//sites := []string { "https://random-status-code.herokuapp.com/", "https://www.alura.com.br", "https://www.caelum.com.br" }
	sites := ReadSitesFile();

	for _, site := range(sites) {
		for i := 0; i < amountOfTry; i++ {
			fmt.Println(BuildMessage(site))

			time.Sleep(secondsToWaitBeforeNewTest * time.Second)
		}
	}
}

func BuildMessage(site string) string {
	siteStatusCode := GetSiteStatusCode(site)

	var message string = "Testing site: " + site

	message += ". " + GetTreatedMessage(siteStatusCode)

	return message
}

func GetSiteStatusCode(site string) int {
	response, err := http.Get(site)

	if err != nil {
		fmt.Println("An error had ocurred trying to make a Get request to", site, err)
	}

	return response.StatusCode
}

func GetTreatedMessage(responseStatusCode int) string {
	var message string

	if responseStatusCode == 200 {
		message = "Success!"
	} else {
		message = "Fail!"
	}

	return message
}

func ReadSitesFile() []string {
	var sites []string

	fileOpen, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("An error occurred trying to open specified file", err)
	}

	reader := bufio.NewReader(fileOpen)

	for  {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}

		line = strings.TrimSpace(line)
		sites = append(sites, line)
	}

	fileOpen.Close()

	return sites
}
