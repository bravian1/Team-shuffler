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

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Sprintf("Failed to create %q\n", filename), false
	}
	defer file.Close()
	return filename, true
}

func WriteStringToFile(mutex *sync.Mutex, filename string, str string) bool {
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

	mutex.Lock()
	_, err = file.WriteString(str)
	mutex.Unlock()

	return err == nil
}

func WriteBytesToFile(mutex *sync.Mutex, filename string, data []byte) bool {
	rootDir, err := os.Getwd()
	if err != nil {
		return false
	}
	filename = path.Join(rootDir, "storage", filename)
	if err := os.Truncate(filename, 0); err != nil {
		return false
	}
	file, err := os.OpenFile(filename, os.O_WRONLY, 0o644)
	if err != nil {
		return false
	}
	defer file.Close()

	mutex.Lock()
	_, err = file.Write(data)
	mutex.Unlock()

	return err == nil
}
