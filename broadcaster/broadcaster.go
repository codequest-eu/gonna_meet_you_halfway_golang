package broadcaster

import (
	"io"

	"github.com/codequest-eu/gonna_meet_you_halfway_golang/models"
)

//Broadcaster is responsible for open a communication channel and message exchange
type Broadcaster interface {
	io.Closer
	PublishMeetingSuggestion(sugestion models.MeetingSuggestion, topic string) error
}
