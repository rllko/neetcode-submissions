func findMin(nums []int) int {
    left := 0
    right := len(nums)-1
    
    smallest := nums[left]
    for left <= right {
        if nums[left] < smallest {
            smallest = nums[left]
        }

        if nums[right] < smallest {
            smallest = nums[right]
        }

        left++
        right--
    }

    return smallest;
}
