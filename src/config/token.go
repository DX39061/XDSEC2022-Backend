package config

import (
	"errors"
	"github.com/spf13/viper"
)

type Token struct {
	SigningKey  string `mapstructure:"signing_key" json:"signing-key" yaml:"signing_key"`
	ExpiresTime int64  `mapstructure:"expires_time" json:"expires-time" yaml:"expires_time"`
	BufferTime  int64  `mapstructure:"buffer_time" json:"buffer-time" yaml:"buffer_time"`
}

var TokenConfig Token

func init() {
	Register(&TokenConfig)
}

func (t *Token) Save() error {
	viper.Set("token", t)
	err := viper.WriteConfig()
	if err != nil {
		return err
	}
	return nil
}

func (t *Token) Load() error {
	configReader := viper.Sub("token")
	if configReader == nil {
		return errors.New("could not read token config")
	}
	err := configReader.Unmarshal(&t)
	if err != nil {
		return err
	}
	return nil
}
