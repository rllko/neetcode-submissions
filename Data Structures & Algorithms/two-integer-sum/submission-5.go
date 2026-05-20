func twoSum(nums []int, target int) []int {
   m := map[int]int{}

   for i,_ := range nums {
        diff := target - nums[i];

        if _, ok := m[diff]; ok {
            return []int{m[diff],i}
        }

        m[nums[i]] = i;  
   } 

   return []int{}
}
