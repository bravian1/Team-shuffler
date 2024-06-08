package core

import (
	"encoding/json"
	"os"
	"path"

	"bravian1/team-shuffler/src/types"
)

func ReadConfig() (*types.Config, bool) {
	rootDir, err := os.Getwd()
	if err != nil {
		return nil, false
	}
	filename := path.Join(rootDir, "config.json")

	stats, err := os.Stat(filename)
	if err != nil {
		return nil, false
	}
	if stats.Size() == 0 {
		return nil, false
	}

	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, false
	}
	config := types.Config{}

	err = json.Unmarshal(content, &config)
	if err != nil {
		return nil, false
	}

	return &config, true
}
