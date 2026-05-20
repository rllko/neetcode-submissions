func productExceptSelf(nums []int) []int {
    resp := make([]int,len(nums));
    for i, _ := range nums{
        result := 1 
        for j, j_num := range nums {
            if i == j {
                continue;
            }
            result *= j_num
        }
        resp[i] = result;
    }
    return resp;
}