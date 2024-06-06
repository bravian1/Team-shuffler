package core

import (
	"math/rand"
	"time"
)

func Shuffle(arr []string) []string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
	return arr
}
