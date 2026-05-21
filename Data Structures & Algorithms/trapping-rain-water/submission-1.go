func trap(height []int) int {
		left,right := 0,len(height)-1
		left_max,right_max := height[left],height[right]
		res := 0
		
		for left < right  {
				if left_max < right_max {
					left++
					left_max = max(left_max,height[left])
					res += left_max - height[left]
				} else {
					right--
					right_max = max(right_max,height[right])
					res += right_max - height[right]
				}
		}

	return res
}
