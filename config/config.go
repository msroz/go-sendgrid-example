package config

import (
	"os"
	"strconv"
	"sync"
)

var config *Config
var onece sync.Once

type Config struct {
	AppName  string
	SendGrid SendGrid
}

type SendGrid struct {
	APIKey   string
	Host     string
	Endpoint string
	Debug    bool
}

func GetConfig() *Config {

	onece.Do(func() {
		config = NewConfig()
	})

	return config
}

func NewConfig() *Config {
	c := &Config{
		AppName: "goexamplesendgrid",
	}
	loadConfig(c)

	return c
}

func loadConfig(c *Config) {
	loadSendGridConfig(c)
}

func loadSendGridConfig(c *Config) {
	c.SendGrid.APIKey = os.Getenv("SENDGRID_API_KEY")
	c.SendGrid.Host = os.Getenv("SENDGRID_HOST")
	c.SendGrid.Endpoint = os.Getenv("SENDGRID_ENDPOINT")
	// TODO: error handling :(
	v, _ := strconv.ParseBool(os.Getenv("SENDGRID_DEBUG"))
	c.SendGrid.Debug = v
}
