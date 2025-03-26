package testgolang

func Shuffle(str string) []string {
	return genPermutations(str)
}

func genPermutations(str string) []string {
	if str == "" {
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
