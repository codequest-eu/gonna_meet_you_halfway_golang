package main

import (
	"encoding/json"
	"net/http"

	"github.com/codequest-eu/gonna_meet_you_halfway_golang/models"
)

type Handler struct {
	// mailer mailer.Mailer
	// store  store.Store
}

func (h *Handler) start(w http.ResponseWriter, r *http.Request) error {
	var sd models.StartData
	if err := json.NewDecoder(r.Body).Decode(&sd); err != nil {
		return err
	}
	t := models.Topics{
		suggestionsTopicName:     "suggestionsTopicName",
		myLocationTopicName:      "myLocationTopicName",
		otherLocationTopicName:   "otherLocationTopicName",
		meetingLocationTopicName: "meetingLocationTopicName",
	}
	return json.NewEncoder(w).Encode(t)
}
