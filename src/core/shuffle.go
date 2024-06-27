package core

import (
	"bravian1/team-shuffler/src/types"
	"math/rand"
	"time"
)

func Shuffle(players []types.Player) []types.Player {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	rand.Shuffle(len(players), func(i, j int) {
		players[i], players[j] = players[j], players[i]
	})
	return players
}
