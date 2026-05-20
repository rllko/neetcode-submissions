func topKFrequent(nums []int, k int) []int {
	// Step 1: Count the frequencies of each element
    counts := make(map[int]int)
    for _, num := range nums {
        counts[num]++
    }

    // Step 2: Create buckets where the array index represents the frequency.
    // The maximum possible frequency is the length of the array itself.
    buckets := make([][]int, len(nums)+1)
    for num, count := range counts {
        buckets[count] = append(buckets[count], num)
    }

    // Step 3: Gather the top k frequent elements.
    // We iterate from the end (highest frequency) to the beginning.
    var result []int
    for i := len(buckets) - 1; i >= 0 && len(result) < k; i-- {
        if buckets[i] != nil {
            result = append(result, buckets[i]...)
        }
    }

    // Return exactly k elements (handles edge cases where multiple numbers 
    // tie in frequency at the boundary, though LeetCode guarantees unique answers)
    return result[:k]
}
