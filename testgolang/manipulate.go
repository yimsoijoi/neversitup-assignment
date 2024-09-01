package testgolang

import (
	"fmt"
)

/*
a b     a b c		a b c d
b a     a c b		a b d c
        b a c		a c b d
		b c a		a c d b
		c a b		a d b c
		c b a		a d c b
				    b a c d
					. . . .
					d c b a
*/

func Shuffle(str string) []string {
	results := []string{}
	perms := GenPermutations(str)
	seen := map[string]struct{}{}
	for i := range perms {
		if _, ok := seen[perms[i]]; !ok {
			results = append(results, perms[i])
			seen[perms[i]] = struct{}{}
		}
	}

	return results
}

func GenPermutations(str string) []string {
	fmt.Println(str)
	if len(str) == 0 {
		return []string{}
	}
	if len(str) == 1 {
		return []string{str}
	}

	permutations := []string{}

	for i, c := range str {
		remaining := str[:i] + str[i+1:]
		subPermutations := GenPermutations(remaining)

		for _, p := range subPermutations {
			permutations = append(permutations, string(c)+p)
		}
	}

	return permutations
}
