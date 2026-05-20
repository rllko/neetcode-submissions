class Solution:
    def longestConsecutive(self, nums: List[int]) -> int:
        if len(nums) == 0:
            return 0

        counter = 1
        best = -1
        nums.sort()

        for i in range(len(nums)-1):
            if nums[i+1] == nums[i]:
                continue; # same value doesnt count

            if nums[i+1] == nums[i]+1:
                counter+=1
            else: # streak broke
                best = max(counter,best)
                counter = 1 # reset counter

        return max(counter,best)
                