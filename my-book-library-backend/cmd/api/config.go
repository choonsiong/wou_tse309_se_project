package main

import (
	"encoding/json"
	"os"
)

// Config type store application wide configuration
type Config struct {
	DSN      string `json:"dsn"`       // database source name
	ENV      string `json:"env"`       // running environment (e.g. dev, prod)
	HttpPort int    `json:"http_port"` // http server listening port
}

const (
	configFilePath = "etc/config.json"
)

// loadConfig loads the user configuration from the config file.
func loadConfig(cfg *Config) error {
	bs, err := os.ReadFile(configFilePath)
	if err != nil {
		return err
	}

	return json.Unmarshal(bs, &cfg)
}
