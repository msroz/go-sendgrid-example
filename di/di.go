package di

import (
	"gosendgridexample/client"
	"gosendgridexample/config"
	"gosendgridexample/driver"
)

func InjectMailClient() *client.Mail {
	return client.NewMail(
		InjectDriver(),
	)
}

func InjectDriver() *driver.SendGridDriver {
	cfg := config.GetConfig()

	return driver.NewSendGrid(
		cfg.SendGrid.APIKey,
		driver.DebugMode(cfg.SendGrid.Debug),
		driver.Host(cfg.SendGrid.Host),
	)
}
