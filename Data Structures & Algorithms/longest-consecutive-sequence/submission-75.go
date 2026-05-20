func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// Sorts in-place. Uses virtually zero extra RAM.
	sort.Ints(nums)

	largestCount := 1
	currentCount := 1

	for i := 1; i < len(nums); i++ {
		// Skip duplicate numbers so they don't break the streak
		if nums[i] == nums[i-1] {
			continue
		}

		// If consecutive, increase the current streak
		if nums[i] == nums[i-1]+1 {
			currentCount++
		} else {
			// Streak broke. Update largestCount using a fast if-statement
			if currentCount > largestCount {
				largestCount = currentCount
			}
			currentCount = 1 // Reset streak
		}
	}

	// Final check for the last sequence
	if currentCount > largestCount {
		largestCount = currentCount
	}

	return largestCount
}
