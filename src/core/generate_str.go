package core

func GenerateString(s1, s2 string) string {
	return s1[0:2] + s2[:len(s2)-3]
}