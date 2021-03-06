package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
	"strconv"
)

const amountOfTries = 3
const secondsToWaitBeforeNewTest = 1

func main() {
	for {
		ShowMenu()

		choice := ReadChoice()
	
		switch choice {
			case 1:
				StartMonitor()
			case 2:
				ShowLog()
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

	return choice
}

func StartMonitor() {
	fmt.Println("Starting...")

	sites := ReadSitesFile();

	for i := 0; i < amountOfTries; i++ {
		for _, site := range(sites) {	
			message := BuildMessage(site)
			
			fmt.Println(message)
			LogResgistry(message)

			time.Sleep(secondsToWaitBeforeNewTest * time.Second)
		}
	}
}

func BuildMessage(site string) string {
	siteStatusCode := GetSiteStatusCode(site)

	var message string = time.Now().Format("02/01/2006 - 15:04:05") + " Testing site: " + site + " - Status Code Response: " + strconv.Itoa(siteStatusCode)

	message += ". " + GetSuccessOrFail(siteStatusCode)

	return message
}

func GetSiteStatusCode(site string) int {
	response, err := http.Get(site)

	if err != nil {
		fmt.Println("An error had ocurred trying to make a Get request to", site, err)
	}

	return response.StatusCode
}

func GetSuccessOrFail(responseStatusCode int) string {
	var message string

	if responseStatusCode >= 200 && responseStatusCode <= 299 {
		message = "Success!"
	} else {
		message = "Fail!"
	}

	return message
}

func ReadSitesFile() []string {
	var sites []string

	sitesFile, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("An error occurred trying to read sites file")
	}

	reader := bufio.NewReader(sitesFile)

	for  {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}

		line = strings.TrimSpace(line)
		sites = append(sites, line)
	}

	sitesFile.Close()

	return sites
}

func LogResgistry(message string) {
	logFile, err := os.OpenFile("log.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("An error ocurred trying to write into log file")
	}

	logFile.WriteString(message + " - " + "\n")

	logFile.Close()
}

func ShowLog() {
	fmt.Println("Logs...")

	file, err := ioutil.ReadFile("log.txt")
	if err != nil {
		fmt.Println("An error had ocurred trying to read log file")
	}

	fmt.Println(string(file))
}
