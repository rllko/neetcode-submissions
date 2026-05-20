func topKFrequent(nums []int, k int) []int {
	occurences := map[int]int{}

	for _, num := range nums {
		if _, ok := occurences[num]; !ok {
			occurences[num] = 0
		}

		occurences[num] += 1
	}

	result := []int{}
	for num := range occurences {
		result = append(result, num)
	}

	sort.Slice(result, func(i, j int) bool {
		return occurences[result[i]] > occurences[result[j]]
	})

	return result[:k]
}
