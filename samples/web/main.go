package main

import (
	"fmt"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func CreateHttpHandler() http.Handler {
	handler := http.NewServeMux()

	handler.HandleFunc("/", rootHandler)

	return handler
}

func RunWebServer() {
	http.ListenAndServe(":8080", CreateHttpHandler())
}

func main() {
	RunWebServer()
}
