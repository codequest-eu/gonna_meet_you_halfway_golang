package mailer

type Mailer interface {
	Mail(tEmail string, sName string, sEmail string) error
}
