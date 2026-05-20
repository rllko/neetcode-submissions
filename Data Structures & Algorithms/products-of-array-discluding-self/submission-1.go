func productExceptSelf(nums []int) []int {
    resp := make([]int,len(nums));
    for i, _ := range nums{
        result := 1 
        for j, _ := range nums {
            if i == j {
                continue;
            }
            result *= nums[j]
        }
        resp[i] = result;
    }
    return resp;
}