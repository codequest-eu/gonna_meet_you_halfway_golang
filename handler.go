package main

import (
	"encoding/json"
	"net/http"

	"github.com/codequest-eu/gonna_meet_you_halfway_golang/broadcaster"
	"github.com/codequest-eu/gonna_meet_you_halfway_golang/mailer"
	"github.com/codequest-eu/gonna_meet_you_halfway_golang/models"
	"github.com/codequest-eu/gonna_meet_you_halfway_golang/storage"
	"github.com/satori/go.uuid"
)

type Handler struct {
	mailer      mailer.Mailer
	broadcaster broadcaster.Broadcaster
	store       storage.Store
}

func (h *Handler) start(w http.ResponseWriter, r *http.Request) error {
	var inviteData models.InviteData
	if err := json.NewDecoder(r.Body).Decode(&inviteData); err != nil {
		return err
	}
	meetingID := uuid.NewV4().String()
	if err := h.mailer.Mail(inviteData, meetingID); err != nil {
		return err
	}
	topics := h.generateTopics()
	if err := h.store.SaveTopics(meetingID, topics); err != nil {
		return err
	}
	meeting := models.NewMeeting{
		Identifier: meetingID,
		Topics:     topics,
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
