package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/codequest-eu/gonna_meet_you_halfway_golang/mailer"
	"github.com/gorilla/mux"
)

type fallibleHandler func(w http.ResponseWriter, r *http.Request) error

func main() {
	m := mailer.NewSendGridMailer(os.Getenv("SENDGRID_API_KEY"))
	h := Handler{mailer: m}
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
