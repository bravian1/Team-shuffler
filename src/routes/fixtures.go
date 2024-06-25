package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path"

	"bravian1/team-shuffler/src/core"
	"bravian1/team-shuffler/src/types"
)

// Fixtures generates and returns the fixtures in JSON format
func Fixtures(w http.ResponseWriter, r *http.Request) {
	rootDir, err := os.Getwd()
	if err != nil {
		http.Error(w, "Error creating fixtures", http.StatusInternalServerError)
		return
	}
	filename := path.Join(rootDir, "storage", "teams.txt")

	stats, err := os.Stat(filename)
	if err != nil {
		http.Error(w, "Error creating fixtures", http.StatusInternalServerError)
		return
	}
	if stats.Size() == 0 {
		http.Error(w, "Error creating fixtures", http.StatusInternalServerError)
		return
	}

	content, err := os.ReadFile(filename)
	if err != nil {
		http.Error(w, "Error creating fixtures", http.StatusInternalServerError)
		return
	}
	teams := []types.Teams{}

	err = json.Unmarshal(content, &teams)
	if err != nil {
		http.Error(w, "Error creating fixtures", http.StatusInternalServerError)
		return
	}

	gameweeks := core.Fixture(teams)

	fmt.Println(gameweeks)
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(gameweeks)
}
