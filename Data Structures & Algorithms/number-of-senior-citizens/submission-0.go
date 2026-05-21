import "slices"

func countSeniors(details []string) int {
	idx := 0
	for idx < len(details) {
		s := details[idx]
		val, _ := strconv.Atoi(s[11:13])

		if val <= 60 {
			details = slices.Delete(details, idx, idx+1)
		} else {
			idx++
		}
		
	}
	return len(details)
}