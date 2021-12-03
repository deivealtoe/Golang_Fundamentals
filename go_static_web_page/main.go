package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}

func indexHandler(responseWriter http.ResponseWriter, request *http.Request) {
	http.ServeFile(responseWriter, request, "./static/index.html")
}
