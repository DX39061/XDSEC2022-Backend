package config

import (
	"errors"
	"github.com/spf13/viper"
)

type Server struct {
	ApiBasePath string `mapstructure:"api_base_path" json:"api-base-path" yaml:"api_base_path"`
	Address     string `mapstructure:"address" json:"address" yaml:"address"`
	Port        int    `mapstructure:"port" json:"port" yaml:"port"`
}

var ServerConfig Server

func init() {
	Register(&ServerConfig)
}

func (s *Server) Save() error {
	viper.Set("server", s)
	err := viper.WriteConfig()
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) Load() error {
	configReader := viper.Sub("server")
	if configReader == nil {
		return errors.New("could not read server config")
	}
	err := configReader.Unmarshal(&s)
	if err != nil {
		return err
	}
	return nil
}
