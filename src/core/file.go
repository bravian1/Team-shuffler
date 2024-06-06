package core

import (
	"fmt"
	"os"
	"path"
	"sync"
)

func OpenOrCreate(filename string) (string, bool) {
	rootDir, err := os.Getwd()
	if err != nil {
		return "Error accessing directory...", false
	}
	filename = path.Join(rootDir, "storage", filename)

	file, err := os.Open(filename)
	if err != nil {
		return fmt.Sprintf("Failed to open %q\n", filename), false
	}
	defer file.Close()
	return filename, true
}

func WriteLine(mutex *sync.Mutex, filename string, data any) bool {
	rootDir, err := os.Getwd()
	if err != nil {
		return false
	}
	filename = path.Join(rootDir, "storage", filename)
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
