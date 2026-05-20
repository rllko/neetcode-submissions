class Solution {
    /**
     * @param {number[]} heights
     * @return {number}
     */
    maxArea(heights) {
			let biggest = -1;

			for(let i = 0,j = heights.length-1; i < j;){

				let h = heights[j];

				if(heights[i] < heights[j]) {
					h = heights[i]
				}

				const val = (j - i) * h;
				
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
