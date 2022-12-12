package config

// Cache is the configuration of the redis cache server.
import (
	"errors"

	"github.com/spf13/viper"
)

type Email struct {
	Smtp_url   string `yaml:"smtp_url"`
	Smtp_port       int    `yaml:"smtp_port"`
	Email_from     string `yaml:"email_from"`
	Email_password string `yaml:"email_password"`
}

var EmailConfig Email

func init() {
	Register(&EmailConfig)
}

func (c *Email) Save() error {
	viper.Set("email", c)
	err := viper.WriteConfig()
	if err != nil {
		return err
	}
	return nil
}

func (e *Email) Load() error {
	configReader := viper.Sub("email")
	if configReader == nil {
		return errors.New("could not read email config")
	}
	err := configReader.Unmarshal(&e)
	if err != nil {
		return err
	}
	return nil
}
