package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/super-hero-battle/config"
	"github.com/super-hero-battle/types"
)

const API_URL = "https://www.superheroapi.com/api"

var log = logrus.New()

func init() {
	log.Formatter = new(logrus.JSONFormatter)
	rand.Seed(time.Now().UnixNano())
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

func GenerateRandomIds(n uint32, lower uint32, upper uint32) ([]uint32, error) {
	size := upper - lower + 1
	log.Debugf("Sampling from a slice of size %d", size)
	if size < n {
		return nil, errors.New("Cannot sample from the given parameters")
	}
	ids := make([]uint32, size)

	var i uint32
	for i = 0; i < uint32(len(ids)); i++ {
		ids[i] = config.Params.HeroIdLower + i
	}
	rand.Shuffle(len(ids), func(i, j int) {
		ids[i], ids[j] = ids[j], ids[i]
	})

	output := make([]uint32, n)

	for i = 0; i < uint32(len(output)); i++ {
		output[i] = ids[i]
	}
	return output, nil
}

func BuildHeroTeam(heroes [5]types.Hero) types.Team {
	goodAlignment := 0

	for _, hero := range heroes {
		if hero.Biography.Alignment == "good" {
			goodAlignment += 1
		}
	}
	teamAlignment := "bad"

	if goodAlignment >= 3 {
		teamAlignment = "good"
	}

	team := types.Team{
		Heroes:    heroes,
		Alignment: teamAlignment,
	}

	return team
}
