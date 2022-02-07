package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type conf struct {
	ApiKey      string `yaml:"api-key"`
	HeroIdLower string `yaml:"hero-id-lower"`
	HeroIdUpper string `yaml:"hero-id-upper"`
}

var Params conf

func init() {
	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(yamlFile, &Params)
	if err != nil {
		panic(err)
	}
}
