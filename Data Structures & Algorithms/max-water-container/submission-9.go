func maxArea(heights []int) int {
	left := 0 
	right := len(heights)- 1
	maxArea := 0

	for left < right {
		width := right - left;

		h := heights[right]

		if heights[right] > heights[left] {
			h = heights[left]
		}

		area := width * h

		if area > maxArea {
			maxArea = area
		}

		if heights[left] < heights[right] {
			left++
		}else {
			right--
		}
	}

	return maxArea;
}
