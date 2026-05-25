
func checkInclusion(s1 string, s2 string) bool {
    count1 := make(map[rune]int)
    for _,ch := range s1 {
        count1[ch]++
    }

    need := len(count1)
    for i:= 0; i< len(s2); i++ {
        count2 := make(map[rune]int)
        curr := 0
        for j := i; j < len(s2); j++ {
            ch_rune := rune(s2[j]);
            count2[ch_rune]++
            if count1[rune(s2[j])] < count2[ch_rune] {
                break
            }
            if count1[rune(s2[j])] == count2[ch_rune] {
                curr++
            }
            if curr == need {
                return true
            }
        }
    }

    return false;
}
