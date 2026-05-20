import "slices"

func lengthOfLongestSubstring(s string) int {
	list := []rune{}
	longestStreak := 0

	for _, c := range s {
		if idx := slices.Index(list, c); idx != -1 {
			list = slices.Delete(list, 0, idx+1)
		}

    list = append(list, c)

		if len(list) > longestStreak {
			longestStreak = len(list)
		}
    }
    return longestStreak
}