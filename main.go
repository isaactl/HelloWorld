package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

func echoRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s: %s", r.Method, r.URL.String())
	w.Write([]byte("Hello World"))
}

func echoByPost(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s: %s", r.Method, r.URL.String())
	b, _ := ioutil.ReadAll(r.Body)
	fmt.Fprintf(w, "post body: %q", b)
}

func echoByPut(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s: %s", r.Method, r.URL.String())
	fmt.Fprintf(w, "put: %q", r.URL.String())
}

func echoByDelete(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s: %s", r.Method, r.URL.String())
	fmt.Fprintf(w, "delete: %q", r.URL.String())
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/get", echoRequest).Methods(http.MethodGet)
	r.HandleFunc("/post", echoByPost).Methods(http.MethodPost)
	r.HandleFunc("/put", echoByPut).Methods(http.MethodPut)
	r.HandleFunc("/delete", echoByDelete).Methods(http.MethodDelete)
	http.ListenAndServe(":80", r)
}
