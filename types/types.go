package types

import "math/rand"

type Stats struct {
	Intelligence float32 `json:"intelligence,string,omitempty"`
	Strength     float32 `json:"strength,string,omitempty"`
	Speed        float32 `json:"speed,string,omitempty"`
	Durability   float32 `json:"durability,string,omitempty"`
	Power        float32 `json:"power,string,omitempty"`
	Combat       float32 `json:"combat,string,omitempty"`
}

type Hero struct {
	Id                   uint32 `json:"id,string,omitempty"`
	Name                 string `json:"name,omitempty"`
	PowerStats           Stats  `json:"powerstats,omitempty"`
	Biography            Bio    `json:"biography,omitempty"`
	FiliationCoefficient float32
	HP                   float32
}

type Bio struct {
	Alignment string `json:"alignment,omitempty"`
}

type Team struct {
	Heroes    [5]Hero
	Alignment string
}

func getActualStamina() float32 {
	return rand.Float32() * 10
}

func updateStat(base float32, fb float32) float32 {
	return float32(2*base+getActualStamina()) / 1.1 * fb

}

func (hero *Hero) updateHp() {
	tmp := hero.PowerStats.Strength*0.8 + hero.PowerStats.Durability*0.7 + hero.PowerStats.Power
	hero.HP = 100 + tmp/2*(1+getActualStamina()/10)
}

// We make it private to avoid issues
func (hero *Hero) updateHeroStats() {
	hero.PowerStats.Strength = updateStat(hero.PowerStats.Strength, hero.FiliationCoefficient)
	hero.PowerStats.Intelligence = updateStat(hero.PowerStats.Intelligence, hero.FiliationCoefficient)
	hero.PowerStats.Combat = updateStat(hero.PowerStats.Combat, hero.FiliationCoefficient)
	hero.PowerStats.Durability = updateStat(hero.PowerStats.Durability, hero.FiliationCoefficient)
	hero.PowerStats.Power = updateStat(hero.PowerStats.Power, hero.FiliationCoefficient)
	hero.PowerStats.Speed = updateStat(hero.PowerStats.Speed, hero.FiliationCoefficient)
	hero.updateHp()
}

func (team *Team) UpdateFiliationCoefficient() {
	for i, hero := range team.Heroes {
		if hero.Biography.Alignment == team.Alignment {
			team.Heroes[i].FiliationCoefficient = float32(1 + rand.Intn(10))
		} else {
			team.Heroes[i].FiliationCoefficient = 1.0 / float32(1+rand.Intn(10))
		}
		// Having the FC we can update power stats
		team.Heroes[i].updateHeroStats()
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
