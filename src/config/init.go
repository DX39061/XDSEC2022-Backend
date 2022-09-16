package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config interface {
	Save() error
	Load() error
}

var configModels []Config

func Register(config Config) {
	configModels = append(configModels, config)
}

func Initialize() error {
	if err := initViper(); err != nil {
		log.Println("Could not init viper.")
		return err
	}
	return nil
}

func initViper() error {
	viper.SetConfigName("config")            // name of config file (without extension)
	viper.SetConfigType("yaml")              // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("/etc/XDSEC-join/")  // path to look for the config file in
	viper.AddConfigPath("$HOME/.XDSEC-join") // call multiple times to add many search paths
	viper.AddConfigPath(".")                 // optionally look for config in the working directory
	if err := viper.ReadInConfig(); err != nil {
		log.Println("Could not access config file, maybe it is not exist?", err)
		return err
	}
	err := readConfig()
	if err != nil {
		log.Println("Could not read config file, maybe it is broken?", err)
		return err
	}
	return nil
}

func readConfig() error {
	for _, model := range configModels {
		err := model.Load()
		if err != nil {
			return err
		}
	}
	return nil
}
