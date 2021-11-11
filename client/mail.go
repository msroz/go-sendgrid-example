package client

import (
	"gosendgridexample/driver"

	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type Mail struct {
	sendgrid *driver.SendGridDriver
}

func NewMail(s *driver.SendGridDriver) *Mail {
	return &Mail{
		sendgrid: s,
	}
}

func (c *Mail) Send(msg *Message) error {
	from := mail.NewEmail(msg.From.Name, msg.From.Address)
	to := mail.NewEmail(msg.To.Name, msg.To.Address)
	subject := msg.Subject
	message := mail.NewSingleEmail(from, subject, to, msg.PlainContent, msg.HTMLContent)

	p := mail.NewPersonalization()
	p.AddTos(to)
	p.AddCCs(mail.NewEmail(msg.CC.Name, msg.CC.Address))
	message.AddPersonalizations(p)

	_, err := c.sendgrid.Send(message)
	if err != nil {
		return err
	}
	return nil
}
