package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	TraderKey string
}

var DefaultConfigFile = "config.json"

func GetConfig() Config {
	var config Config
	data, _ := ioutil.ReadFile(DefaultConfigFile)
	_ = json.Unmarshal(data, &config)
	return config
}
