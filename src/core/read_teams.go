package core

import (
	"bufio"
	"log"
	"os"
	"path"
	"strings"
)

func ReadTeams(filename string) ([]string, []string, error) {
	rootDir, err := os.Getwd()
	if err != nil {
		return nil, nil, err
	}
	filename = path.Join(rootDir, "storage", filename)
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var strikers []string
	var defenders []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(line, " ")
		if strings.ToLower(words[0]) == "striker:" {
			strikers = append(strikers, words[1])
		} else if strings.ToLower(words[0]) == "defender:" {
			defenders = append(defenders, words[1])
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
	// if len(striker) == 0 || len(defender) == 0 || len(striker) != len(defender) {
	// 	return striker, defender, fmt.Errorf("number of strikers and defenders are not equal")
	// }
	return strikers, defenders, nil
}
