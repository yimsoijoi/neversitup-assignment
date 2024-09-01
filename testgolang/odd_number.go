package testgolang

func OddNumber(nums []int) int {
	odd := make(map[int]struct{})
	for i := range nums {
		if _, ok := odd[nums[i]]; ok {
			delete(odd, nums[i])
			continue
		}
		odd[nums[i]] = struct{}{}
	}

	return len(odd)
}
