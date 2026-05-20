class Solution {
    /**
     * @param {number[]} nums
     * @param {number} k
     * @return {number[]}
     */
    topKFrequent(nums, k) {
        const occurrences = {}

        for(let num of nums){
            if(!occurrences[num]){
                occurrences[num] = 0
            }

            occurrences[num] += 1 
        }

        const output = Object.keys(occurrences)
                            .sort((a, b) => occurrences[b] - occurrences[a])
                            .slice(0,k)


        return output
    }
}
