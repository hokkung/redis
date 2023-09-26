package config

import "github.com/kelseyhightower/envconfig"

const APP_PREFIX = "APP_REDIS"

type Config struct {
	Addr       string `envconfig:"ADDR" default:"localhost:6379"`
	Password   string `envconfig:"PASSWORD" default:""`
	DB         int    `envconfig:"DB" default:"0"`
	MaxRetries int    `envconfig:"MAX_RETRIES" default:"3"`
}

func New() *Config {
	var config Config
	envconfig.MustProcess(APP_PREFIX, &config)
	return &config
}
