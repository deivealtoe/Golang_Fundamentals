package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", indexPageHandler)
	http.HandleFunc("/docker", dockerPageHandler)
	http.ListenAndServe(":8080", nil)
}

func indexPageHandler(responseWriter http.ResponseWriter, request *http.Request) {
	http.ServeFile(responseWriter, request, "./static/index.html")
}

func dockerPageHandler(responseWriter http.ResponseWriter, request *http.Request) {
	http.ServeFile(responseWriter, request, "./static/docker.html")
}
