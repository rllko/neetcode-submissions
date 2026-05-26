func characterReplacement(s string, k int) int {
	l := 0
	m := 0
	frequence := make(map[rune]int)
	maxFreq := 0

	for r,ch := range s {
		frequence[ch]++
		
		maxFreq = max(maxFreq,frequence[ch])

		length := (r-l+1) - maxFreq

		for length > k {
			frequence[rune(s[l])]--
			l++
			length = (r-l+1) - maxFreq
		}
		
		m = max(m,r - l + 1)
	}
 
	return m
}
