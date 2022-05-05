package chapter05

import (
	"encoding/json"
	"fmt"
	"os"
)

type loadConfigError struct {
	msg string
	err error
}

func (e *loadConfigError) Error() string {
	return fmt.Sprintf("cannot load config: %s (%s)", e.msg, e.err.Error())
}

func (e *loadConfigError) Unwrap() error {
	return e.err
}

type Config struct{}

func LoadConfig(configFilePath string) (*Config, error) {
	var cfg *Config
	data, err := os.ReadFile(configFilePath)
	if err != nil {
		return nil, &loadConfigError{
			msg: fmt.Sprintf("read file `%s`", configFilePath),
			err: err,
		}
	}

	if err = json.Unmarshal(data, cfg); err != nil {
		return nil, &loadConfigError{
			msg: fmt.Sprintf("parse config file `%s`", configFilePath),
			err: err,
		}
	}

	return cfg, nil
}
