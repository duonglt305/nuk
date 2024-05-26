package email

import (
	"fmt"
	"log"
	"net/smtp"
)

type EmailSender interface {
	Send(to string, subject string, body []byte) error
}

type SMTPSender struct {
	from     string
	host     string
	port     int
	user     string
	password string
}

type SMTPOptions struct {
	Host     string
	Port     int
	User     string
	Password string
	From     string
}

func NewSMTPSender(opts SMTPOptions) EmailSender {
	return &SMTPSender{
		host:     opts.Host,
		port:     opts.Port,
		user:     opts.User,
		password: opts.Password,
		from:     opts.From,
	}
}

func (s *SMTPSender) Send(to string, subject string, body []byte) error {
	auth := smtp.PlainAuth("", s.user, s.password, s.host)
	server := fmt.Sprintf("%s:%d", s.host, s.port)
	err := smtp.SendMail(server, auth, s.from, []string{to}, body)
	if err != nil {
		log.Printf("failed to send email: %+v\n", err)
		return err
	}
	return nil
}
