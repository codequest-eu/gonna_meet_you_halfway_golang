package storage

import "github.com/codequest-eu/gonna_meet_you_halfway_golang/models"

type Store interface {
	SaveTopics(meetingID string, topics models.Topics) error
}
