package config

import (
	"errors"
	"github.com/spf13/viper"
)

type Captcha struct {
	SecretKey string `mapstructure:"secret_key" json:"secret-key" yaml:"secret_key"`
}

var CaptchaConfig Captcha

func init() {
	Register(&CaptchaConfig)
}

func (s *Captcha) Save() error {
	viper.Set("captcha", s)
	err := viper.WriteConfig()
	if err != nil {
		return err
	}
	return nil
}

func (s *Captcha) Load() error {
	configReader := viper.Sub("captcha")
	if configReader == nil {
		return errors.New("could not read captcha config")
	}
	err := configReader.Unmarshal(&s)
	if err != nil {
		return err
	}
	return nil
}
