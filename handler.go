package main

import (
	"encoding/json"
	"net/http"

	"fmt"

	"github.com/codequest-eu/gonna_meet_you_halfway_golang/broadcaster"
	"github.com/codequest-eu/gonna_meet_you_halfway_golang/mailer"
	"github.com/codequest-eu/gonna_meet_you_halfway_golang/meeting"
	"github.com/codequest-eu/gonna_meet_you_halfway_golang/models"
	"github.com/codequest-eu/gonna_meet_you_halfway_golang/storage"
	"github.com/codequest-eu/gonna_meet_you_halfway_golang/util"
	"github.com/gorilla/mux"
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

	meetingSuggestion := models.MeetingSuggestion{LocationA: inviteData.Location}
	if err := h.store.SaveMeetingSuggestion(meetingID, meetingSuggestion); err != nil {
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

func (h *Handler) acceptMeeting(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	meetingID := vars["meetingID"]

	var acceptData models.AcceptData
	if err := json.NewDecoder(r.Body).Decode(&acceptData); err != nil {
		return err
	}

	meetingSuggestion, err := h.store.GetMeetingSuggestion(meetingID)
	if err != nil {
		return err
	}
	meetingSuggestion.SetLocationB(acceptData.Location)
	middlePoint, err := meeting.CalculateMiddlePoint(meetingSuggestion.LocationA, meetingSuggestion.LocationB)
	if err != nil {
		return err
	}
	meetingSuggestion.SetCenter(*middlePoint)

	venues, err := meeting.AskForVenues(*middlePoint)
	if err != nil {
		return err
	}
	meetingSuggestion.SetVenues(*venues)
	fmt.Println(meetingSuggestion)

	if err := h.store.SaveMeetingSuggestion(meetingID, meetingSuggestion); err != nil {
		return err
	}

	topics, err := h.store.GetTopics(meetingID)
	if err != nil {
		return err
	}

	if err := h.broadcaster.PublishMeetingSuggestion(meetingSuggestion, topics.SuggestionsTopicName); err != nil {
		return err
	}

	meeting := models.NewMeeting{
		Identifier: meetingID,
		Topics:     h.swapTopicsForOtherDevice(topics),
	}
	return json.NewEncoder(w).Encode(meeting)
}

func (h *Handler) generateTopics() models.Topics {
	char := '-'
	return models.Topics{
		SuggestionsTopicName:     util.RemoveCharacter(char, uuid.NewV4().String()),
		MyLocationTopicName:      util.RemoveCharacter(char, uuid.NewV4().String()),
		OtherLocationTopicName:   util.RemoveCharacter(char, uuid.NewV4().String()),
		MeetingLocationTopicName: util.RemoveCharacter(char, uuid.NewV4().String()),
	}
}

func (h *Handler) swapTopicsForOtherDevice(topics models.Topics) models.Topics {
	return models.Topics{
		topics.SuggestionsTopicName,
		topics.OtherLocationTopicName,
		topics.MyLocationTopicName,
		topics.MeetingLocationTopicName,
	}
}
