package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"bravian1/team-shuffler/src/core"
	"bravian1/team-shuffler/src/types"
)

// Fixtures generates and returns the fixtures in JSON format
func Fixtures(w http.ResponseWriter, r *http.Request) {

	teams := []types.Team{}
	if err := DB.Find(&teams).Error; err != nil {
		http.Error(w, "Error fetching teams", http.StatusInternalServerError)
		return
	}
	gameweeks := core.Fixture(teams)
	for _, gw:=range gameweeks{
		if err:=DB.Create(&gw).Error; err!=nil{
			http.Error(w, "Error creating gameweek", http.StatusInternalServerError)
            return
        }
	}


	fmt.Println(gameweeks)
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(gameweeks)
}
