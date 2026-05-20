import "slices"

func lengthOfLongestSubstring(s string) int {
	list := []rune{}
	longestStreak := 0

	for _, c := range s {
        idx := slices.Index(list, c)

		if idx != -1 {
			list = slices.Delete(list, 0, idx+1)
		}

        list = append(list, c)

		// The size of our list IS our current streak! No need for a separate counter.
		if len(list) > longestStreak {
			longestStreak = len(list)
		}
    }
    return longestStreak
}