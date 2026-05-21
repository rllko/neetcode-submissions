func characterReplacement(s string, k int) int {
	l := 0
	m := 0
	frequence := map[rune]int{}
	maxFreq := 0

	for r,ch := range s {
		if _,ok := frequence[ch]; !ok {
			frequence[ch] = 1
		} else {
			frequence[ch]++
		}
		
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
