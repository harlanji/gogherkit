package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

var server *http.Server

func RunWebServer() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)

}

func StopWebServer() {
	// TODO figure this out... 
}

func main() {
	RunWebServer()
}
