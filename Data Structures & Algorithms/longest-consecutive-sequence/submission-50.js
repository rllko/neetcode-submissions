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
        
        // Start counts at 1, because any non-empty array has a minimum sequence of 1
        let currentCount = 1;
        let longestCount = 1;

        // Start checking from the second element (index 1) to compare with the previous
        for (let i = 1; i < sorted.length; i++) {
            
            // Only do something if the numbers are different (this safely ignores duplicates)
            if (sorted[i] !== sorted[i - 1]) {
                
                // If they are exactly 1 apart, the sequence continues
                if (sorted[i] === sorted[i - 1] + 1) {
                    currentCount += 1;
                } else {
                    // The sequence broke! Save the highest count so far and reset
                    longestCount = Math.max(longestCount, currentCount);
                    currentCount = 1;
                }
            }
        }

        // Return the max of longestCount and currentCount 
        // (needed in case the longest sequence goes all the way to the end of the array)
        return Math.max(longestCount, currentCount);
    }
}