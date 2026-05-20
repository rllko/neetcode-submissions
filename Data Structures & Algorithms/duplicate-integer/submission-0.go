func hasDuplicate(nums []int) bool {
    for i,el := range nums {
        for j := i + 1; j <  len(nums); j++ {
            if nums[j] == el{
                return true
            }
        }
    }
    return false;
}
