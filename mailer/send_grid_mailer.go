package mailer

import (
	"log"
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

func (sgm *sendGridMailer) Mail(tEmail string, sName string, sEmail string) error {
	request := sgm.request
	request.Body = sgm.buildInviteMail(tEmail, sName, sEmail)
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
		return  err
	}
	log.Println("Mail from" + sEmail + " to " + tEmail + " sent with: " + response.Body)
	return nil
}

func (sgm *sendGridMailer) buildInviteMail(tEmail string, sName string, sEmail string) []byte {
	from := mail.NewEmail("[Half Way]", "support@halfway.io")
	to := mail.NewEmail("", tEmail)
	content := mail.NewContent("text/html", "<p>Your friend " + sName + " (" + sEmail + ") want you to meet with you.</p><p>Use this <a href= >link</a> find best place to see him/her</p><br />")
	subject := "[Half Way] Let's meet!!!"
	m := mail.NewV3MailInit(from, subject, to, content)
	return mail.GetRequestBody(m)
}
