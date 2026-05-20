func isPalindrome(s string) bool {
	re := regexp.MustCompile(`(?i)[ .!@#$%^&*()_+\-=?,:\\'/]`)
	new_s := re.ReplaceAllString(s, "")
	new_s = strings.ToLower(new_s)
    for i, j := 0, len(new_s)-1; i <= j; {
        if new_s[i] != new_s[j] {
            return false
        }

        i++;
        j--;
    }

    return true;
}
