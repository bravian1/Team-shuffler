package data

import (
	"fmt"
	"os"
	"strconv"

	"bravian1/team-shuffler/src/types"

	"github.com/go-gota/gota/dataframe"
)

func GetPlayers() []types.Table {
	file, err := os.Open("storage/team.csv")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
	}

	csvDf := dataframe.ReadCSV(file)
	sorted := csvDf.Arrange(
		dataframe.RevSort("Points"),
	)
	// convert the dataframe to a slice of structs
	table := make([]types.Table, 0)
	records := sorted.Records()
	for i, row := range records {
		// Skip the header row
		if i == 0 {
			continue
		}
		table = append(table, types.Table{
			TeamName:     row[0],
			Played:       toInt(row[1]),
			Wins:         toInt(row[2]),
			Draws:        toInt(row[3]),
			Losses:       toInt(row[4]),
			GoalsFor:     toInt(row[5]),
			GoalsAgainst: toInt(row[6]),
			Points:       toInt(row[7]),
		})
	}
	return table
}

func toInt(str string) int {
	val, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return val
}
