package types

import "math/rand"

type Stats struct {
	Intelligence  uint32 `json:"intelligence,string,omitempty"`
	Strength      uint32 `json:"strength,string,omitempty"`
	Speed         uint32 `json:"speed,string,omitempty"`
	Durability    uint32 `json:"durability,string,omitempty"`
	Power         uint32 `json:"power,string,omitempty"`
	Combat        uint32 `json:"combat,string,omitempty"`
	ActualStamina uint8  `json:"actual_stamina,string,omitempty"`
}

type Hero struct {
	Id                   uint32 `json:"id,string,omitempty"`
	Name                 string `json:"name,omitempty"`
	PowerStats           Stats  `json:"powerstats,omitempty"`
	Biography            Bio    `json:biography,omitempty`
	FiliationCoefficient float32
}

type Bio struct {
	Alignment string `json:"alignment,omitempty"`
}

type Team struct {
	Heroes    [5]Hero
	Alignment string
}

func (team *Team) UpdateFiliationCoefficient() {
	for i, hero := range team.Heroes {
		if hero.Biography.Alignment == team.Alignment {
			team.Heroes[i].FiliationCoefficient = float32(1 + rand.Intn(10))
		} else {
			team.Heroes[i].FiliationCoefficient = 1.0 / float32(1+rand.Intn(10))
		}
	}
}

func BuildHeroTeam(heroes [5]Hero) Team {
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

	team := Team{
		Heroes:    heroes,
		Alignment: teamAlignment,
	}

	return team
}
