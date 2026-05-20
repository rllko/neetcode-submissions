import "slices"

func threeSum(nums []int) [][]int {
    slices.Sort(nums)
    
    res := [][]int{}
    l:= 0
    r:= 0

    for i:= 0; i < len(nums); i++ {
        r = len(nums)-1
        l = i+1

        if i > 0 && nums[i-1] == nums[i] {
            continue;
        }

        for r > l {
            val := nums[r] + nums[l] + nums[i]

            if val == 0 {
                res = append(res,[]int{nums[i],nums[l],nums[r]})
                l++
                r--

                for l < r && nums[l] == nums[l-1] {
                    l++
                }
                
                for l < r && nums[r] == nums[r+1] {
                    r--
                }
            } else if val > 0 {
                r--
            } else {
                l++
            }
        }
    }

    return res;
}
