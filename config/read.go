package config

import (
	"encoding/json"
	"os"
)

func ReadConfig(configPath string) (Config, error) {
	var c Config
	all, err := os.ReadFile(configPath)
	if err != nil {
		return c, err
	}

	return c, json.Unmarshal(all, &c)
}

func MustReadConfig(configPath string) Config {
	c, err := ReadConfig(configPath)
	if err != nil {
		panic(err)
	}
	return c
}
