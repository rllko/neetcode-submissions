func isPalindrome(s string) bool {
	input := strings.ToLower(s);

    cleanedText := ""

    for _,txt := range input {
        if (txt >= 'a' && txt <= 'z') || (txt >= '0' && txt <= '9') {
            cleanedText += string(txt);
        }
    }

    left := 0
    right := len(cleanedText) - 1;
    
    for left < right {
        if cleanedText[left] != cleanedText[right] {
            return false
        }

        left++
        right--
    }

    return true;
}
