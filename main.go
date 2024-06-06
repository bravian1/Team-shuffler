package main

import (
	"log"
	"net/http"
	"sync"

	"bravian1/team-shuffler/src/core"
	handler "bravian1/team-shuffler/src/routes"
)

func main() {
	result, success := core.OpenOrCreate("teams.txt")
	if !success {
		log.Fatalf(result)
	} else {
		println("Successfully created or opened: ", result)
	}

	mutex := &sync.Mutex{}

	http.HandleFunc("/players", handler.Players)
	http.HandleFunc("/fixtures", handler.Fixtures)
	http.HandleFunc("/shuffle", handler.Shuffle(mutex))
	http.HandleFunc("/register", handler.Register(mutex))
	http.HandleFunc("/", handler.Index)

	print("\n\n\tServer started on port 8000\n\n")
	http.ListenAndServe(":8080", nil)
}
