func getConcatenation(nums []int) []int {
	l := len(nums)
	ans := make([]int, 0, l * 2)
	
	for i := 0; i < 2; i++ {
		for _, n := range(nums){
			ans = append(ans, n)
		}
	}

	return ans
}