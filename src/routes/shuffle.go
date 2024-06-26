package routes

import (
	"encoding/json"
	"net/http"
	"sync"

	"bravian1/team-shuffler/src/core"
	"bravian1/team-shuffler/src/types"

	"gorm.io/gorm"
)



func Shuffle(mutex *sync.Mutex) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mutex.Lock()
		defer mutex.Unlock()

		var players []types.Player
		if err := DB.Find(&players).Error; err != nil {
			http.Error(w, "Error fetching players", http.StatusInternalServerError)
			return
		}

		var strikers, defenders []string
		for _, player := range players {
			if player.Role == "Striker" {
				strikers = append(strikers, player.Name)
			} else if player.Role == "Defender" {
				defenders = append(defenders, player.Name)
			}
		}

		defenders = core.Shuffle(defenders)

		var teams []types.Team
		for i := 0; i < len(strikers); i++ {
			teamName := core.GenerateString(strikers[i], defenders[i])
			team := types.Team{Name: teamName, Striker: strikers[i], Defender: defenders[i]}
			teams = append(teams, team)
		}

		// Clear existing teams and insert new ones
		if err := DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&types.Team{}).Error; err != nil {
			http.Error(w, "Error clearing existing teams", http.StatusInternalServerError)
			return
		}

		if err := DB.Create(&teams).Error; err != nil {
			http.Error(w, "Error creating new teams", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(teams)
	}
}
