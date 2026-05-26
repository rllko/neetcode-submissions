func minWindow(s string, t string) string {
    arr := make(map[byte]int)
    for i := 0; i < len(t); i++ {
        arr[t[i]]++
    }

    need := len(arr)
    have := 0
    match_arr := make(map[byte]int)

    l := 0
    bestL, bestLen := 0, len(s)+1

    for r := 0; r < len(s); r++ {
        c := s[r]
        match_arr[c]++

        if match_arr[c] == arr[c] {
            have++
        }

        for have == need {
            if r-l+1 < bestLen {
                bestL = l
                bestLen = r - l + 1
            }
            match_arr[s[l]]--
            if match_arr[s[l]] < arr[s[l]] {
                have--
            }
            l++
        }
    }

    if bestLen > len(s) {
        return ""
    }
    return s[bestL : bestL+bestLen]
}