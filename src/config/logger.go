package config

import (
	"errors"
	"github.com/spf13/viper"
)

type Logger struct {
	Level     string `mapstructure:"level" json:"level" yaml:"level"`
	Format    string `mapstructure:"format" json:"format" yaml:"format"`
	Directory string `mapstructure:"directory" json:"directory"  yaml:"directory"`
	MaxAge    int64  `mapstructure:"max_age" json:"max-age" yaml:"max_age"`
	// LinkName is the name of the symlink to the current log file.
	LinkName      string `mapstructure:"link_name" json:"link-name" yaml:"link_name"`
	ShowLine      bool   `mapstructure:"show_line" json:"show-line" yaml:"show_line"`
	EncodeLevel   string `mapstructure:"encode_level" json:"encode-level" yaml:"encode_level"`
	StacktraceKey string `mapstructure:"stacktrace_key" json:"stacktrace-key" yaml:"stacktrace_key"`
	LogInConsole  bool   `mapstructure:"log_in_console" json:"log-in-console" yaml:"log_in_console"`
}

var LoggerConfig Logger

func init() {
	Register(&LoggerConfig)
}

func (l *Logger) Save() error {
	viper.Set("logger", l)
	err := viper.WriteConfig()
	if err != nil {
		return err
	}
	return nil
}

func (l *Logger) Load() error {
	configReader := viper.Sub("logger")
	if configReader == nil {
		return errors.New("could not read logger config")
	}
	err := configReader.Unmarshal(&l)
	if err != nil {
		return err
	}
	return nil
}
