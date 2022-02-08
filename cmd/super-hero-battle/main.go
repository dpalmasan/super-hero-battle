package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/mailgun/mailgun-go/v4"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"github.com/super-hero-battle/config"
	"github.com/super-hero-battle/handlers"
	"github.com/super-hero-battle/types"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	lvl, ok := os.LookupEnv("LOG_LEVEL")
	// LOG_LEVEL not set, let's default to debug
	if !ok {
		lvl = "debug"
	}
	// parse string, this is built-in feature of logrus
	ll, err := logrus.ParseLevel(lvl)
	if err != nil {
		ll = logrus.DebugLevel
	}
	// set global log level
	logrus.SetLevel(ll)
}

func main() {

	ids, err := handlers.GenerateRandomIds(10, config.Params.HeroIdLower, config.Params.HeroIdUpper)
	if err != nil {
		log.Fatal("Could not sample random ids")
		return
	}
	log.Debugf("Generated ids: %+v", ids)

	var heroes1 [5]types.Hero
	var heroes2 [5]types.Hero

	k1 := 0
	k2 := 0
	for i, id := range ids {
		hero, err := handlers.LoadHeroById(id)
		log.Debugf("Loaded hero: %+v", hero)
		if err != nil {
			log.Fatal(err)
			return
		}
		if i < 5 {
			heroes1[k1] = hero
			k1++
		} else {
			heroes2[k2] = hero
			k2++
		}
	}

	team1 := types.BuildHeroTeam(heroes1)
	team1.UpdateFiliationCoefficient()
	team2 := types.BuildHeroTeam(heroes2)
	team2.UpdateFiliationCoefficient()

	log.Debug("Team 1: %+v", team1)
	log.Debug("Team 2: %+v", team2)

	battleFinished := false
	team1Pointer := 0
	team2Pointer := 0
	team1Turn := rand.Float32() > 0.5
	hero1Hp := team1.Heroes[0].HP
	hero2Hp := team2.Heroes[0].HP
	output := ""
	var dmg float32
	for !battleFinished {
		output += fmt.Sprintf("Team 1\nCurrent Hero: %s (%f HP)\n", team1.Heroes[team1Pointer].Name, hero1Hp)
		output += fmt.Sprintf("Team 2\nCurrent Hero: %s (%f HP)\n", team2.Heroes[team2Pointer].Name, hero2Hp)
		if team1Turn {
			dmg = team1.Heroes[team1Pointer].Attack()
			hero2Hp -= dmg
			output += fmt.Sprintf("Hero %s received %f damage!\n", team2.Heroes[team2Pointer].Name, dmg)
		} else {
			dmg = team2.Heroes[team2Pointer].Attack()
			hero1Hp -= dmg
			output += fmt.Sprintf("Hero %s received %f damage!\n", team1.Heroes[team1Pointer].Name, dmg)
		}

		if hero1Hp <= 0 {
			output += fmt.Sprintf("Hero %s from team 1 died!\n", team1.Heroes[team1Pointer].Name)
			team1Pointer++

			if team1Pointer < 5 {
				hero1Hp = team1.Heroes[team1Pointer].HP
			}
			hero2Hp = team1.Heroes[team2Pointer].HP
		}

		if hero2Hp <= 0 {
			output += fmt.Sprintf("Hero %s from team 2 died!\n", team2.Heroes[team2Pointer].Name)
			team2Pointer++
			hero1Hp = team1.Heroes[team1Pointer].HP

			if team2Pointer < 5 {
				hero2Hp = team1.Heroes[team2Pointer].HP
			}
		}
		team1Turn = !team1Turn
		battleFinished = team1Pointer >= 5 || team2Pointer >= 5
	}

	var won int
	if team1Pointer >= 5 {
		won = 2
	} else {
		won = 1
	}

	output += fmt.Sprintf("Team %d won the battle!\n", won)
	fmt.Print(output)

	mg := mailgun.NewMailgun(config.Params.Email.ApiDomain, config.Params.Email.ApiKey)

	sender := "super-hero-battle@test.io"
	subject := "Epic Hero Battle!!"
	body := output
	recipient := config.Params.Email.To

	// The message object allows you to add attachments and Bcc recipients
	message := mg.NewMessage(sender, subject, body, recipient)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	// Send the message with a 10 second timeout
	_, _, err = mg.Send(ctx, message)

	if err != nil {
		log.Fatal(err)
	}
}
