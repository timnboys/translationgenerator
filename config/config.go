package config

import "github.com/BurntSushi/toml"

type Config struct {
	Server struct{
		Host string
	}
}

func LoadConfig() (conf Config) {
	if _, err := toml.DecodeFile("config.toml", &conf); err != nil {
		panic(err)
	}

	return
}
