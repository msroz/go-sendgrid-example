# 📧 go-sendgrid-example

Example of sending mail via SendGrid in Golang.

## Get it started 

```
$ make setup

# Edit environment variables
$ vim ./env/local.env

# Run docker containers
$ make up

# Healthcheck of mock-sendgrid
$ make ping

# Send email
$ docker-compose exec app go run main.go

# Send email with options
$ docker-compose exec app go run main.go -to=hoge@example.com -cc=fuga@example.com -subject=hi
```

### Mail

Use [ykanazawa/sendgrid-maildev](https://hub.docker.com/r/ykanazawa/sendgrid-maildev) as mock SendGrid API Server.

- Mock SendGrid API Server running at `http://0.0.0.0:3030`
- MailDev webapp running at `http://0.0.0.0:1080`
- MailDev SMTP Server running at `0.0.0.0:1025`
