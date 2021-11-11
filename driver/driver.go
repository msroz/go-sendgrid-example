package driver

import (
	"fmt"

	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type SendGridDriver struct {
	svc      *sendgrid.Client
	host     string
	endpoint string
	debug    bool
}

type SendGridOption func(*SendGridDriver)

func Host(v string) SendGridOption {
	return func(s *SendGridDriver) {
		s.host = v
	}
}

func Endpoint(v string) SendGridOption {
	return func(s *SendGridDriver) {
		s.endpoint = v
	}
}

func DebugMode(v bool) SendGridOption {
	return func(s *SendGridDriver) {
		s.debug = v
	}
}

func NewSendGrid(apiKey string, opts ...SendGridOption) *SendGridDriver {
	s := &SendGridDriver{
		endpoint: "/v3/mail/send",
	}

	for _, opt := range opts {
		opt(s)
	}

	request := sendgrid.GetRequest(apiKey, s.endpoint, s.host)
	request.Method = "POST"
	s.svc = &sendgrid.Client{
		Request: request,
	}

	if s.debug {
		fmt.Printf("%+v\n", s)
	}

	return s
}

func (d *SendGridDriver) Send(msg *mail.SGMailV3) (*rest.Response, error) {
	res, err := d.svc.Send(msg)
	if d.debug {
		fmt.Printf("%+v\n", res)

	}

	return res, err
}
