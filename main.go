package main

import (
	"fmt"
	"net/http"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("get request for %s", r.URL.String())
	w.Write([]byte("\nv2: Hello world\n"))
}

func launchServer() {
	http.HandleFunc("/", handleRequest)
	http.ListenAndServe(":80", nil)
}

func main() {
	launchServer()
}

// Print testing
func Print() {
	fmt.Printf("hello world")
}
