package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type ServerConfig struct {
	Port string `json:"port"`
}

func NewServerConfig() *ServerConfig {
	file, err := os.ReadFile("config.json")
	if err != nil {
		fmt.Println("Can't find file config.json! Using default port: 8080")
		return &ServerConfig{Port: "8080"}
	}

	config := &ServerConfig{}
	err = json.Unmarshal(file, config)
	if err != nil {
		return nil
	}

	return config
}
