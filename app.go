package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/codequest-eu/gonna_meet_you_halfway_golang/broadcaster"
	"github.com/codequest-eu/gonna_meet_you_halfway_golang/mailer"
	"github.com/codequest-eu/gonna_meet_you_halfway_golang/storage"
	"github.com/gorilla/mux"
)

type fallibleHandler func(w http.ResponseWriter, r *http.Request) error

func main() {
	m := mailer.NewSendGridMailer(os.Getenv("SENDGRID_API_KEY"))

	host := os.Getenv("MQTT_BROKER_HOST")
	user := os.Getenv("MQTT_BROKER_USER")
	pass := os.Getenv("MQTT_BROKER_PASSWORD")
	b, err := broadcaster.NewMQTTBroadcaster(host, user, pass)
	defer b.Close()
	if err != nil {
		log.Fatal(err)
	}

	projectID := os.Getenv("DATASTORE_PROJECT_ID")
	s, err := storage.NewGoogleStorage("datastore.key", projectID)
	if err != nil {
		log.Fatal(err)
	}

	h := Handler{mailer: m, broadcaster: b, store: s}
	r := mux.NewRouter()
	r.HandleFunc("/start", catchError(h.start)).Methods("POST")
	r.HandleFunc("/accept_meeting", catchError(h.acceptMeeting)).Methods("POST")
	r.HandleFunc("/accept_meeting/{meetingID}", catchError(h.acceptMeetingRedirect)).Methods("GET")
	r.HandleFunc("/suggest_meeting_location", catchError(h.suggestMeetingLocation)).Methods("POST")
	r.HandleFunc("/accept_meeting_location", catchError(h.acceptMeetingLocation)).Methods("POST")
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
