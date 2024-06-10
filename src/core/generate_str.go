package core

import "strings"

func GenerateString(s1, s2 string) string {
	news2 := strings.Fields(s2)[1]
	return s1[:2] + news2
}
