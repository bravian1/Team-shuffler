package core

import (
	"os"
	"sync"
)

func WriteLine(mutex *sync.Mutex, filename string, data any) bool {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		return false
	}
	defer file.Close()

	str, ok := data.(string)
	if !ok {
		return false
	}

	mutex.Lock()
	_, err = file.WriteString(str)
	mutex.Unlock()

	return err == nil
}
