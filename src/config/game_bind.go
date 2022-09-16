package config

import (
	"errors"
	"github.com/spf13/viper"
)

type GameBind struct {
	PlatformUrl string `mapstructure:"platform_url" json:"platform-url" yaml:"platform_url"`
	GameID      int    `mapstructure:"game_id" json:"game-id" yaml:"game_id"`
	AuthToken   string `mapstructure:"auth_token" json:"auth-token" yaml:"auth_token"`
}

var GameBindConfig GameBind

func init() {
	Register(&GameBindConfig)
}

func (g *GameBind) Save() error {
	viper.Set("game_bind", g)
	err := viper.WriteConfig()
	if err != nil {
		return err
	}
	return nil
}

func (g *GameBind) Load() error {
	configReader := viper.Sub("game_bind")
	if configReader == nil {
		return errors.New("could not read game_bind config")
	}
	err := configReader.Unmarshal(&g)
	if err != nil {
		return err
	}
	return nil
}
