package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"bravian1/team-shuffler/src/core"
	"bravian1/team-shuffler/src/types"
)

func Shuffle(mutex *sync.Mutex) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		strikers, defenders, err := core.ReadTeams("players.txt")
		if err != nil {
			log.Fatal(err)
		}
		defenders = core.Shuffle(defenders)
		teams := []types.Teams{}
		for i := 0; i < len(strikers); i++ {
			teamName := core.GenerateString(strikers[i], defenders[i])
			teams = append(teams, types.Teams{Name: teamName, Striker: strikers[i], Defender: defenders[i]})
		}
		// fmt.Println(teams)

		content, err := json.Marshal(teams)
		if err != nil {
			http.Error(w, "Error creating teams json structure", http.StatusInternalServerError)
			return
		}

		if ok := core.WriteBytesToFile(mutex, "teams.txt", content); !ok {
			http.Error(w, "Error creating fixtures", http.StatusInternalServerError)
			return
		}

		w.Header().Set("content-type", "application/json")
		json.NewEncoder(w).Encode(teams)
	}
}
