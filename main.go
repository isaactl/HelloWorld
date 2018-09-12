package main

import (
	"fmt"
	"net/http"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("get request for %s", r.URL.String())
	w.Write([]byte("\nHello world\n"))
}

func launchServer() {
	http.HandleFunc("/", handleRequest)
	http.ListenAndServe(":8080", nil)
}

func main() {
	launchServer()
}
