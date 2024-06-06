package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"bravian1/team-shuffler/types"
	handler "bravian1/team-shuffler/routes"
)

var mu sync.Mutex

func main() {
	file, err := os.Open("teams.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	http.HandleFunc("/players", playerHandler)
	http.HandleFunc("/shuffle", shuffleHandlers)
	http.HandleFunc("/register", registerHandlers)
	http.HandleFunc("/", handler.Index)
	fmt.Println("Server started on port 8000")
	http.ListenAndServe(":8000", nil)
}


func playerHandler(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("teams.txt")
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
	}
	defer file.Close()
	strikers, defenders, err := ReadTeams(file)
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
func shuffleHandlers(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("teams.txt")
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
	}
	defer file.Close()
	strikers, defenders, err := ReadTeams(file)
	if err != nil {
		log.Fatal(err)
	}
	rand.Seed(time.Now().UnixNano())
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

// func Shuffle(file *os.File)map[string]string {
// 	strikers, defenders, err := ReadTeams(file)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	rand.Seed(time.Now().UnixNano())
// 	rand.Shuffle(len(defenders), func(i, j int) {
// 		defenders[i], defenders[j] = defenders[j], defenders[i]
// 	})
// 	team := make(map[string]string)
// 	for i := 0; i < len(strikers); i++ {
// 		team[strikers[i]] = defenders[i]
// 	}
// 	return team
// }
func ReadTeams(file *os.File) ([]string, []string, error) {
	var striker []string
	var defender []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(line, " ")
		if strings.ToLower(words[0]) == "striker:" {
			striker = append(striker, words[1])
		} else if strings.ToLower(words[0]) == "defender:" {
			defender = append(defender, words[1])
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
	// if len(striker) == 0 || len(defender) == 0 || len(striker) != len(defender) {
	// 	return striker, defender, fmt.Errorf("number of strikers and defenders are not equal")
	// }
	return striker, defender, nil
}

func registerHandlers(w http.ResponseWriter, r *http.Request) {
	role := r.FormValue("role")
	name := r.FormValue("name")
	fmt.Println(role + " " + name)
	if role == "" || name == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	line := fmt.Sprintf("\n%s: %s", role, name)

	file, err := os.OpenFile("teams.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		http.Error(w, "Failed to open file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	_, err = file.WriteString(line)
	if err != nil {
		http.Error(w, "Failed to write to file", http.StatusInternalServerError)
		return
	}
	// players := []Player{}
	mu.Lock()
	player := types.Player{Role: role, Name: name}
	mu.Unlock()
	w.Header().Set("content-type", "application.json")
	json.NewEncoder(w).Encode(player)
	// http.Redirect(w, r, "/", http.StatusSeeOther)

}
