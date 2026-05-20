class Solution {
    /**
     * @param {number[]} nums
     * @return {number}
     */
    longestConsecutive(nums) {
        if (nums.length === 0) {
            return 0;
        }

        const sorted = nums.sort((a, b) => a - b);
        
        let currentCount = 1;
        let longestCount = 1;

        for (let i = 1; i < sorted.length; i++) {
            
            if (sorted[i] === sorted[i - 1]) {
                continue;
            }

            if (sorted[i] === sorted[i - 1] + 1) {
                currentCount += 1;
            } else {
                longestCount = Math.max(longestCount, currentCount);

                if(nums.length <= currentCount){
                    return currentCount
                }
                
                currentCount = 1;
            }
        }

       
        return Math.max(longestCount, currentCount);
    }
}