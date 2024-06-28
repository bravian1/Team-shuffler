package core

import (
	"time"

	"bravian1/team-shuffler/src/types"
)

func Fixture(teams []types.Team) []types.Fixture {
	n := len(teams)
	if n%2 != 0 {
		teams = append(teams, types.Team{Name: "", Striker: "", Defender: ""})
		n += 1
	}
	
	var fixtures []types.Fixture
	week := 1
	for round := 1; round < n-1; round++ {
		for match := 0; match < n/2; match++ {
			home := (round + match) % (n - 1)
			away := (n - 1 - match + round) % (n - 1)
			if match == 0 {
				away = n - 1
			}
			if teams[home].Name != "" && teams[away].Name != "" {
				fixtures = append(fixtures, types.Fixture{
					Home: teams[home].Name,
					Away: teams[away].Name,
					Date: time.Now().AddDate(0, 0, week-1).Format("2006-01-02"),
					Week: week,
				})
			}
		}
		week++
	}

	return fixtures
}
