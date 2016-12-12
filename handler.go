package main

import (
	"encoding/json"
	"net/http"

	"github.com/codequest-eu/gonna_meet_you_halfway_golang/mailer"
	"github.com/codequest-eu/gonna_meet_you_halfway_golang/models"
	"github.com/satori/go.uuid"
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
	t := h.generateTopics()
	meeting := models.NewMeeting{
		Identifier: uuid.NewV4().String(),
		Topics:     t,
	}
	return json.NewEncoder(w).Encode(meeting)
}

func (h *Handler) generateTopics() models.Topics {
	return models.Topics{
		SuggestionsTopicName:     uuid.NewV4().String(),
		MyLocationTopicName:      uuid.NewV4().String(),
		OtherLocationTopicName:   uuid.NewV4().String(),
		MeetingLocationTopicName: uuid.NewV4().String(),
	}
}
