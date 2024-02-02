package mailer

import (
	"log"
	"net/smtp"

	"github.com/aadi-1024/straysafe/internals/models"
)

//make it use a worker pool instead of a single worker

type Mailer struct {
	Mail chan models.MailData
	
	addr string
	auth smtp.Auth
}

func NewMailer(username string, pass string, host string) *Mailer {
	auth := smtp.PlainAuth("", "admin@example.com", "", host)
	defMail := "admin@example.com"

	return &Mailer{
		Mail: make(chan models.MailData),
		addr: defMail,
		auth: auth,
	}
}

func (m *Mailer) StartService() {
	for {
		mail := <-m.Mail
		err := smtp.SendMail("localhost:1025", m.auth, mail.From, mail.To, []byte(mail.Content))
		if err != nil {
			log.Println(err)
		}
	}
}