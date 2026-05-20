class Solution {
    /**
     * @param {number[]} heights
     * @return {number}
     */
    maxArea(heights) {
			let biggest = -1;
			let i = 0,j = heights.length-1;

			while(i < j){
				const min = Math.min(heights[i], heights[j])
				biggest = Math.max((j - i) * min, biggest)

				if(heights[i] < heights[j]){
					i++;
				} else {
					j--;
				}
			}

			return biggest
		}
}
