class Solution {
    /**
     * @param {number[]} heights
     * @return {number}
     */
    maxArea(heights) {
			let biggest = -1;

			for(let i = 0; i < heights.length; i++){
				for(let j = heights.length;j > i;j--){
					const val = (j - i) * Math.min(heights[i],heights[j]);
					if(val > biggest){
						biggest = val
					}
				}
			}

			return biggest
		}
}
