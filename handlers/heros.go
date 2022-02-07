package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/super-hero-battle/config"
	"github.com/super-hero-battle/types"
)

const API_URL = "https://www.superheroapi.com/api"

var log = logrus.New()

func init() {
	log.Formatter = new(logrus.JSONFormatter)
}

func LoadHeroById(id uint32) (types.Hero, error) {
	var hero types.Hero

	resp, err := http.Get(fmt.Sprintf("%s/%s/%d", API_URL, config.Params.ApiKey, id))
	if err != nil {
		log.Fatal(err)
		return hero, err
	}

	err = json.NewDecoder(resp.Body).Decode(&hero)
	if err != nil {
		log.Fatal(err)
		return hero, err
	}
	return hero, nil
}
