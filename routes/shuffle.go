package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"bravian1/team-shuffler/core"
	"bravian1/team-shuffler/types"
)

func Shuffle(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("teams.txt")
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
	}
	defer file.Close()
	strikers, defenders, err := core.ReadTeams(file)
	if err != nil {
		log.Fatal(err)
	}
	rand.New(rand.NewSource(time.Now().UnixNano()))
	rand.Shuffle(len(defenders), func(i, j int) {
		defenders[i], defenders[j] = defenders[j], defenders[i]
	})
	team := []types.Teams{}
	for i := 0; i < len(strikers); i++ {
		team = append(team, types.Teams{Striker: strikers[i], Defender: defenders[i]})
	}
	fmt.Println(team)

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(team)
}
