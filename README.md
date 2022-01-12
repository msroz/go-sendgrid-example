# 📧 go-sendgrid-example

Example of sending mail via SendGrid in Golang.

## 🦾 Get it started

### 📦 Run containers

```
$ make setup

# Edit environment variables
$ vim ./env/local.env

# Run docker containers
$ make up
```

### 📮 Send mails

```
# Healthcheck of mock-sendgrid
$ make ping

# Send a single email with default value
$ docker-compose exec app go run cmd/single/single.go

# Send a single email with some options
$ docker-compose exec app go run cmd/single/single.go -to=hoge@example.com -cc=fuga@example.com -subject=hi

# Send a single email with substitusions
$ docker-compose exec app go run cmd/substitution/substitution.go -to=foo@example.com -to=bar@example.com -to=baz@example.com -cc=fuga@example.com -subject=hi

# Send a single email with Dynamic Transactional Templates
# NOTE: This command does not work with sendgrid-maildev mock server because of no support for dynamic transactional templates. Use "real" sendgrid API instead.
# See: https://docs.sendgrid.com/ui/sending-email/how-to-send-an-email-with-dynamic-transactional-templates
$ docker-compose exec app go run cmd/template/template.go -to=hoge@example.com -cc=fuga@example.com -subject=hi -templateID=[Your template ID goes here.]
```

### 📥 Mail in mock server

Use [ykanazawa/sendgrid-maildev](https://hub.docker.com/r/ykanazawa/sendgrid-maildev) as mock SendGrid API Server.

- Mock SendGrid API Server running at `http://0.0.0.0:3030`
- MailDev webapp running at `http://0.0.0.0:1080`
- MailDev SMTP Server running at `0.0.0.0:1025`

## 🔗 References

- [SendGrid v3 API Documentation](https://sendgrid.api-docs.io/v3.0/mail-send/v3-mail-send)
- [v3 Mail Send API概要 - ドキュメント | SendGrid](https://sendgrid.kke.co.jp/docs/API_Reference/Web_API_v3/Mail/index.html)
- :octocat: [sendgrid/sendgrid-go: The Official Twilio SendGrid Led, Community Driven Golang API Library](https://github.com/sendgrid/sendgrid-go)
- :octocat: [yKanazawa/sendgrid-maildev: SendGrid MailDev is SengGrid mock API + MailDev. For test your sendgrid emails during development.](https://github.com/yKanazawa/sendgrid-maildev)
- :octocat: [yKanazawa/sendgrid-dev: SendGrid Dev is SengGrid mock API for test your sendgrid emails during development.](https://github.com/yKanazawa/sendgrid-dev)