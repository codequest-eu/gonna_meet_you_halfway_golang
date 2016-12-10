package main

import "github.com/gorilla/mux"
import "net/http"
import "fmt"

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler).Methods("GET")
	http.ListenAndServe(":8080", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
