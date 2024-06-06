package core

import (
	"bravian1/team-shuffler/src/types"
)

func Fixture(teams []types.Teams) []string {
	n := len(teams)
	if n%2 != 0 {
		teams = append(teams, types.Teams{Name: "", Striker: "", Defender: ""})
		n += 1
	}

	fixtures := []string{}
	for round := 1; round < n; round++ {
		round_match := []string{}
		for match := 0; match <= n/2-1; match++ {
			home := (round + match) % (n - 1)
			away := (n - 1 - match + round) % (n - 1)
			if match == 0 {
				away = n - 1
			}
			if !(teams[away].Name == "" || teams[home].Name == "") {
				round_match = append(round_match, teams[home].Name+":"+teams[away].Name)
			}
		}
		fixtures = append(fixtures, round_match...)
	}
	return fixtures
}
