package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type fallibleHandler func(w http.ResponseWriter, r *http.Request) error

func main() {
	h := Handler{}
	r := mux.NewRouter()
	r.HandleFunc("/start", catchError(h.start)).Methods("POST")
	http.ListenAndServe(":8080", r)
}

func catchError(fn fallibleHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			fmt.Println(err)
			http.Error(w, "You broke the internet 🐮💩😱", http.StatusInternalServerError)
		}
	}
}
