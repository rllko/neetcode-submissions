func productExceptSelf(nums []int) []int {
    length := len(nums)
    resp := make([]int,len(nums));

    prefix := 1
    for i, num := range nums{
      resp[i] = prefix
      prefix *= num
    }

    suffix := 1
    for j:= length -1; j>= 0; j-- {
        resp[j] *= suffix
        suffix *= nums[j]
    }

    return resp;
}