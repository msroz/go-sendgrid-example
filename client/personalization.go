package client

type Personalization struct {
	To                  []*Email
	From                *Email
	CC                  []*Email
	BCC                 []*Email
	Subject             string
	Substitutions       map[string]string
	DynamicTemplateData map[string]interface{}
}

func NewPersonalization() *Personalization {
	return &Personalization{
		Substitutions:       make(map[string]string),
		DynamicTemplateData: make(map[string]interface{}),
	}
}

func (p *Personalization) AddTos(to ...*Email) *Personalization {
	p.To = append(p.To, to...)
	return p
}

func (p *Personalization) AddCCs(cc ...*Email) *Personalization {
	p.CC = append(p.CC, cc...)
	return p
}

func (p *Personalization) AddBCCs(bcc ...*Email) *Personalization {
	p.BCC = append(p.BCC, bcc...)
	return p
}

func (p *Personalization) SetFrom(from *Email) *Personalization {
	p.From = from
	return p
}

func (p *Personalization) SetSubject(sub string) *Personalization {
	p.Subject = sub
	return p
}

func (p *Personalization) SetSubstitution(key, value string) *Personalization {
	p.Substitutions[key] = value
	return p
}

func (p *Personalization) SetDynamicTemplateData(key string, value interface{}) *Personalization {
	p.DynamicTemplateData[key] = value
	return p
}
