package config

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	SEVER SeverConfig
	DB    DBConfig
}

type SeverConfig struct {
	Host string
	Port uint
}

type DBConfig struct {
	Connection string
}

func NewConfig() Config {
	var config Config

	_, err := toml.DecodeFile("config.toml", &config)
	if err != nil {
		// Error Handling
	}

	return config
}
