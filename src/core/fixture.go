package core

import (
	"strconv"

	"bravian1/team-shuffler/src/types"
)

func Fixture(teams []types.Team) []types.GameWeek {
	n := len(teams)
	if n%2 != 0 {
		teams = append(teams, types.Team{Name: "", Striker: "", Defender: ""})
		n += 1
	}
	gameweek := []types.GameWeek{}

	date := 1
	for round := 1; round < n-1; round++ {
		week := types.GameWeek{
			Week:    date,
			Matches: []types.Fixture{},
		}
		for match := 0; match < n/2; match++ {
			home := (round + match) % (n - 1)
			away := (n - 1 - match + round) % (n - 1)
			if match == 0 {
				away = n - 1
			}
			week.Matches = append(week.Matches, types.Fixture{Date: strconv.Itoa(date), Home: teams[home].Name, Away: teams[away].Name})
			// if !(teams[away].Name == "" || teams[home].Name == "") {
			// 	match_list = append(match_list, teams[home].Name+" vs "+teams[away].Name+" \n")
			// }
			// fixtures = append(fixtures, types.Fixture{Date: "Week " + fmt.Sprintf("%d", round) + " Fixtures", Home: teams[home].Name, Away: teams[away].Name})
		}
		gameweek = append(gameweek, week)
		date++
	}
	// gameweek = append(gameweek, types.GameWeek{Week: 1, Matches: fixtures})

	return gameweek
}
