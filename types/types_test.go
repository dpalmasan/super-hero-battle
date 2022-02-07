package types

import (
	"testing"
)

func TestBuildHeroTeam(t *testing.T) {

	heroes := [5]Hero{
		{
			Biography: Bio{
				Alignment: "good",
			},
		},
		{
			Biography: Bio{
				Alignment: "good",
			},
		},
		{
			Biography: Bio{
				Alignment: "bad",
			},
		},
		{
			Biography: Bio{
				Alignment: "bad",
			},
		},
		{
			Biography: Bio{
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

func TestUpdateFiliationCoefficient(t *testing.T) {

	heroes := [5]Hero{
		{
			Biography: Bio{
				Alignment: "good",
			},
		},
		{
			Biography: Bio{
				Alignment: "good",
			},
		},
		{
			Biography: Bio{
				Alignment: "bad",
			},
		},
		{
			Biography: Bio{
				Alignment: "bad",
			},
		},
		{
			Biography: Bio{
				Alignment: "bad",
			},
		},
	}

	team := BuildHeroTeam(heroes)
	team.UpdateFiliationCoefficient()

	if team.Heroes[0].FiliationCoefficient > 1 || team.Heroes[0].FiliationCoefficient == 0 {
		t.Fatal(`Expected FC to be <= 1 and != 0 for hero 0`)
	}

	if team.Heroes[3].FiliationCoefficient < 1 || team.Heroes[0].FiliationCoefficient == 0 {
		t.Fatal(`Expected FC to be >= 1 and != 0 for hero 3`)
	}
}
