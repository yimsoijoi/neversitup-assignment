package testgolang

import "regexp"

var (
	smileyRegex = regexp.MustCompile(`^[:;][-~]?[)D]$`)
)

func CountSmileys(arr []string) int {
	c := 0
	for _, a := range arr {
		if smileyRegex.MatchString(a) {
			c++
		}
	}

	return c
}
