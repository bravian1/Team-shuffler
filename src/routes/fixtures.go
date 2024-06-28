package routes

import (
	"encoding/json"
	"net/http"

	"bravian1/team-shuffler/src/core"
	"bravian1/team-shuffler/src/data"
	"bravian1/team-shuffler/src/types"
)

// Fixtures generates and returns the fixtures in JSON format
func Fixtures(w http.ResponseWriter, r *http.Request) {
	// Check if DB is nil
	if DB == nil {
		http.Error(w, "Database connection is not initialized", http.StatusInternalServerError)
		return
	}
	db, err := data.ConnectDB()
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}

	teams := []types.Team{}
	if err := db.Find(&teams).Error; err != nil {
		http.Error(w, "Error fetching teams", http.StatusInternalServerError)
		return
	}

	fixtures := core.Fixture(teams)
	for _, fixture := range fixtures {
		if err := db.Create(&fixture).Error; err != nil {
			http.Error(w, "Error creating fixture", http.StatusInternalServerError)
			return
		}
	}

	content, err := json.Marshal(fixtures)
	if err != nil {
		http.Error(w, "Error creating fixtures json", http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(content)
}
