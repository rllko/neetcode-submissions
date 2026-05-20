func longestConsecutive(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    sort.Ints(nums);

    largestCount := 1;
    currentCount := 1;

    for i:= 1; i < len(nums); i++ {
        val := nums[i]
        val_behind := nums[i-1]

        if val == val_behind {
            continue;
        }

        if val == val_behind + 1 {
            currentCount += 1
        } else {
            if currentCount > largestCount {
                 largestCount = currentCount
            }

            currentCount = 1
        }
    }

	if currentCount > largestCount {
		largestCount = currentCount
	}

	return largestCount
}
