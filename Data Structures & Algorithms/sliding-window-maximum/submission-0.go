func maxSlidingWindow(nums []int, k int) []int {
    left := 0
    res := make([]int,0)

    for i:= k-1; i < len(nums);i++ {
        idx := left
        max := nums[idx]
        for idx <= i{
            if nums[idx] > max {
                max = nums[idx]
            }
            idx++
        }
        res = append(res,max)
        left++
    } 
    return res;
}
