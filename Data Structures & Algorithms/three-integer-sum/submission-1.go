import "slices"

func threeSum(nums []int) [][]int {
    slices.Sort(nums)
    res := [][]int{}
    l:= 0
    r:= 0
    for i:= 0; i < len(nums); i++ {

        r = len(nums)-1
        l = i+1
        for r > l {
            val := nums[r] + nums[l]

            if r < len(nums)-1 && (nums[r+1] == nums[r]) && nums[l-1] == nums[l] {
                fmt.Println("NIGGER")
                break
            }

            if val + nums[i] == 0 {
                res = append(res,[]int{nums[i],nums[l],nums[r]})
                l++
            } 

            if val + nums[i] > 0 {
                r--
            } else {
                l++
            }
        }
    }

    return res;
}
