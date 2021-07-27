package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	TraderKey string `json:"TraderKey"`
	Env       string `json:"ENV"`
}

var DefaultConfigFileName = "config.json"

func GetConfig() *Config {
	env, ok := os.LookupEnv("ENV")
	var config Config
	data, _ := ioutil.ReadFile(DefaultConfigFileName)
	_ = json.Unmarshal(data, &config)
	if !ok {
		env = "Dev"
	}
	config.Env = env
	return &config
}
