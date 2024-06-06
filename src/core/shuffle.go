package core

import (
	"log"
	"math/rand"
	"time"
)

func Shuffle(filename string) map[string]string {
	strikers, defenders, err := ReadTeams(filename)
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
