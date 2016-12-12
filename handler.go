package main

import (
	"encoding/json"
	"net/http"

	"github.com/codequest-eu/gonna_meet_you_halfway_golang/mailer"
	"github.com/codequest-eu/gonna_meet_you_halfway_golang/models"
)

type Handler struct {
	mailer mailer.Mailer
	// store  store.Store
}

func (h *Handler) start(w http.ResponseWriter, r *http.Request) error {
	var sd models.StartData
	if err := json.NewDecoder(r.Body).Decode(&sd); err != nil {
		return err
	}
	if err := h.mailer.Mail(sd.Email, sd.Name, sd.OtherEmail); err != nil {
		return err
	}
	t := models.Topics{
		SuggestionsTopicName:     "suggestionsTopicName",
		MyLocationTopicName:      "myLocationTopicName",
		OtherLocationTopicName:   "otherLocationTopicName",
		MeetingLocationTopicName: "meetingLocationTopicName",
	}
	return json.NewEncoder(w).Encode(t)
}
