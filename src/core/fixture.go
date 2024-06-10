package core

import (
	"bravian1/team-shuffler/src/types"
)

func Fixture(teams []types.Teams) [][]string {
	n := len(teams)
	if n%2 != 0 {
		teams = append(teams, types.Teams{Name: "", Striker: "", Defender: ""})
		n += 1
	}

	fixtures := [][]string{}
	for round := 1; round < n-1; round++ {
		//round_match :=[]string{}
		match_list := []string{}
		for match := 0; match < n/2; match++ {
			home := (round + match) % (n - 1)
			away := (n - 1 - match + round) % (n - 1)
			if match == 0 {
				away = n - 1
			}
			if !(teams[away].Name == "" || teams[home].Name == "") {
				match_list = append(match_list, teams[home].Name+" vs "+teams[away].Name+" \n")
			}
		}

		fixtures = append(fixtures, match_list)
	}

	return fixtures
}
