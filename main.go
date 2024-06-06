package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	handler "bravian1/team-shuffler/routes"
)

func main() {
	mutex :=  &sync.Mutex{}

	file, err := os.Open("teams.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	http.HandleFunc("/players", handler.Players)
	http.HandleFunc("/shuffle", handler.Shuffle)
	http.HandleFunc("/register", handler.Register(mutex))
	http.HandleFunc("/", handler.Index)
	fmt.Println("Server started on port 8000")
	http.ListenAndServe(":8000", nil)
}
