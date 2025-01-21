package config

import "time"

type Config struct {
	LogInterval  time.Duration
	LogFormat    string
	LogDirectory string
}

func DefaultConfig() *Config {
	return &Config{
		LogInterval:  time.Nanosecond,
		LogFormat:    ".log",
		LogDirectory: "./logs/",
	}
}
