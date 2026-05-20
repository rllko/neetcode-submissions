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

		longestStreak = max(len(list),longestStreak)
		
    }
    return max(len(list),longestStreak)
}