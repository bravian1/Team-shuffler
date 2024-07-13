package routes

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"bravian1/team-shuffler/src/types"

	"gorm.io/gorm"
)

func Register(mutex *sync.Mutex) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("name")
		role := r.FormValue("role")

		if name == "" || role == "" {
			http.Error(w, "Name and role are required", http.StatusBadRequest)
			return
		}

		player := types.Player{
			Name: name,
			Role: role,
		}

		if err := DB.Create(&player).Error; err != nil {
			http.Error(w, "Error creating player", http.StatusInternalServerError)
			return
		}

		// Return HTML fragment for the new player row
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, "<tr><td>%s</td><td>%s</td></tr>", player.Role, player.Name)
	}
}
func deletePlayerByName(db *gorm.DB, name string) error {
	player := types.Player{
		Name: name,
	}
	result := db.Where("name = ? AND deleted_at IS NULL", player.Name).Delete(&player)
	if result.Error != nil {
		return result.Error
	}

	log.Printf("Number of rows deleted: %d\n", result.RowsAffected)
	return nil
}

func DeletePlayer(mutex *sync.Mutex) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("name")

		if name == "" {
			http.Error(w, "Name required", http.StatusBadRequest)
			return
		}

		// Ensure thread safety using the mutex
		mutex.Lock()
		defer mutex.Unlock()

		// Perform the delete operation
		if err := deletePlayerByName(DB, name); err != nil {
			http.Error(w, "Error deleting player", http.StatusInternalServerError)
			return
		}

		// Respond with success message
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Player with Name: %s deleted successfully", name)
	}
}
