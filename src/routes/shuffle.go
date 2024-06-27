package routes

import (
	"encoding/json"
	"net/http"

	"bravian1/team-shuffler/src/core"
	"bravian1/team-shuffler/src/data"
	"bravian1/team-shuffler/src/types"
)

func Shuffle(w http.ResponseWriter, r *http.Request) {
	db, err := data.ConnectDB()
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}

	var strikers []types.Player
	var defenders []types.Player

	// Fetch strikers and defenders
	if err := db.Where("role = ?", "Striker").Find(&strikers).Error; err != nil {
		http.Error(w, "Database query error for strikers", http.StatusInternalServerError)
		return
	}
	if err := db.Where("role = ?", "Defender").Find(&defenders).Error; err != nil {
		http.Error(w, "Database query error for defenders", http.StatusInternalServerError)
		return
	}

	if len(strikers) != len(defenders) {
		http.Error(w, "Uneven number of strikers and defenders", http.StatusInternalServerError)
		return
	}

	// Shuffle defenders
	shuffledDefenders := core.Shuffle(defenders)

	// Create teams
	teams := []types.Team{}
	for i := 0; i < len(strikers); i++ {
		teamName := core.GenerateString(strikers[i].Name, shuffledDefenders[i].Name)
		team := types.Team{Name: teamName, Striker: strikers[i].Name, Defender: shuffledDefenders[i].Name}
		teams = append(teams, team)
	}

	// Save teams to the database
	// mutex.Lock()
	for _, team := range teams {
		if err := db.Create(&team).Error; err != nil {
			// mutex.Unlock()
			http.Error(w, "Error saving teams to database", http.StatusInternalServerError)
			return
		}
	}
	// mutex.Unlock()

	content, err := json.Marshal(teams)
	if err != nil {
		http.Error(w, "Error creating teams json", http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(content)
}
