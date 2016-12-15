package storage

import "github.com/codequest-eu/gonna_meet_you_halfway_golang/models"

type Store interface {
	SaveTopics(meetingID string, topics models.Topics) error
	GetTopics(meetingID string) (models.Topics, error)

	SaveMeetingSuggestion(meetingID string, meetingSuggestion models.MeetingSuggestion) error
	GetMeetingSuggestion(meetingID string) (models.MeetingSuggestion, error)

	SavePlaceSuggestion(meetingID string, placeSuggestion models.PlaceSuggestion) error
	GetPlaceSuggestion(placeSuggestionID string) (models.PlaceSuggestion, error)
}
