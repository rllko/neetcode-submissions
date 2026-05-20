func hasDuplicate(nums []int) bool {
    seen := make(map[int]bool)
    for _, el := range nums {
        if seen[el] {
            return true
        }
        seen[el] = true
    }
    return false
}