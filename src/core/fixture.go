package core

import (
	"fmt"
	"time"

	"bravian1/team-shuffler/src/types"
)

func Fixture(teams []types.Teams, startDate string, intervalDays int) []string {
	n := len(teams)
	if n%2 != 0 {
		teams = append(teams, types.Teams{Name: "", Striker: "", Defender: ""})
		n += 1
	}

	start, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		fmt.Println("Error parsing start date:", err)
		return []string{}
	}

	fixtures := []string{}
	for round := 1; round < n; round++ {
		round_date := start.AddDate(0, 0, (round-1)*intervalDays).Format("2006-01-02")
		round_match := " [" + round_date + ": "
		match_list := []string{}
		for match := 0; match <= n/2-1; match++ {
			home := (round + match) % (n - 1)
			away := (n - 1 - match + round) % (n - 1)
			if match == 0 {
				away = n - 1
			}
			if !(teams[away].Name == "" || teams[home].Name == "") {
				match_list = append(match_list, teams[home].Name+" vs "+teams[away].Name)
			}
		}
		round_match += fmt.Sprintf("%s", match_list)
		round_match += "]"
		fixtures = append(fixtures, round_match)
	}

	return fixtures
}
