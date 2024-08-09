package main

import (
	"os"

	"github.com/zeppelinmc/zeppelin/properties"
)

func loadConfig() properties.ServerProperties {
	file, err := os.ReadFile("server.properties")
	if err != nil {
		file, err := os.Create("server.properties")
		if err == nil {
			properties.Marshal(file, properties.Default)
			file.Close()
		}
		return properties.Default
	}
	var cfg properties.ServerProperties

	err = properties.Unmarshal(string(file), &cfg)
	if err != nil {
		cfg = properties.Default
	}

	return cfg
}
