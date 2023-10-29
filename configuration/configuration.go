package configuration

import (
	"os"

	"gopkg.in/yaml.v2"
)

const (
	CONFIG_PATH = "configuration/config.yml"
)

func GetConfig() *Configuration {
	config := Configuration{}

	content, err := os.ReadFile(CONFIG_PATH)
	if err != nil {
		return setDefaultConfig(&config)
	}

	err = yaml.Unmarshal(content, &config)
	if err != nil {
		panic(err)
	}

	return &config
}
