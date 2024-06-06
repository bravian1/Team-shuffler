package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"bravian1/team-shuffler/src/core"
	"bravian1/team-shuffler/src/types"
)

func Players(w http.ResponseWriter, r *http.Request) {
	strikers, defenders, err := core.ReadTeams("players.txt")
	if err != nil {
		http.Error(w, "Error reading users", http.StatusInternalServerError)
		return
	}
	allplayer := []types.Player{}
	for _, ch := range defenders {
		allplayer = append(allplayer, types.Player{Role: "Defender", Name: ch})
	}
	for _, ch := range strikers {
		allplayer = append(allplayer, types.Player{Role: "Striker", Name: ch})
	}

	fmt.Println(allplayer)
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(allplayer)
}
