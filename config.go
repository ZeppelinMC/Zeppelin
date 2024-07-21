package main

import (
	"os"

	"github.com/pelletier/go-toml"
	"github.com/zeppelinmc/zeppelin/server/config"
)

func loadConfig() config.ServerConfig {
	file, err := os.Open("config.toml")
	if err != nil {
		defer file.Close()
		file, err = os.Create("config.toml")
		if err == nil {
			toml.NewEncoder(file).CompactComments(true).Encode(config.DefaultConfig)
			file.Close()
		}
		return config.DefaultConfig
	}
	var cfg config.ServerConfig
	err = toml.NewDecoder(file).Decode(&cfg)
	if err != nil {
		cfg = config.DefaultConfig
	}

	return cfg
}
