package config

import (
	"encoding/json"
	"log"
	"os"
	pkg_os "turtle-service/pkg/os"
)

type Config struct {
	Port int `json:"port"`
}

func New() *Config {
	var cfg *Config
	data, err := os.ReadFile(pkg_os.ProcessPath("config.json"))
	if err != nil {
		log.Fatal(err)
	}

	if err = json.Unmarshal(data, &cfg); err != nil {
		log.Fatal(err)
	}

	return cfg
}
