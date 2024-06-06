package core

import (
	"log"
	"math/rand"
	"os"
	"time"
)

func Shuffle(file *os.File) map[string]string {
	strikers, defenders, err := ReadTeams(file)
	if err != nil {
		log.Fatal(err)
	}
	rand.New(rand.NewSource(time.Now().UnixNano()))
	rand.Shuffle(len(defenders), func(i, j int) {
		defenders[i], defenders[j] = defenders[j], defenders[i]
	})
	team := make(map[string]string)
	for i := 0; i < len(strikers); i++ {
		team[strikers[i]] = defenders[i]
	}
	return team
}
