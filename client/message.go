package client

type Message struct {
	From         *Email
	To           *Email
	CC           *Email
	Subject      string
	PlainContent string
	HTMLContent  string
}

type Email struct {
	Name    string
	Address string
}

func NewMessage(subject string) *Message {
	return &Message{
		Subject: subject,
	}
}

func NewEmail(name, address string) *Email {
	return &Email{
		Name:    name,
		Address: address,
	}
}

func (m *Message) AddPlaintContent(plain string) *Message {
	m.PlainContent = plain
	return m
}

func (m *Message) AddHTMLContent(html string) *Message {
	m.HTMLContent = html
	return m
}

func (m *Message) AddFrom(email *Email) *Message {
	m.From = email
	return m
}

func (m *Message) AddTo(email *Email) *Message {
	m.To = email
	return m
}

func (m *Message) AddCC(email *Email) *Message {
	m.CC = email
	return m
}
