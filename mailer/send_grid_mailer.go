package mailer

import (
	"log"

	"github.com/codequest-eu/gonna_meet_you_halfway_golang/models"
	"github.com/codequest-eu/gonna_meet_you_halfway_golang/util"
	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

const (
	sendMailEndpoint = "/v3/mail/send"
	sendMailHost     = "https://api.sendgrid.com"
)

type sendGridMailer struct {
	request rest.Request
}

func NewSendGridMailer(key string) Mailer {
	r := sendgrid.GetRequest(key, sendMailEndpoint, sendMailHost)
	r.Method = "POST"
	return &sendGridMailer{r}
}

func (sgm *sendGridMailer) Mail(inviteData models.InviteData, meetingID string) error {
	request := sgm.request
	sEmail := inviteData.Email
	tEmail := inviteData.OtherEmail
	sName := inviteData.Name
	request.Body = sgm.buildInviteMail(sEmail, sName, tEmail, meetingID)
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("Mail from" + sEmail + " to " + tEmail + " sent with: " + response.Body)
	return nil
}

func (sgm *sendGridMailer) buildInviteMail(sEmail string, sName string, tEmail string, meetingID string) []byte {
	from := mail.NewEmail("[Half Way]", "support@halfway.io")
	to := mail.NewEmail("", tEmail)
	invitePath := util.InvitePath(meetingID)
	content := mail.NewContent("text/html", "<p>Your friend "+sName+" ("+sEmail+") want you to meet with you.</p><p>Use this <a href="+invitePath+">link</a> find best place to see him/her</p><br />")
	subject := "[Half Way] Let's meet!!!"
	m := mail.NewV3MailInit(from, subject, to, content)
	return mail.GetRequestBody(m)
}
