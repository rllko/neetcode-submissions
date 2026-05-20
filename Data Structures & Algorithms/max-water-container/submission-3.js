class Solution {
    /**
     * @param {number[]} heights
     * @return {number}
     */
    maxArea(heights) {
			let biggest = -1;

			for(let i = 0,j = heights.length-1; i < j;){
				const val = (j - i) * Math.min(heights[i],heights[j]);
				
				if(val > biggest){
					biggest = val
				}

				if(heights[i] < heights[j]){
					i++;
				} else {
					j--;
				}
			}

			return biggest
		}
}
