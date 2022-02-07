package config

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type conf struct {
	ApiKey      string `yaml:"api-key"`
	HeroIdLower uint32 `yaml:"hero-id-lower"`
	HeroIdUpper uint32 `yaml:"hero-id-upper"`
}

var Params conf

func init() {
	if os.Getenv("ENV") == "test" {
		return
	}

	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(yamlFile, &Params)
	if err != nil {
		panic(err)
	}
}
