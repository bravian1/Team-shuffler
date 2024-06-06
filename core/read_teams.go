package core

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func ReadTeams(file *os.File) ([]string, []string, error) {
	var striker []string
	var defender []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(line, " ")
		if strings.ToLower(words[0]) == "striker:" {
			striker = append(striker, words[1])
		} else if strings.ToLower(words[0]) == "defender:" {
			defender = append(defender, words[1])
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
	// if len(striker) == 0 || len(defender) == 0 || len(striker) != len(defender) {
	// 	return striker, defender, fmt.Errorf("number of strikers and defenders are not equal")
	// }
	return striker, defender, nil
}
