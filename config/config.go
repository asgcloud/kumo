package config

import (
	"encoding/json"
	"io/ioutil"
)

// Config contains the configuration information of the kumo backend.
type Config struct {
	Postgres struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		DB       string `json:"db"`
	} `json:"postgres"`
}

// FromFile returns a configuration object parsed from the given file.
func FromFile(path string) (*Config, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := json.Unmarshal(b, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
