package testgolang

import (
	"fmt"
)

func Shuffle(str string) []string {
	results := []string{}
	perms := genPermutations(str)
	seen := map[string]struct{}{}
	for i := range perms {
		if _, ok := seen[perms[i]]; !ok {
			results = append(results, perms[i])
			seen[perms[i]] = struct{}{}
		}
	}

	return results
}

func genPermutations(str string) []string {
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
		subPermutations := genPermutations(remaining)

		for _, p := range subPermutations {
			permutations = append(permutations, string(c)+p)
		}
	}

	return permutations
}
