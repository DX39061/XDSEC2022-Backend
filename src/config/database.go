package config

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
)

// Database is the configuration of MySQL database.
type Database struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Port     string `mapstructure:"port" json:"port" yaml:"port"`
	User     string `mapstructure:"user" json:"user" yaml:"user"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	DBName   string `mapstructure:"db_name" json:"db-name" yaml:"db_name"`
	SSLMode  string `mapstructure:"ssl_mode" json:"ssl-mode" yaml:"ssl_mode"`
	TimeZone string `mapstructure:"time_zone" json:"time-zone" yaml:"time_zone"`
}

func (d *Database) Dsn() string {
	// MySQL connection
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True&loc=Local",
		d.User, d.Password, d.Host, d.Port, d.DBName)
}

var DatabaseConfig Database

func init() {
	Register(&DatabaseConfig)
}

func (d *Database) Save() error {
	viper.Set("database", d)
	err := viper.WriteConfig()
	if err != nil {
		return err
	}
	return nil
}

func (d *Database) Load() error {
	configReader := viper.Sub("database")
	if configReader == nil {
		return errors.New("could not read database config")
	}
	err := configReader.Unmarshal(&d)
	if err != nil {
		return err
	}
	return nil
}
