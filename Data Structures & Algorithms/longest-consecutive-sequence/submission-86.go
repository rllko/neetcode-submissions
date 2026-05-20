func longestConsecutive(nums []int) int {
	  if len(nums) == 0 {
        return 0
    }

    sort.Ints(nums);

    length := 1;
		biggestLength := 0

		left := 0
    for right,num := range nums {
			left = right + 1
			if left >= len(nums) {
				break;
			}
			if (num+1) == nums[left] {
				length++
			} else if num == nums[left] {
				continue;
			} else {
				if length > biggestLength { biggestLength = length}
				length = 1
			}
		}
		if length > biggestLength { biggestLength = length}

	return biggestLength
}
