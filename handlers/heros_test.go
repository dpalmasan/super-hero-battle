package handlers

import (
	"testing"

	"github.com/super-hero-battle/types"
)

func TestGenerateRandomIds(t *testing.T) {

	ids, err := GenerateRandomIds(10, 3, 9)
	if err == nil || len(ids) != 0 {
		t.Fatal(`Expected error when size < n`)
	}

	ids, err = GenerateRandomIds(10, 3, 100)
	if len(ids) != 10 {
		t.Fatal("Expected an array of ids of length 10!")
	}
}

func TestBuildHeroTeam(t *testing.T) {

	heroes := [5]types.Hero{
		{
			Biography: types.Bio{
				Alignment: "good",
			},
		},
		{
			Biography: types.Bio{
				Alignment: "good",
			},
		},
		{
			Biography: types.Bio{
				Alignment: "bad",
			},
		},
		{
			Biography: types.Bio{
				Alignment: "bad",
			},
		},
		{
			Biography: types.Bio{
				Alignment: "bad",
			},
		},
	}

	team := BuildHeroTeam(heroes)
	if team.Alignment != "bad" {
		t.Fatal(`Expected team alignment to be bad`)
	}

	heroes[2].Biography.Alignment = "good"
	team = BuildHeroTeam(heroes)
	if team.Alignment != "good" {
		t.Fatal(`Expected team alignment to be good`)
	}
}
