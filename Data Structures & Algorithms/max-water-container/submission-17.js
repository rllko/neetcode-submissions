class Solution {
    /**
     * @param {number[]} heights
     * @return {number}
     */
    maxArea(heights) {
			let biggest = -1;
			let i = 0,j = heights.length-1;
			let h = heights[j];

			let output = (i-j) * h;

			while(i < j){
				
				h = heights[j];
				if(heights[i] < heights[j]) {
					h = heights[i]
				}

				output = (j - i) * h;

				if(output > biggest){
					biggest = output
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
