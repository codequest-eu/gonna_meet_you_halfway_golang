package mailer

import "github.com/codequest-eu/gonna_meet_you_halfway_golang/models"

type Mailer interface {
	Mail(inviteData models.InviteData, meetingID string) error
}
