package main

import (
	"os"

	"github.com/dynamitemc/aether/server"
	"github.com/pelletier/go-toml"
)

func loadConfig() server.ServerConfig {
	file, err := os.Open("config.toml")
	if err != nil {
		defer file.Close()
		file, err = os.Create("config.toml")
		if err == nil {
			toml.NewEncoder(file).Encode(server.DefaultConfig)
			file.Close()
		}
		return server.DefaultConfig
	}
	var cfg server.ServerConfig
	err = toml.NewDecoder(file).Decode(&cfg)
	if err != nil {
		cfg = server.DefaultConfig
	}

	return cfg
}
