func isPalindrome(s string) bool {
	input := strings.ToLower(s);

    cleanedText := ""

    for _,txt := range input {
        if (txt >= 'a' && txt <= 'z') || (txt >= '0' && txt <= '9') {
            cleanedText += string(txt);
        }
    }
    
    for i := 0; i < len(cleanedText)/2;i++ {
        if cleanedText[i] != cleanedText[len(cleanedText)-1-i] {
            return false
        }
    }

    return true;
}
