package client

import "errors"

type ContentType string

const (
	ContentTypePlain ContentType = "text/plain"
	ContentTypeHTML  ContentType = "text/html"
)

// See: https://sendgrid.api-docs.io/v3.0/mail-send/v3-mail-send

type Message struct {
	From             *Email
	ReplyTo          *Email
	Subject          string // overridden by subject lines set in personalizations.
	Personalizations []*Personalization
	Content          []*Content
	TemplateID       string
	// Attachments
}

type Content struct {
	Type  ContentType
	Value string
}

type Email struct {
	Name    string
	Address string
}

func NewMessage(subject string, content ...*Content) *Message {
	m := new(Message)
	m.Subject = subject
	m.AddContents(content...)

	return m
}

func NewContent(typ ContentType, val string) *Content {
	return &Content{
		Type:  typ,
		Value: val,
	}
}

func NewEmail(name, address string) *Email {
	return &Email{
		Name:    name,
		Address: address,
	}
}

func (m *Message) Validate() error {
	if m.From == nil {
		return errors.New("from object required")
	}

	if m.Subject == "" {
		return errors.New("subject should not be empty string")
	}

	if m.TemplateID == "" && len(m.Content) < 1 {
		return errors.New("at leaset one content required")
	}

	for _, p := range m.Personalizations {
		if len(p.To) < 1 {
			return errors.New("substitutions should not be used with templateID")
		}
	}

	if m.TemplateID != "" {
		for _, p := range m.Personalizations {
			for range p.Substitutions {
				return errors.New("substitutions should not be used with templateID")
			}
		}
	}

	return nil
}

func (m *Message) AddPersonalizations(p ...*Personalization) *Message {
	m.Personalizations = append(m.Personalizations, p...)
	return m
}

func (m *Message) AddContents(content ...*Content) *Message {
	m.Content = append(m.Content, content...)
	return m
}

func (m *Message) SetFrom(email *Email) *Message {
	m.From = email
	return m
}

func (m *Message) SetTemplateID(id string) *Message {
	m.TemplateID = id
	return m
}
