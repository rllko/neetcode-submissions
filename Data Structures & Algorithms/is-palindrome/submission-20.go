func isPalindrome(s string) bool {
	re := regexp.MustCompile("[^a-zA-Z0-9]+")
	new_s := re.ReplaceAllString(s, "")
	new_s = strings.ToLower(new_s)
    for i := 0; i < len(new_s)/2;i++ {
        if new_s[i] != new_s[len(new_s)-1-i] {
            return false
        }
    }

    return true;
}
