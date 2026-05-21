func trap(height []int) int {
	res := 0

	for i := range height {
	left,right := 0,len(height)-1
        left_max:= 0
        right_max:= 0
        
        for right > left  {

            if height[left] > left_max && left < i {
                left_max = height[left]
            } else if left < i {
                left++
            }

            if height[right] > right_max && right > i {
                right_max = height[right]
            } else if right > i {             
                right--
            }
        }

        water := min(left_max,right_max) - height[i]

        if water > 0 {
            res += water
        }
	}

    return res
}
