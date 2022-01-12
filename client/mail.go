package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gosendgridexample/driver"

	"github.com/pkg/errors"
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
	if err := msg.Validate(); err != nil {
		return errors.Wrap(err, "validate message")
	}

	from := mail.NewEmail(msg.From.Name, msg.From.Address)
	subject := msg.Subject

	message := mail.NewV3MailInit(from, subject, &mail.Email{})

	if msg.TemplateID != "" {
		message.SetTemplateID(msg.TemplateID)
	}

	for _, c := range msg.Content {
		message.AddContent(mail.NewContent(string(c.Type), c.Value))
	}

	var personalizations []*mail.Personalization
	for _, pe := range msg.Personalizations {
		p := mail.NewPersonalization()

		for _, to := range pe.To {
			t := mail.NewEmail(to.Name, string(to.Address))
			p.AddTos(t)
		}
		for _, cc := range pe.CC {
			t := mail.NewEmail(cc.Name, string(cc.Address))
			p.AddCCs(t)
		}

		for _, bcc := range pe.BCC {
			t := mail.NewEmail(bcc.Name, string(bcc.Address))
			p.AddBCCs(t)
		}

		for k, v := range pe.Substitutions {
			p.SetSubstitution(k, v)
		}

		for k, v := range pe.DynamicTemplateData {
			p.SetDynamicTemplateData(k, v)
		}

		personalizations = append(personalizations, p)
	}

	// Overwrite `Personalications`,
	// because only one To address can be included in `Personalization`
	// which is initialized in `NewSingleEmail` method.
	message.Personalizations = personalizations

	b := mail.GetRequestBody(message)
	j, _ := prettyJSON(b)
	fmt.Printf("Request Body:%s\n", j)

	_, err := c.sendgrid.Send(message)
	if err != nil {
		return err
	}
	return nil
}

// For debug
func prettyJSON(b []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "    ")
	if err != nil {
		return b, err
	}
	return out.Bytes(), nil
}
