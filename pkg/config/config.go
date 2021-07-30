package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var C *Config

type Config struct {
	TraderKey          string `json:"TraderKey"`
	Env                string `json:"ENV"`
	GoogleClientId     string `json:"GoogleClientId"`
	GoogleClientSecret string `json:"GoogleClientSecret"`
	HOST               string `json:"Host"`
	REDIS_HOST         string `json:"RedisHost"`
	REDIS_PASSWORD     string `json:"RedisPassword"`
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
