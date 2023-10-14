package emailConf

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Address string `yaml:"address"`
}

func GetEmailConfig() Config {
	var config Config

	configData, err := os.ReadFile("./internal/config/emailConf/conf.yaml")
	if err != nil {
		panic(err)
	}

	if err := yaml.Unmarshal(configData, &config); err != nil {
		panic(err)
	}

	return config
}
