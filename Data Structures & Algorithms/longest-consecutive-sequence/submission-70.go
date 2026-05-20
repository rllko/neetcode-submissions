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
            if len(nums)  <= largestCount {
                return currentCount
            }

            largestCount = int(math.Max(float64(largestCount),float64(currentCount)))

            currentCount = 1
        }
    }

    return int(math.Max(float64(largestCount),float64(currentCount)))

}
