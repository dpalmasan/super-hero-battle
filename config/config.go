package config

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type emailConf struct {
	To              string `yaml:"to"`
	MailgunUsername string `yaml:"mailgun-username"`
	MailgunPassword string `yaml:"mailgun-password"`
}

type conf struct {
	ApiKey      string    `yaml:"api-key"`
	HeroIdLower uint32    `yaml:"hero-id-lower"`
	HeroIdUpper uint32    `yaml:"hero-id-upper"`
	Email       emailConf `yaml:"email"`
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
