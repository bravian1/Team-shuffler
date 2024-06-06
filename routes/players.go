package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"bravian1/team-shuffler/types"
	"bravian1/team-shuffler/core"
)

func Players(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("teams.txt")
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
	}
	defer file.Close()
	strikers, defenders, err := core.ReadTeams(file)
	if err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
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
