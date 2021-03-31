// Package mail defines smtp client to send emails
package mail

import (
	"fmt"
	"net/smtp"
)

// GoMailer defines methods on mailer client
type GoMailer interface {
	SendMail(toAddresses []string, message string) error
}

type mailer struct {
	host     string
	port     int
	fromMail string
	password string
}

// NewMailer creates new instance of maielr
func NewMailer(host string, port int, fromMail, password string) GoMailer {
	return &mailer{
		host:     host,
		port:     port,
		fromMail: fromMail,
		password: password,
	}
}

func (m *mailer) address() string {
	return fmt.Sprintf("%s:%d", m.host, m.port)
}

func (m *mailer) SendMail(toAddresses []string, message string) error {
	auth := smtp.PlainAuth("", m.fromMail, m.password, m.host)
	return smtp.SendMail(m.address(), auth, m.fromMail, toAddresses, []byte(message))
}
